package database

import (
	"database/sql"
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

	// 実際のMySQL接続
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	// 接続テスト
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	// テーブル作成
	if err := createTables(db); err != nil {
		return nil, fmt.Errorf("failed to create tables: %v", err)
	}

	// Entクライアントの作成（実際のDB接続付き）
	client := ent.NewClientWithDB(db)

	log.Println("Database connection established successfully")
	return client, nil
}

// createTables 必要なテーブルを作成します
func createTables(db *sql.DB) error {
	// templesテーブル
	templeSQL := `
	CREATE TABLE IF NOT EXISTS temples (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		name_en VARCHAR(255) NOT NULL,
		description TEXT,
		description_en TEXT,
		latitude DOUBLE NOT NULL,
		longitude DOUBLE NOT NULL,
		address VARCHAR(500),
		phone VARCHAR(50),
		website VARCHAR(500),
		instagram VARCHAR(255),
		twitter VARCHAR(255),
		opening_hours VARCHAR(255),
		goshuin_fee VARCHAR(100),
		goshuin_office VARCHAR(255),
		is_active BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
	`

	// goshuin_collectionsテーブル
	goshuinSQL := `
	CREATE TABLE IF NOT EXISTS goshuin_collections (
		id INT AUTO_INCREMENT PRIMARY KEY,
		temple_id INT NOT NULL,
		image_url VARCHAR(500),
		notes TEXT,
		collected_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (temple_id) REFERENCES temples(id) ON DELETE CASCADE
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
	`

	// テーブル作成実行
	if _, err := db.Exec(templeSQL); err != nil {
		return fmt.Errorf("failed to create temples table: %v", err)
	}

	if _, err := db.Exec(goshuinSQL); err != nil {
		return fmt.Errorf("failed to create goshuin_collections table: %v", err)
	}

	// サンプルデータの挿入
	if err := insertSampleData(db); err != nil {
		log.Printf("Warning: failed to insert sample data: %v", err)
	}

	return nil
}

// insertSampleData サンプルデータを挿入します
func insertSampleData(db *sql.DB) error {
	// 既存データをチェック
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM temples").Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return nil // 既にデータがある場合はスキップ
	}

	// サンプル寺社データ
	sampleTemples := []struct {
		name, nameEn, description string
		lat, lng                  float64
	}{
		{
			name:        "浅草寺",
			nameEn:      "Senso-ji Temple",
			description: "東京最古の寺院で、雷門と五重塔が有名",
			lat:         35.7148,
			lng:         139.7967,
		},
		{
			name:        "明治神宮",
			nameEn:      "Meiji Shrine",
			description: "明治天皇と昭憲皇太后を祀る神社",
			lat:         35.6764,
			lng:         139.6993,
		},
		{
			name:        "金閣寺",
			nameEn:      "Kinkaku-ji",
			description: "京都の有名な禅寺、金箔で覆われた建物",
			lat:         35.0394,
			lng:         135.7292,
		},
	}

	for _, temple := range sampleTemples {
		_, err := db.Exec(`
			INSERT INTO temples (name, name_en, description, latitude, longitude, is_active)
			VALUES (?, ?, ?, ?, ?, ?)
		`, temple.name, temple.nameEn, temple.description, temple.lat, temple.lng, true)

		if err != nil {
			return fmt.Errorf("failed to insert temple %s: %v", temple.name, err)
		}
	}

	log.Println("Sample data inserted successfully")
	return nil
}
