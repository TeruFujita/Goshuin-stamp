package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Load 環境変数を読み込みます
func Load() error {
	// .envファイルが存在する場合は読み込み
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			return err
		}
	}
	return nil
}

// GetDBConfig データベース設定を取得します
func GetDBConfig() map[string]string {
	return map[string]string{
		"host":     getEnv("DB_HOST", "localhost"),
		"port":     getEnv("DB_PORT", "3306"),
		"name":     getEnv("DB_NAME", "stamp_db"),
		"user":     getEnv("DB_USER", "stamp_user"),
		"password": getEnv("DB_PASSWORD", "stamp_password"),
	}
}

// GetJWTSecret JWTシークレットを取得します
func GetJWTSecret() string {
	return getEnv("JWT_SECRET", "default-secret-key")
}

// GetS3Config S3設定を取得します
func GetS3Config() map[string]string {
	return map[string]string{
		"bucket":          getEnv("S3_BUCKET", "stamp-app-uploads"),
		"region":          getEnv("S3_REGION", "ap-northeast-1"),
		"access_key_id":   getEnv("S3_ACCESS_KEY_ID", ""),
		"secret_access_key": getEnv("S3_SECRET_ACCESS_KEY", ""),
	}
}

// getEnv 環境変数を取得し、デフォルト値を設定します
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
