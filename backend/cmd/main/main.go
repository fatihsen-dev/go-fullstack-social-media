package main

import (
	"log"
	"net/http"

	"github.com/fatihsen-dev/go-fullstack-social-media/pkg/routes"
	"github.com/fatihsen-dev/go-fullstack-social-media/pkg/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main(){
	gin.SetMode(gin.ReleaseMode)

	//? create gin router
	router := gin.Default()

	//? allow cors
	config := cors.DefaultConfig()
   config.AllowOrigins = []string{"*"}
   router.Use(cors.New(config))

	//? Routers
	mainRouter := router.Group("/api")
	{
		routes.UsersRouter(mainRouter.Group("/user"))
		routes.PostsRouter(mainRouter.Group("/post"))
	}

	router.GET("/api",func (c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"message":"Api Ready"})
	})
	
	//? starting server
	log.Printf("Server started on port %s", utils.GetEnvVariable("PORT"))
	http.ListenAndServe(":"+utils.GetEnvVariable("PORT"), router)
}