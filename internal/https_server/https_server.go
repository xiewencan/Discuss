package https_server

import (
	v1 "discuss/api/v1"
	// "discuss/internal/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var GE *gin.Engine

func init() {
	GE = gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	GE.Use(cors.New(corsConfig))
	// GE.Use(ssl.TlsHandler(config.GetConfig().MainConfig.Host, config.GetConfig().MainConfig.Port))
	// GE.Static("/static/avatars", config.GetConfig().StaticAvatarPath)
	// GE.Static("/static/files", config.GetConfig().StaticFilePath)
	GE.POST("/login", v1.Login)
	// GE.POST("/register", v1.Register)

}
