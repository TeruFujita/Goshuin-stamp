package database

import (
	"context"
	"fmt"
	"log"

	"stamp-backend/internal/config"
	"stamp-backend/internal/ent"

	_ "github.com/go-sql-driver/mysql"
)

// Init データベース接続を初期化します
func Init() (*ent.Client, error) {
	dbConfig := config.GetDBConfig()
	
	// DSN (Data Source Name) の構築
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig["user"],
		dbConfig["password"],
		dbConfig["host"],
		dbConfig["port"],
		dbConfig["name"],
	)

	// Entクライアントの作成
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to mysql: %v", err)
	}

	// 接続テスト
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Printf("failed creating schema resources: %v", err)
	}

	log.Println("Database connection established successfully")
	return client, nil
}
