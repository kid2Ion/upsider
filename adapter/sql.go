package adapter

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// adapter は外部接続を実装します。

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	conn, err := sql.Open("sqlite3", "db/upsider")
	if err != nil {
		panic(err)
	}
	// TODO: sqliteでは文字列全般はtextで扱われるが、postgresqlなどに移行してvarchar(255)のようなカラムへの制限がある方がベター
	cmd := `
		create table if not exists companies (
			uuid text primary key not null,
			name text not null,
			representative_name text,
			phone_number text,
			postal_code text,
			address text
		);
		create table if not exists users (
			uuid text primary key,
			company_uuid text references companies(uuid) on delete cascade,
			name text not null,
			email text unique not null,
			password text not null
		);
		create table if not exists clients (
			uuid text primary key,
			company_uuid text references companies(uuid) on delete cascade,
			name text not null,
			representative_name text,
			phone_number text,
			postal_code text,
			address text
		);
		create table if not exists client_bank_account (
			uuid text primary key,
			client_uuid text references clients(uuid) on delete cascade,
			bank_name text not null,
			branch_name text not null,
			account_number text not null, -- 口座番号は先頭が0の可能性もあるためtextとして扱う
			account_name text not null
		);
		create table if not exists invoices (
			uuid text primary key,
			company_uuid text references companies(uuid),
			client_uuid text references clients(uuid),
			issued_date date not null,
			amount decimal(10, 2) not null,
			fee decimal(10, 2),
			fee_rate decimal(4, 2),
			tax decimal(10, 2),
			tax_rate decimal(4, 2),
			total_amount decimal(10, 2),
			due_date date not null,
			status integer default 0 -- デフォルトは未処理を表す0
		);
	`
	_, err = conn.Exec(cmd)
	if err != nil {
		panic(err)
	}
	return &SqlHandler{Conn: conn}
}
