package middleware

import (
	"github.com/auth0-community/go-auth0"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/square/go-jose.v2"
	"net/http"

	"github.com/derhabicht/rose-park/controllers"
)

func GetValidator() *auth0.JWTValidator {
	client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: viper.GetString("AUTH0_JWK")}, nil)
	audience := []string{viper.GetString("AUTH0_API_AUDIENCE")}
	configuration := auth0.NewConfiguration(client, audience, "", jose.RS256)

	return auth0.NewValidator(configuration, nil)
}

func Authorize(validator *auth0.JWTValidator) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		tok, err := validator.ValidateRequest(c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, controllers.ControllerError{"error": "invalid token"})
			logrus.WithFields(logrus.Fields{
				"module": "auth",
				"error": err,
			}).Error("Invalid auth token provided to API")
			return
		}

		claims := make(map[string]interface{})
		err = validator.Claims(c.Request, tok, &claims)
		logrus.WithFields(logrus.Fields{
			"module": "auth",
			"claims": claims,
		}).Debug("")

		c.Next()
	})
}
