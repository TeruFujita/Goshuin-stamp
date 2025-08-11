# 御朱印アプリ開発進捗

## 📋 プロジェクト概要

**Goshuin App MVP** - 海外観光客向けデジタル御朱印アプリケーション

### 目標
- 海外観光客が日本の寺社で御朱印をデジタルで収集・管理
- GPSベースの寺社検索機能
- 御朱印のマナーと文化のガイド機能
- 美しいUI/UXで日本文化を体験

## 🏗️ 技術スタック

### フロントエンド
- **Next.js 14** (App Router)
- **TypeScript**
- **Tailwind CSS**
- **Framer Motion** (アニメーション)
- **React Hook Form** (フォーム管理)
- **Zustand** (状態管理)
- **Lucide React** (アイコン)
- **Axios** (HTTP通信)
- **React Hot Toast** (通知)

### バックエンド
- **Go 1.24+** (標準ライブラリのみ使用)
- **MySQL 8.0** (データベース)
- **Ent** (ORM - 手動実装)
- **Docker** (コンテナ化)

### インフラ
- **Docker Compose** (マルチコンテナ環境)
- **AWS S3** (ファイルストレージ予定)

## ✅ 完了した機能

### 1. プロジェクト構造の構築
- [x] ディレクトリ構造の作成
- [x] Docker設定ファイル
- [x] 環境変数設定
- [x] README.mdの作成

### 2. バックエンド開発
- [x] Goプロジェクトの初期化
- [x] 依存関係の管理 (`go.mod`)
- [x] 環境変数設定 (`godotenv`)
- [x] データベース接続設定
- [x] Entクライアントの手動実装
- [x] HTTPサーバーの構築 (標準ライブラリ)
- [x] CORS設定
- [x] APIエンドポイントの実装

#### 実装済みAPI
- `GET /health` - ヘルスチェック
- `GET /api/v1/temples` - 寺社一覧取得
- `GET /api/v1/temples/{id}` - 特定寺社取得
- `GET /api/v1/temples/nearby` - 近隣寺社検索
- `GET /api/v1/goshuin` - 御朱印コレクション一覧
- `POST /api/v1/goshuin` - 御朱印コレクション作成
- `GET /api/v1/goshuin/{id}` - 特定御朱印取得
- `PUT /api/v1/goshuin/{id}` - 御朱印コレクション更新
- `DELETE /api/v1/goshuin/{id}` - 御朱印コレクション削除
- `GET /api/v1/guide` - ガイド情報取得

### 3. データベース設計
- [x] MySQLテーブル設計
- [x] テーブル自動作成機能
- [x] サンプルデータの挿入
- [x] 外部キー制約の設定

#### テーブル構造
```sql
-- temples テーブル
CREATE TABLE temples (
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
);

-- goshuin_collections テーブル
CREATE TABLE goshuin_collections (
    id INT AUTO_INCREMENT PRIMARY KEY,
    temple_id INT NOT NULL,
    image_url VARCHAR(500),
    notes TEXT,
    collected_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (temple_id) REFERENCES temples(id) ON DELETE CASCADE
);
```

### 4. フロントエンド開発
- [x] Next.jsプロジェクトの初期化
- [x] TypeScript設定
- [x] Tailwind CSS設定
- [x] 依存関係のインストール
- [x] 基本的なページ構造
- [x] レスポンシブデザイン
- [x] アニメーション実装

#### 実装済みページ
- **ホームページ** (`/`) - ランディングページ
- **寺社一覧** (`/temples`) - GPSベース検索
- **御朱印コレクション** (`/collection`) - 収集管理
- **ガイド** (`/guide`) - 御朱印マナー

### 5. Docker環境
- [x] MySQLコンテナ設定
- [x] Goバックエンドコンテナ設定
- [x] Next.jsフロントエンドコンテナ設定
- [x] ネットワーク設定
- [x] ボリューム設定

## 🔧 解決した技術的課題

### 1. Entスキーマ生成エラー
**問題**: Go 1.24との互換性問題でEntツールが動作しない
**解決策**: 手動でEntクライアントを実装し、実際のデータベース操作に対応

### 2. データベース接続
**問題**: 簡易的なEntクライアントでは実際のDB操作ができない
**解決策**: MySQLに直接接続し、SQLクエリでCRUD操作を実装

### 3. フレームワーク依存の削除
**問題**: Ginフレームワークの使用
**解決策**: Go標準ライブラリ（net/http）のみでHTTPサーバーを構築

## 📊 現在の状況

### 起動可能なサービス
- ✅ **バックエンド**: http://localhost:8080
- ✅ **フロントエンド**: http://localhost:3000
- ✅ **データベース**: MySQL (Docker)

### 動作確認済み機能
- ✅ APIエンドポイントの応答
- ✅ データベース接続
- ✅ サンプルデータの取得
- ✅ フロントエンドの表示

## 🚀 次のステップ

### 短期目標
1. **フロントエンドとバックエンドの連携**
   - API呼び出しの実装
   - データの表示・更新機能

2. **GPS機能の実装**
   - 現在地取得
   - 近隣寺社検索ロジック

3. **画像アップロード機能**
   - S3連携
   - 御朱印画像の保存

### 中期目標
1. **ユーザー認証システム**
2. **プッシュ通知機能**
3. **オフライン対応**
4. **多言語対応の拡充**

### 長期目標
1. **PWA対応**
2. **ネイティブアプリ化**
3. **ソーシャル機能**
4. **AI機能（寺社推薦など）**

## 📝 開発メモ

### 重要な決定事項
- **フレームワーク不使用**: Go標準ライブラリのみでバックエンド構築
- **手動Ent実装**: 自動生成ツールの代わりに手動でORM実装
- **Docker優先**: 開発環境はDockerで統一

### 学習ポイント
- Go標準ライブラリでのWebサーバー構築
- MySQLとの直接接続とクエリ実行
- Next.js 14のApp Router
- TypeScriptとTailwind CSSの組み合わせ
- Docker Composeでのマルチコンテナ環境

## 📅 開発履歴

- **2025/08/11**: プロジェクト初期化、基本構造構築
- **2025/08/11**: バックエンド開発、Entクライアント実装
- **2025/08/11**: データベース設計、テーブル作成
- **2025/08/11**: フロントエンド開発、UI実装
- **2025/08/11**: Docker環境構築、サービス起動確認

---

**開発者**: AI Assistant + User  
**プロジェクト**: Goshuin App MVP  
**バージョン**: 0.1.0  
**最終更新**: 2025年8月11日
