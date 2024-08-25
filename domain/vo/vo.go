package vo

// voはvalue objectを管理します。

type (
	// FeeRate は手数料率です。
	FeeRate float64
	// TaxRate は消費税率です。
	TaxRate float64
	// InvoiceStatus は請求書のステータスです。
	InvoiceStatus int
)

const (
	DefaultFeeRate FeeRate = 0.04
	DefaultTaxRate TaxRate = 1.10
	// InvoiceStatusUnprocessed は「未処理」です。
	InvoiceStatusUnprocessed InvoiceStatus = 0
	// InvoiceStatusProcessing は「処理中」です。
	InvoiceStatusProcessing InvoiceStatus = 1
	// InvoiceStatusPaid は「支払い済み」です。
	InvoiceStatusPaid InvoiceStatus = 2
	// InvoiceStatusError は「エラー」です。
	InvoiceStatusError InvoiceStatus = 3
)

func (t FeeRate) Rate() float64 {
	return float64(t)
}

func (t TaxRate) Rate() float64 {
	return float64(t)
}

func (t InvoiceStatus) String() string {
	switch t {
	case InvoiceStatusUnprocessed:
		return "未処理"
	case InvoiceStatusProcessing:
		return "処理中"
	case InvoiceStatusPaid:
		return "支払い済み"
	case InvoiceStatusError:
		return "エラー"
	default:
		return "不明"
	}
}

func (t InvoiceStatus) Int() int {
	switch t {
	case InvoiceStatusUnprocessed:
		return 0
	case InvoiceStatusProcessing:
		return 1
	case InvoiceStatusPaid:
		return 2
	case InvoiceStatusError:
		return 3
	default:
		return 4
	}
}
