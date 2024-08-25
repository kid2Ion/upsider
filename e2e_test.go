package main

// シンプルなe2eテスト
import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"upsider/adapter"
	"upsider/router"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func clearDatabase(db *sql.DB) error {
	_, err := db.Exec("delete from invoices")
	return err
}

func TestInvoice(t *testing.T) {
	e := echo.New()
	router.InitInvoiceRouter(e)

	sqlHandler := adapter.NewSqlHandler()
	defer sqlHandler.Conn.Close()

	t.Run("正常系: 一つの請求書データを作成し、取得できる", func(t *testing.T) {
		// testケースが干渉しないようケースごとにレコード削除
		if err := clearDatabase(sqlHandler.Conn); err != nil {
			panic(err)
		}

		// requestParam
		createReq := map[string]interface{}{
			"company_uuid": "company-uuid-1",
			"client_uuid":  "client-uuid-1",
			"amount":       "10000",
			"due_date":     "2025-06-10",
		}
		createBody, _ := json.Marshal(createReq)
		req := httptest.NewRequest(http.MethodPost, "/api/invoices", bytes.NewReader(createBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		// postリクエスト
		e.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)

		// getリクエスト
		req = httptest.NewRequest(http.MethodGet, "/api/invoices?from_date=2025-06-01&to_date=2025-06-29", nil)
		rec = httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)

		// レスポンス検証
		var got []map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &got)

		if assert.Len(t, got, 1) {
			assert.Equal(t, "company-uuid-1", got[0]["CompanyUUID"])
			assert.Equal(t, "client-uuid-1", got[0]["ClientUUID"])
			assert.Equal(t, float64(10000), got[0]["Amount"])
		} else {
			t.Fatal("Expected one invoice but got none")
		}
	})
	t.Run("正常系: 複数の請求書データを作成し、期間内のもののみ取得できる", func(t *testing.T) {
		// testケースが干渉しないようケースごとにレコード削除
		if err := clearDatabase(sqlHandler.Conn); err != nil {
			panic(err)
		}

		// 請求書データを3つ作成
		invoices := []map[string]interface{}{
			{
				"company_uuid": "company-uuid-1",
				"client_uuid":  "client-uuid-1",
				"amount":       "10000",
				"due_date":     "2025-06-10", // 期間内
			},
			{
				"company_uuid": "company-uuid-2",
				"client_uuid":  "client-uuid-2",
				"amount":       "20000",
				"due_date":     "2025-06-15", // 期間内
			},
			{
				"company_uuid": "company-uuid-3",
				"client_uuid":  "client-uuid-3",
				"amount":       "30000",
				"due_date":     "2025-07-01", // 期間外
			},
		}

		for _, invoice := range invoices {
			createBody, _ := json.Marshal(invoice)
			req := httptest.NewRequest(http.MethodPost, "/api/invoices", bytes.NewReader(createBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			// postリクエスト
			e.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusOK, rec.Code)
		}

		// getリクエスト: 期間内のデータのみ取得
		req := httptest.NewRequest(http.MethodGet, "/api/invoices?from_date=2025-06-01&to_date=2025-06-29", nil)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)

		// レスポンス検証
		var got []map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &got)

		if assert.Len(t, got, 2) { // 期待通り2件取得できているか
			assert.Equal(t, "company-uuid-1", got[0]["CompanyUUID"])
			assert.Equal(t, "client-uuid-1", got[0]["ClientUUID"])
			assert.Equal(t, float64(10000), got[0]["Amount"])

			assert.Equal(t, "company-uuid-2", got[1]["CompanyUUID"])
			assert.Equal(t, "client-uuid-2", got[1]["ClientUUID"])
			assert.Equal(t, float64(20000), got[1]["Amount"])
		} else {
			t.Fatal("Expected two invoices but got none or incorrect number")
		}
	})
}
