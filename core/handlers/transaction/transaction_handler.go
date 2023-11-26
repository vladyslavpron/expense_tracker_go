package transactionHandler

import "tracker/ent"

type TransactionHandler struct {
	DB *ent.Client
}
