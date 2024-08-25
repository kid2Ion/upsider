package infra

import (
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

func (t *invoiceInfra) Create(req *entity.Invoice) error {
	cmd := `
	insert into invoices (
		uuid, company_uuid, client_uuid, issued_date, amount, fee, fee_rate, tax, tax_rate, total_amount, due_date, status
	) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := t.Conn.Exec(cmd,
		req.UUID,
		req.CompanyUUID,
		req.ClientUUID,
		req.IssuedDate,
		req.Amount,
		req.Fee,
		req.FeeRate.Rate(),
		req.Tax,
		req.TaxRate.Rate(),
		req.TotalAmount,
		req.DueDate,
		req.Status.Int(),
	)
	return err
}
