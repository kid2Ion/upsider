package repository

import (
	"time"
	"upsider/domain/entity"
)

// クリーンアーキテクチャの依存性逆転に則るためにrepositoryではinterface定義のみにする
// 複雑なドメイン固有のロジックを持ちたい場合は、service層などで実装する(今回のcaseでは不要)
type (
	// InvoiceRepository は請求書データ操作を抽象化します。
	InvoiceRepository interface {
		Create(req *entity.Invoice) error
		FindByDateRange(fromDate time.Time, toDate time.Time) ([]*entity.Invoice, error)
	}
)
