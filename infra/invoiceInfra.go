package infra

import (
	"fmt"
	"upsider/adapter"
	"upsider/domain/entity"
	"upsider/domain/repository"
)

// infraはデータベースアクセスメソッドの具体を実装します

type (
	invoiceInfra struct {
		adapter.SqlHandler
	}
)

func NewInvoiceRepository(sh adapter.SqlHandler) repository.InvoiceRepository {
	return &invoiceInfra{SqlHandler: sh}
}

func (t *invoiceInfra) Create(*entity.Invoice) error {
	fmt.Println("bunbun hello")
	return nil
}
