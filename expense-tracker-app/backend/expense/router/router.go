package router

import (
	"net/http"

	"github.com/Tabed23/Fun-Application-Projects/tree/main/expense-tracker-app/backend/expense/controller"
	"github.com/gin-gonic/gin"
)



const (
	port = ":8081"
)




func StartApplication(){
	r := gin.Default()
	r.GET("/api/v1/", func (ctx * gin.Context)  {
		ctx.String(http.StatusOK, "API is Alive")
	})
	v1 := r.Group("/api/v1/")
	{
		v1.GET("/get-expenses", controller.GetExpense)
		v1.DELETE("delete-expense/:id", controller.DeleteExpense)
		v1.POST("/add-expense/", controller.AddExpense)
	}
	r.Run(port)
}
