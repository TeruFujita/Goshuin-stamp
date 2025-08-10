package main

import (
	"log"
	"os"

	"stamp-backend/internal/config"
	"stamp-backend/internal/database"
	"stamp-backend/internal/server"
)

func main() {
	// 環境変数の読み込み
	if err := config.Load(); err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// データベース接続の初期化
	db, err := database.Init()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// サーバーの初期化と起動
	srv := server.New(db)
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := srv.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
