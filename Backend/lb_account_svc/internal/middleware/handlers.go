package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"lb_account_svc/configs"
	"lb_account_svc/internal/middleware/parsers"
	"lb_account_svc/internal/model/contracts/request"
	"lb_account_svc/internal/model/contracts/response"
	"lb_account_svc/internal/model/database"
	"lb_account_svc/internal/tooling"
	"net/http"
	"time"
)

var env *configs.ViperConfigReader

func SetConfig(conf *configs.ViperConfigReader) {
	env = conf
}

// HANDLERS
// Lists all the products for a specific client
func ListMyProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//Variables declaration
		var user database.User

		//Catching the request params
		logrus.Info("Catching the request params")
		values := c.Request.URL.Query()
		user.Username = values.Get("user")

		//Starting database search process
		logrus.Info("Starting database search process")
		db.Preload("Role").Where(&user).First(&user)

		//var transactions []database.Transaction
		//db.Where("contract = ?", 1).Find(&transactions)

		//Loading database values
		logrus.Info("Loading database values")
		if user.Active == true {
			logrus.Info("The user is active!")
			logrus.Info("Searching for contracted services")
			var client database.Client
			client.UserInfo = user
			client.UserInfoID = user.ID
			client.ID = user.ID
			db.Preload("Contracts").
				Preload("Contracts.Product").
				Preload("Contracts.Client").
				Preload("Contracts.Account").
				Preload("Contracts.CreditCard").
				Preload("Contracts.Loan").
				Preload("Contracts.Product.Type").
				Find(&client)
			logrus.Info("Parsing response")
			var jsonResponse response.ContractedServices
			for _, v := range client.Contracts {
				if v.CreditCard.ID != 0 {
					var cc response.CreditCardResponse
					cc.Name = v.Product.Name
					cc.CardLimit = v.CreditCard.CardLimit
					cc.CardNumber = v.CreditCard.CardNumber
					cc.GoodThru = v.CreditCard.GoodThru
					cc.Cvv = v.CreditCard.Cvv
					cc.MonthlyCut = v.CreditCard.MonthlyCut
					cc.InterestAmount = v.CreditCard.InterestAmount
					cc.InterestRate = v.CreditCard.InterestRate
					jsonResponse.CreditCards = append(jsonResponse.CreditCards, cc)
				}
				if v.AccountID != 0 {
					var ac response.AccountResponse
					ac.Name = v.Product.Name
					ac.Number = v.Account.Number
					ac.Nickname = v.Account.NameID
					jsonResponse.Accounts = append(jsonResponse.Accounts, ac)
				}
				if v.LoanID != 0 {
					var lo response.LoanResponse
					lo.Name = v.Product.Name
					lo.Amount = v.Loan.Amount
					lo.InterestRate = v.Loan.InterestRate
					lo.DuePayment = v.Loan.DuePayment
					jsonResponse.Loans = append(jsonResponse.Loans, lo)
				}
			}
			logrus.Info("showing response")
			c.JSON(http.StatusOK, jsonResponse)
		} else {
			logrus.Warn("USER NOT FOUND")
			c.JSON(http.StatusBadRequest, response.Failure{Code: http.StatusBadRequest, Message: "User not found"})
		}

	}
}

// Lists all the transactions for a specific account
func ProductTransactions(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		values := c.Request.URL.Query()

		logrus.Info("Validating query")

		accountNumber := values.Get("accountID")
		start := values.Get("start")
		end := values.Get("end")

		if (accountNumber != "") && (start != "") && (end != "") {
			start = start + "T00:00:00Z"
			end = end + "T00:00:00Z"
			//	Dates validation
			timeStart, err := time.Parse(time.RFC3339, start)
			if err != nil {
				c.JSON(http.StatusBadRequest, response.Failure{http.StatusBadRequest, "The request period is major than 3 months"})
				return
			}
			timeEnd, err := time.Parse(time.RFC3339, end)
			if err != nil {
				c.JSON(http.StatusBadRequest, response.Failure{http.StatusBadRequest, "The request period is major than 3 months"})
				return
			}

			if timeStart.Unix() > timeEnd.Unix() {
				c.JSON(http.StatusBadRequest, response.Failure{http.StatusBadRequest, "The start time cannot be supperior to the end time"})
				return
			}
			period := timeEnd.Unix() - timeStart.Unix()
			threeMonths := int64(60 * 60 * 24 * 30 * 3)
			if period > threeMonths {
				logrus.Info("Period is too long")
				c.JSON(http.StatusBadRequest, response.Failure{http.StatusBadRequest, "The request period is major than 3 months"})
				return
			} else {
				logrus.Info("The period is acceptable")
				logrus.Info("Searching for account and transactions existence")
				logrus.Info("Searching for the account: ", accountNumber)

				var account database.Account
				account.Number = accountNumber
				db.Where("id = ?", accountNumber).First(&account)

				if account.ID == 0 {
					c.JSON(http.StatusNotFound, response.Failure{http.StatusNotFound, "The account number doesn´t exist"})
					return
				}

				var contract database.Contract
				contract.Account = account
				contract.AccountID = account.ID
				db.Preload("Client").
					Preload("CreditCard").
					Preload("Loan").
					Preload("Product").
					First(&contract)

				var transactions []database.Transaction
				db.
					Preload("Contract").
					Preload("Contract.Product").
					Preload("Contract.CreditCard").
					Preload("Contract.Loan").
					Preload("Contract.Account").
					Where("contract = ?", contract.ID).
					Find(&transactions)
				c.JSON(http.StatusOK, parsers.ParseTransactions(transactions))
			}
		} else {
			c.JSON(http.StatusBadRequest, response.Failure{http.StatusBadRequest, "Bad parameters"})
		}

	}
}

