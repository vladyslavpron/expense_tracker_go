package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	balanceHandler "tracker/core/handlers/balance"
	categoryHandler "tracker/core/handlers/category"
	transactionHandler "tracker/core/handlers/transaction"
	viewHandler "tracker/core/handlers/view"
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

	tmpl := template.Must(template.ParseGlob("./core/handlers/view/templates/*"))

	r := gin.Default()

	bh := balanceHandler.BalanceHandler{DB: client}
	ch := categoryHandler.CategoryHandler{DB: client}
	th := transactionHandler.TransactionHandler{DB: client}
	vh := viewHandler.ViewHandler{BalanceHandler: &bh, CategoryHandler: &ch, TransactionHandler: &th, Template: tmpl}

	r.GET("/", vh.Main)
	r.GET("/balance/:id", vh.Balance)

	r.Static("/static", "./static")

	api := r.Group("/api")

	api.GET("/ping", PingHandler)
	api.GET("/test", TestHandler)

	api.GET("/balances", bh.GetBalancesHandler)
	api.POST("/balance", bh.CreateBalance)

	api.GET("/categories", ch.GetCategoriesHandler)
	api.POST("/category", ch.CreateCategory)

	api.GET("/transactions", th.GetTransactionsHandler)
	api.POST("/transaction", th.CreateTransaction)

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
