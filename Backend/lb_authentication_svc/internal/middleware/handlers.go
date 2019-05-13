package middleware

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"lb_authentication_svc/configs"
	"lb_authentication_svc/internal/model/contracts/request"
	"lb_authentication_svc/internal/model/contracts/response"
	"lb_authentication_svc/internal/model/database"
	token "lb_authentication_svc/internal/model/jwt"
	"net/http"
	"time"
)

var env *configs.ViperConfigReader

func SetConfig(conf *configs.ViperConfigReader) {
	env = conf
}

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestJson request.Credentials
		// Get the JSON body and decode into credentials

		err := json.NewDecoder(c.Request.Body).Decode(&requestJson)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			c.JSON(http.StatusBadRequest, response.AccessDenied{http.StatusBadRequest, env.GetString("messages.response.login.failure.bad-request")})
			return
		}

		var user database.User
		db.
			Where("username = ?", requestJson.Username).
			Where("password = ?", requestJson.Password).
			First(&user)

		if user.ID == 0 {
			//	User not found
			c.JSON(http.StatusUnauthorized, response.AccessDenied{http.StatusUnauthorized, env.GetString("messages.response.login.failure.unauthorized")})
			return
		}

		//If the password is expected then Start the token creation process
		// Declare the expiration time of the token
		// here, we have kept it as 5 minutes
		expirationTime := time.Now().Add(30 * time.Minute)
		// Create the JWT claims, which includes the username and expiry time
		claims := &token.Claims{
			UserID:   requestJson.ID,
			IP:       c.Request.RemoteAddr,
			Username: requestJson.Username,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
				Audience:  requestJson.Username,
				Subject:   env.GetString("jwt.claims.standard.subject"),
				IssuedAt:  time.Now().Unix(),
				Issuer:    env.GetString("jwt.claims.standard.issuer"),
			},
		}

		// Declare the token with the algorithm used for signing, and the claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// Create the JWT string
		key := []byte(env.GetString("jwt.signature"))
		tokenString, err := token.SignedString(key)
		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			c.JSON(http.StatusInternalServerError, response.AccessDenied{http.StatusInternalServerError, env.GetString("messages.response.login.failure.internal-server-error.jwt")})
			return
		}

		// Setting the client cookie for "token" as the JWT we just generated
		// we also set an expiry time which is the same as the token itself
		c.SetCookie(env.GetString("cookie.name"), tokenString, expirationTime.Minute(), "", "", true, true)

		// Finally, if everything went well, so it is time to create a success response
		c.JSON(http.StatusOK, response.AccessPermitted{http.StatusOK, env.GetString("messages.response.login.success"), tokenString})
	}
}
