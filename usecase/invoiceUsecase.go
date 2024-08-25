package usecase

import (
	"upsider/domain/entity"
	"upsider/domain/repository"
)

type (
	invoiceUsecase struct {
		repo repository.InvoiceRepository
	}
	InvoiceUsecase interface {
		Create() error
	}
)

func NewInvoiceUsecase(repo repository.InvoiceRepository) InvoiceUsecase {
	return &invoiceUsecase{repo: repo}
}

func (t *invoiceUsecase) Create() error {
	return t.repo.Create(&entity.Invoice{})
}
