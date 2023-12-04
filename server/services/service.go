package services

import (
	"database/sql"
)

type MyAppService struct {
	db *sql.DB
}

// コンストラクタ関数
func NewMyAppService(db *sql.DB) *MyAppService {
	return &MyAppService{db: db}
}
