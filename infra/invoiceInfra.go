package infra

import (
	"time"
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

func (t *invoiceInfra) FindByDateRange(fromDate time.Time, toDate time.Time) ([]*entity.Invoice, error) {
	q := `
	select uuid, company_uuid, client_uuid, issued_date, amount, fee, fee_rate, tax, tax_rate, total_amount, due_date, status
	from invoices
	where due_date between ? and ?
	`
	rows, err := t.Conn.Query(q, fromDate, toDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*entity.Invoice
	for rows.Next() {
		var invoice entity.Invoice
		err := rows.Scan(
			&invoice.UUID,
			&invoice.CompanyUUID,
			&invoice.ClientUUID,
			&invoice.IssuedDate,
			&invoice.Amount,
			&invoice.Fee,
			&invoice.FeeRate,
			&invoice.Tax,
			&invoice.TaxRate,
			&invoice.TotalAmount,
			&invoice.DueDate,
			&invoice.Status,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, &invoice)
	}
	return res, nil
}
