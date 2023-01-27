package main

import (
	"/controllers"
	"/database"
	"/middleware"
	"/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)
func main(){
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := controllers.NewApplication(database.ProductData(database.Client,"Product"),database.UserData(database.Client))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	routes.Use(middleware.Authentication())
	
	router.GET("/addtocart",app.AddToCart())
	router.GET("/romoveitem",app.RemoveItem())
	router.GET("/cartcheckout",app.ByFromCart())
	router.GET("/instantbuy",app.InstanBuy())

	log.Fatal(router.Run((":" + port)))

}