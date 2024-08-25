package usecase

import (
	"errors"
	"strconv"
	"time"
	"upsider/domain/entity"
	"upsider/domain/repository"
)

type (
	invoiceUsecase struct {
		repo repository.InvoiceRepository
	}
	InvoiceUsecase interface {
		Create(req *InvoiceCreateReq) error
		Get(fromDateStr string, toDateStr string) ([]*entity.Invoice, error)
	}

	InvoiceCreateReq struct {
		CompanyUUID string `json:"company_uuid"`
		ClientUUID  string `json:"client_uuid"`
		Amount      string `json:"amount"`
		DueDate     string `json:"due_date"`
	}
)

func NewInvoiceUsecase(repo repository.InvoiceRepository) InvoiceUsecase {
	return &invoiceUsecase{repo: repo}
}

func (t *invoiceUsecase) Create(req *InvoiceCreateReq) error {
	amount, err := strconv.Atoi(req.Amount)
	if err != nil {
		return err
	}
	invoice, err := entity.NewInvoice(
		req.CompanyUUID,
		req.ClientUUID,
		float64(amount),
		req.DueDate,
	)
	if err != nil {
		return err
	}
	return t.repo.Create(invoice)
}

func (t *invoiceUsecase) Get(fromDateStr string, toDateStr string) ([]*entity.Invoice, error) {
	fromDate, err := time.Parse("2006-01-02", fromDateStr)
	if err != nil {
		return nil, errors.New("from_dateが不正なフォーマットです")
	}
	toDate, err := time.Parse("2006-01-02", toDateStr)
	if err != nil {
		return nil, errors.New("to_dateが不正なフォーマットです")
	}
	invoices, err := t.repo.FindByDateRange(fromDate, toDate)
	if err != nil {
		return nil, err
	}
	return invoices, nil
}
