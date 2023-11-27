package main

import (
	"context"
	"html/template"
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

	tmpl := template.Must(template.ParseGlob("./templates/*"))

	r := gin.Default()

	bh := balanceHandler.BalanceHandler{DB: client}
	ch := categoryHandler.CategoryHandler{DB: client}
	th := transactionHandler.TransactionHandler{DB: client}

	r.GET("/", func(ctx *gin.Context) {
		b, _ := bh.GetBalances(ctx)
		categories, _ := ch.GetCategories(ctx)
		t, _ := th.GetTransactions(ctx, transactionHandler.GetTransactionsParams{})
		categoryTitles := make(map[int]string, len(b))
		for _, category := range categories {
			categoryTitles[category.ID] = category.Title
		}

		data := MainTemplateData{
			Balances:       b,
			Categories:     categories,
			Transactions:   t,
			CategoryTitles: categoryTitles,
		}

		tmpl.ExecuteTemplate(ctx.Writer, "main", data)
	})

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

type MainTemplateData struct {
	Balances       []*ent.Balance
	Categories     []*ent.Category
	Transactions   []*ent.Transaction
	CategoryTitles map[int]string
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
