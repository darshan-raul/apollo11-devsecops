package routes

import (
	"github.com/darshan-raul/ecommerce/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("users/signup", controllers.Singup())
	incomingRoutes.POST("users/login",controllers.Login())
	incomingRoutes.POST("users/addprduct",controllers.ProductViewAdmin())
	incomingRoutes.POST("users/productview",controllers.SearchProduct())
	incomingRoutes.POST("users/search",controllers.SearchProduct())
}