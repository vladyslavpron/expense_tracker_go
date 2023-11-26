package main

import (
	"context"
	"log"
	"net/http"
	balanceHandler "tracker/core/handlers/balance"
	categoryHandler "tracker/core/handlers/category"
	transactionHandler "tracker/core/handlers/transaction"
	"tracker/core/logger"
	"tracker/ent"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	client, err := ent.Open("sqlite3", "file:db.db?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	r := gin.Default()
	r.GET("/ping", PingHandler)
	r.GET("/test", TestHandler)

	bh := balanceHandler.BalanceHandler{DB: client}
	r.GET("/balances", bh.GetBalances)
	r.POST("/balance", bh.CreateBalance)

	ch := categoryHandler.CategoryHandler{DB: client}
	r.GET("/categories", ch.GetCategories)
	r.POST("/category", ch.CreateCategory)

	th := transactionHandler.TransactionHandler{DB: client}
	r.GET("/transactions", th.GetTransactions)
	r.POST("/transaction", th.CreateTransaction)

	r.Run()
}

func TestHandler(ctx *gin.Context) {
	logger.Log("TEST")

	// if err := ctx.BindJSON(&body); err != nil {
	// 	logger.Err("ERR " + string(err.Error()))
	// 	ctx.JSON(http.StatusBadRequest, body)
	// 	return
	// }
	// logger.Log(fmt.Sprintf("%#v", body))
	ctx.JSON(http.StatusOK, `body`)

}

func PingHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
