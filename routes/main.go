package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"log"
	"net/http"
)

var (
	Router *gin.Engine
)

func tlsHandler(port string) gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":" + port,
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		if err != nil {
			return
		}
		c.Next()
	}
}

// Run will start the server
func Run(port, sslCertPath, sslKeyPath string, trustedProxies []string, getRoutes func(*gin.Engine)) {
	Router = gin.Default()
	if len(trustedProxies) != 0 {
		Router.ForwardedByClientIP = true
		err := Router.SetTrustedProxies(trustedProxies)
		if err != nil {
			log.Fatalf("Wrong TrustedProxies, %v", err)
		}
	}
	Router.Use(passCors)
	getRoutes(Router)
	if len(sslCertPath) != 0 && len(sslKeyPath) != 0 {
		log.Println("SSL is enabled.")
		Router.Use(tlsHandler(port))
		log.Fatal(Router.RunTLS(":"+port, sslCertPath, sslKeyPath))
	} else {
		log.Println("SSL is not enabled.")
		log.Fatal(Router.Run(":" + port))
	}
}

func passCors(c *gin.Context) {
	method := c.Request.Method

	// FIXME: Not safe
	origin := c.GetHeader("Origin")
	c.Header("Access-Control-Allow-Origin", origin)
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	// 放行所有OPTIONS请求
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.Next()
}
