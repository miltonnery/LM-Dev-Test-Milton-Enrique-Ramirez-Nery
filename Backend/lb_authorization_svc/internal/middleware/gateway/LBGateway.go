package gateway

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"lb_authorization_svc/configs"
	"lb_authorization_svc/internal/middleware/security"
	"lb_authorization_svc/internal/model/config"
	"lb_authorization_svc/internal/model/contracts/response"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func Redirector(env *configs.ViperConfigReader) gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Method == "OPTIONS" {
			//	Handle preflight here
			buildPreFlightResponse(c)
			c.JSON(http.StatusOK, response.Accepted{env.GetString("messages.response.pre-flight.status.accepted")})
			return
		}

		//Validating requested endpoint´s existence
		logrus.Info("VALIDATING EXISTENCE OF REQUESTED ENDPOINT")
		result, uri := Find(c.Request.Method, c.Request.URL.Path)
		if result {
			logrus.Info("ENDPOINT MAPPING FOUND")

			logrus.Info("LOOKING FOR SECURITY HEADER EXISTENCE")
			header := c.Request.Header.Get("Authorization")

			if header == "" {
				logrus.Info("NO AUTHORIZATION HEADER PROVIDED")
				if uri.AccessLevel == 0 {
					logrus.Info("REQUESTED ENDPOINT IS PUBLIC")
					buildPublicRequest(c.Request, uri)
					redirect(c)
				} else {
					logrus.Info("REQUESTED ENDPOINT HAS A RESTRICTED ACCESS")
					c.JSON(
						http.StatusServiceUnavailable,
						response.AccessDenied{http.StatusServiceUnavailable, env.GetString("messages.response.authorization.failure.access-denied")})
				}
			} else {
				logrus.Info("GETTING JWT FROM AUTHORIZATION HEADER")
				splitToken := strings.Split(header, "Bearer ")
				header = splitToken[1]

				logrus.Info("TOKEN FOUND! EXTRACTING VALUES")
				_, validToken := security.ExtractClaims(header)

				//Initializing token validation layer
				if validToken {
					logrus.Info("EXTRACTING CLAIMS")
					//c1, _ := claims["tid"].(string)
					//c2, _ := claims["sub"].(string)

					if uri.AccessLevel == 0 {
						logrus.Info("REQUESTED ENDPOINT IS PUBLIC")
						buildPublicRequest(c.Request, uri)
						redirect(c)
					} else {
						logrus.Info("REQUESTED ENDPOINT HAS A RESTRICTED ACCESS")
						logrus.Info("VERIFYING TOKEN STATUS")
						logrus.Info("REDIRECTING TO REQUESTED ENDPOINT")
						buildSecuredRequest(c.Request, uri)
						redirect(c)
					}
				} else {
					logrus.Error("THE TOKEN PROVIDED IS INVALID")
					c.JSON(http.StatusServiceUnavailable, response.AccessDenied{http.StatusServiceUnavailable, env.GetString("messages.response.authorization.failure.access-denied")})
				}
			}
		} else {
			logrus.Info("ENDPOINT NOT FOUND")
			c.JSON(http.StatusNotFound, response.AccessDenied{http.StatusNotFound, env.GetString("messages.response.authorization.failure.access-denied")})
		}
	}
}

func buildPreFlightResponse(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Request-Headers", "Authorization, Content-Type, Logout")
	c.Header("Access-Control-Request-Methods", "")
	c.Header("Access-Control-Expose-Headers", "Access-Level")
	c.Header("Connection", "keep-alive")
	c.Header("Date", time.Now().String())
}

func makeDirector(req *http.Request) {
	// Save a copy of this request for debugging.
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
}

func buildPublicRequest(req *http.Request, uri *config.Uri) {
	req.Host = uri.RedirectionHost
	req.URL.Scheme = "http"
	req.URL.Host = uri.RedirectionHost
	req.URL.Path = uri.RedirectionPath

	//Cleaning headers
	req.Header.Del("tkn")

	//Adding headers
	req.Header.Add("X-Forwarded-For", req.RemoteAddr)
}

func buildSecuredRequest(req *http.Request, uri *config.Uri) {
	req.Host = uri.RedirectionHost
	req.URL.Scheme = "http"
	req.URL.Host = uri.RedirectionHost
	req.URL.Path = uri.RedirectionPath

	//Cleaning headers
	req.Header.Del("Authorization")

	//Adding headers
	req.Header.Add("X-Forwarded-For", req.RemoteAddr)
}

func redirect(c *gin.Context) {
	Director := makeDirector
	proxy := &httputil.ReverseProxy{Director: Director}
	proxy.ServeHTTP(c.Writer, c.Request)
}
