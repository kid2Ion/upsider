package usecase

import (
	"strconv"
	"upsider/domain/entity"
	"upsider/domain/repository"
)

type (
	invoiceUsecase struct {
		repo repository.InvoiceRepository
	}
	InvoiceUsecase interface {
		Create(req *InvoiceCreateReq) error
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
