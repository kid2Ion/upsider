package di

import (
	"upsider/adapter"
	"upsider/domain/repository"
	"upsider/handler"
	"upsider/infra"
	"upsider/usecase"
)

// di は依存性の注入を行います。

// TODO: DIコンテナー等に置き換えて、拡張性高くする。
// db
func injectDB() adapter.SqlHandler {
	sh := adapter.NewSqlHandler()
	return *sh
}

// repository
func injectInvoiceRepository() repository.InvoiceRepository {
	sh := injectDB()
	return infra.NewInvoiceRepository(sh)
}

// usecase
func injectInvoiceUsecase() usecase.InvoiceUsecase {
	ir := injectInvoiceRepository()
	return usecase.NewInvoiceUsecase(ir)
}

// handler
func InjectInvoiceHandler() handler.InvoiceHandler {
	iu := injectInvoiceUsecase()
	return handler.NewInvoiceHandler(iu)
}
