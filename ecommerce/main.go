package main

import(

	"github.com/darshan-raul/ecommerce/controllers"
	"github.com/darshan-raul/ecommerce/database"
	"github.com/darshan-raul/ecommerce/middleware"
	"github.com/darshan-raul/ecommerce/routes"
	"github.com/gin-gonic/gin"
	"os"
	"log"
)

func main(){

	
	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"),database.UserData(database.Client,"Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication)

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.ButFromCart())
	router.GET("/instantbuy",app.InstantBuy())

	log.Fatal(router.Run(":" + port))


}	