package main

import (
  "fmt"
	"go_trial_2/handlers"

  //"net/http"
  "github.com/gin-gonic/gin"
   // _ "github.com/lib/pq"
  //  "database/sql"
    )

func main() { 

	router := gin.Default()  
	router.POST("/campaign", handlers.AddCampaign)
  router.POST("/addecision", handlers.AddDecision)
  router.GET("/:url", handlers.GetURL)
  router.GET("/campaign/:id", handlers.GetCampaign)
  fmt.Println("Running on 8080")
	router.Run("localhost:8080")   
}


