package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Tabed23/Fun-Application-Projects/tree/main/expense-tracker-app/backend/expense/db"
	"github.com/Tabed23/Fun-Application-Projects/tree/main/expense-tracker-app/backend/expense/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var expDB = db.Connect()

func AddExpense(c *gin.Context){
	var exp model.Expense

	if err := c.ShouldBindJSON(&exp); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg": err.Error(),
		})
		return
	}

	col := expDB.Database("test").Collection("expense")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data := bson.M{
		"title": exp.Title,
		"amount": exp.Amount,
		"type": exp.Type,
		"date": exp.Date,
		"category": exp.Category,
		"description": exp.Description,
		"createdAt": time.Now().Unix(),}

	_, err := col.InsertOne(ctx, data)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "msg": err })
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg": "ok",
	})

}

func GetExpense(c *gin.Context){

	col := expDB.Database("test").Collection("expense")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var exps []model.Expense

	courser, err := col.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "msg": err})
	}

	if err = courser.All(context.TODO(), &exps); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "msg": err})
	}

	fmt.Println(exps)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"expenses": exps,
	})
}


func DeleteExpense(c *gin.Context){

	id := c.Param("id")

	col := expDB.Database("test").Collection("expense")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}

	_, err := col.DeleteOne(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "err": err})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg": "expense deleted",
	})

}