// Enrollment of new beneficiary
func Enrollment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var jsonRequest request.BeneficiaryEnrollment
		err := json.NewDecoder(c.Request.Body).Decode(&jsonRequest)
		if err != nil {
			logrus.Error(err)
		}

		//Account
		var ownerAccount database.Account
		db.Where("number = ?", jsonRequest.OwnerAccount).First(&ownerAccount)

		var ownerContract database.Contract
		ownerContract.AccountID = ownerAccount.ID
		db.
			Preload("Client").
			Preload("Client.UserInfo").
			Preload("Client.UserInfo.Role").
			Preload("Account").
			Preload("CreditCard").
			Preload("Loan").
			Preload("Product").
			Where("account = ?", ownerAccount.ID).
			First(&ownerContract)

		//Beneficiary
		var beneficiaryAccount database.Account
		db.Where("number = ?", jsonRequest.BeneficiaryAccount).First(&beneficiaryAccount)

		var beneficiaryContract database.Contract
		db.
			Preload("Client").
			Preload("Client.UserInfo").
			Preload("Client.UserInfo.Role").
			Preload("Account").
			Preload("CreditCard").
			Preload("Loan").
			Preload("Product").
			Where("account = ?", beneficiaryAccount.ID).First(&beneficiaryContract)

		//Saving beneficiary
		var beneficiary database.Beneficiary
		beneficiary.Active = true
		beneficiary.OwnerID = ownerContract.Client.UserInfoID
		beneficiary.Owner = ownerContract.Client.UserInfo
		beneficiary.ReceiverID = beneficiaryContract.Client.UserInfoID
		beneficiary.Identifier = tooling.MakeUUID()
		beneficiary.Email = jsonRequest.Email
		beneficiary.Receiver = beneficiaryContract.Client.UserInfo
		beneficiary.CreatedDate = time.Now().UTC()
		beneficiary.CreatedBy = "lb_configurator_account"
		beneficiary.ModifiedDate = time.Now().UTC()
		beneficiary.ModifiedBy = "lb_configurator_account"

		if beneficiary.OwnerID != beneficiary.ReceiverID {
			// The owner and receiver are different, we can continue then
			if beneficiaryContract.Client.Email == jsonRequest.Email {
				//	The emails are ok
				if (beneficiaryContract.Client.FirstName == jsonRequest.BeneficiaryFirstName) && (beneficiaryContract.Client.LastName == jsonRequest.BeneficiaryLastName) {
					//	First and last name coincides
					//Cheking duplicity
					var ben database.Beneficiary
					db.
						Where("owner = ?", beneficiary.OwnerID).
						Where("receiver = ?", beneficiary.ReceiverID).
						First(&ben)
					if ben.ID != 0 {
						//	Beneficiary already linked
						c.JSON(http.StatusConflict, response.Failure{Code: http.StatusConflict, Message: "The beneficiary is already linked to this user"})
						return
					} else {
						//	Save beneficiary
						db.Save(&beneficiary)
						c.JSON(http.StatusAccepted, response.Success{Code: http.StatusAccepted, Message: "Beneficiary saved successfully"})
						return
					}
				} else {
					//	Name doesn´t match
					c.JSON(http.StatusBadRequest, response.Failure{Code: http.StatusBadRequest, Message: "The name doesn´t match with the one of the internal registry"})
					return
				}
			} else {
				//	Error the emails doesn´t match
				c.JSON(http.StatusBadRequest, response.Failure{Code: http.StatusBadRequest, Message: "The email doesn´t match with the one of the internal registry"})
				return
			}
		} else {
			// The owner and receiver are the same! we can´t continue
			c.JSON(http.StatusBadRequest, response.Failure{Code: http.StatusBadRequest, Message: "The owner is the same as the receiver"})
			return
		}

	}
}

//Updates the email of the beneficiary specified
func UpdateEmail(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		values := c.Request.URL.Query()
		beneficiaryID := values.Get("beneficiaryID")
		email := values.Get("email")
		var beneficiary database.Beneficiary
		db.
			Preload("Receiver").
			Where("identifier = ?", beneficiaryID).
			First(&beneficiary)

		//Beneficiary code validation
		if beneficiary.Identifier == beneficiaryID {
			//	Beneficiary code found!
			if tooling.ValidateEmail(email) {
				//	Valid Email
				beneficiary.Email = email
				db.Save(&beneficiary)
				c.JSON(http.StatusNoContent, "")
				return
			} else {
				//	Invalid Email
				c.JSON(http.StatusBadRequest, response.Failure{Code: http.StatusBadRequest, Message: "Invalid Email"})
				return
			}
		} else {
			c.JSON(http.StatusNotFound, response.Failure{Code: http.StatusNotFound, Message: "The beneficiary code does not exists"})
			return
		}
	}
}

//Deletes the email of the beneficiary specified
func Delete(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		values := c.Request.URL.Query()
		beneficiaryID := values.Get("beneficiaryID")
		var beneficiary database.Beneficiary
		db.
			Preload("Receiver").
			Where("identifier = ?", beneficiaryID).
			First(&beneficiary)

		//Beneficiary code validation
		if beneficiary.Identifier == beneficiaryID {
			//	Beneficiary code found!
			db.Delete(&beneficiary)
			c.JSON(http.StatusNoContent, "")
			return
		} else {
			c.JSON(http.StatusNotFound, response.Failure{Code: http.StatusNotFound, Message: "The beneficiary code does not exists"})
			return
		}
	}
}
