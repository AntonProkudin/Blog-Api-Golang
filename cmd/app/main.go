package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	_ "test/cmd/app/docs"
	"test/controllers"
	"test/initializers"
	"test/routes"
)

var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	PostController      controllers.PostController
	PostRouteController routes.PostRouteController

	CommController      controllers.CommController
	CommRouteController routes.CommRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	PostController = controllers.NewPostController(initializers.DB)
	PostRouteController = routes.NewRoutePostController(PostController)

	CommController = controllers.NewCommController(initializers.DB)
	CommRouteController = routes.NewRouteCommController(CommController)

	server = gin.Default()

}

// @title						APIs
// @version					1.0
// @description				Testing Swagger APIs.
// @termsOfService				http://swagger.io/terms/
// @contact.name				API Support
// @contact.url				http://www.swagger.io/support
// @contact.email				support@swagger.io
// @securityDefinitions.apiKey	JWT
// @in							header
// @name						token
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @host						localhost:8081
// @BasePath					/api/v1
// @securityDefinitions.apiKey	JWT
// @in							header
// @name						token
func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	PostRouteController.PostRoute(router)
	CommRouteController.CommRoute(router)
	//server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Fatal(server.Run(":" + config.ServerPort))

}
