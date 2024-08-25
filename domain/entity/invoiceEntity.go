package entity

import (
	"errors"
	"time"
	"upsider/domain/vo"

	"github.com/google/uuid"
)

// Invoice は請求書を表現するエンティティです。

type Invoice struct {
	UUID        string           // UUID は請求書を一意に識別するためのIDです
	CompanyUUID string           // CompanyUUID は請求書を発行した企業を識別するためのIDです
	ClientUUID  string           // ClientUUID は請求書の宛先である取引先を識別するためのIDです
	IssuedDate  time.Time        // IdduedDate は請求書が発行された日付です
	Amount      float64          // Amount は請求金額です
	Fee         float64          // Fee は請求金額に対する手数料です
	FeeRate     vo.FeeRate       // FeeRate は手数料率です
	Tax         float64          // Tax は手数料に対する消費税です
	TaxRate     vo.TaxRate       // TaxRate は消費税率です
	TotalAmount float64          // TotalAmount は請求金額、手数料、消費税を含めた合計金額です
	DueDate     time.Time        // DueDate は請求書の支払期日です
	Status      vo.InvoiceStatus // Status は請求書の現在のステータスです
}

func NewInvoice(
	companyUUID string,
	clientUUID string,
	amount float64,
	dueDateStr string,
) (*Invoice, error) {
	if companyUUID == "" {
		return nil, errors.New("CompanyUUID は必須です")
	}
	if clientUUID == "" {
		return nil, errors.New("ClientUUID は必須です")
	}
	if amount <= 0 {
		return nil, errors.New("請求金額は1以上である必要があります")
	}
	dueDate, err := time.Parse("2006-01-02", dueDateStr)
	if err != nil {
		return nil, errors.New("支払期日のフォーマットが不正です")
	}
	// 支払期日は23:59:59まで
	dueDate = dueDate.Add(time.Hour*23 + time.Minute*59 + time.Second*59)

	// 発行日初期値は現在時刻
	issuedDate := time.Now().UTC()
	feeRate := vo.DefaultFeeRate
	taxRate := vo.DefaultTaxRate
	fee := amount * feeRate.Rate()
	tax := fee * (taxRate.Rate() - 1)
	totalAmount := amount + fee + tax
	// 初期ステータスは「未処理」
	status := vo.InvoiceStatusUnprocessed

	if issuedDate.After(dueDate) {
		return nil, errors.New("支払期日が過ぎた請求書は作成できません")
	}

	return &Invoice{
		UUID:        uuid.New().String(),
		CompanyUUID: companyUUID,
		ClientUUID:  clientUUID,
		IssuedDate:  issuedDate,
		Amount:      amount,
		Fee:         fee,
		FeeRate:     feeRate,
		Tax:         tax,
		TaxRate:     taxRate,
		TotalAmount: totalAmount,
		DueDate:     dueDate,
		Status:      status,
	}, nil
}
