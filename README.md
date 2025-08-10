# 御朱印アプリ MVP (Goshuin App MVP)

## プロジェクト概要

日本の寺社を巡る「デジタル文化パスポート」アプリケーション。海外観光客向けに、御朱印の収集と寺社巡りのガイドを提供します。

### コンセプト
- 日本の寺社を巡る「デジタル文化パスポート」
- 公式情報も確認しながら、芸術的な御朱印を集める
- 旅の思い出を記録するガイド＆コレクションアプリ

### ターゲット
- 本物の日本文化に触れたい外国人観光客
- ユニークなお土産や体験を求める人

## 技術スタック

### フロントエンド
- **Next.js 14** (TypeScript)
- **Tailwind CSS** (スタイリング)
- **React Hook Form** (フォーム管理)
- **Zustand** (状態管理)

### バックエンド
- **Go 1.21+** (メイン言語)
- **標準ライブラリ** (net/http) - フレームワーク不使用
- **Ent** (ORM)
- **MySQL 8.0** (データベース)

### インフラ
- **Docker** (コンテナ化)
- **Docker Compose** (ローカル開発環境)
- **AWS S3** (ファイルストレージ)

## プロジェクト構造

```
stamp/
├── frontend/          # Next.js フロントエンド
│   ├── src/
│   │   ├── app/       # App Router
│   │   ├── components/ # React コンポーネント
│   │   ├── lib/       # ユーティリティ関数
│   │   └── types/     # TypeScript型定義
│   ├── public/        # 静的ファイル
│   └── package.json
├── backend/           # Go バックエンド
│   ├── cmd/          # エントリーポイント
│   ├── internal/     # 内部パッケージ
│   ├── pkg/          # 公開パッケージ
│   ├── migrations/   # データベースマイグレーション
│   └── go.mod
├── docker/           # Docker設定
│   ├── frontend/
│   ├── backend/
│   └── mysql/
├── docs/             # ドキュメント
└── docker-compose.yml
```

## 環境構築手順

### 前提条件
- **Docker** と **Docker Compose** がインストールされていること
- **Node.js 18+** がインストールされていること
- **Go 1.21+** がインストールされていること（標準ライブラリのみ使用）

### 1. リポジトリのクローン
```bash
git clone <repository-url>
cd stamp
```

### 2. 環境変数の設定
```bash
# バックエンド用の環境変数
cp backend/.env.example backend/.env

# フロントエンド用の環境変数
cp frontend/.env.example frontend/.env.local
```

### 3. Docker環境の起動
```bash
# データベースとバックエンドを起動
docker-compose up -d mysql backend

# フロントエンドを起動（開発モード）
cd frontend
npm install
npm run dev
```

### 4. データベースのセットアップ
```bash
# バックエンドディレクトリに移動
cd backend

# 依存関係のインストール
go mod download

# Entスキーマの生成
go run -mod=mod entgo.io/ent/cmd/ent generate ./internal/ent/schema

# データベースマイグレーションの実行
go run cmd/server/main.go
```

### 5. アプリケーションの確認
- フロントエンド: http://localhost:3000
- バックエンドAPI: http://localhost:8080
- データベース: localhost:3306

## 開発ガイド

### フロントエンド開発
```bash
cd frontend
npm run dev          # 開発サーバー起動
npm run build        # プロダクションビルド
npm run lint         # リンター実行
```

### バックエンド開発
```bash
cd backend
go run cmd/server/main.go    # サーバー起動
go test ./...                # テスト実行
go mod tidy                  # 依存関係整理
go mod download              # 依存関係ダウンロード
```

### データベース操作
```bash
# Entスキーマの生成
go run -mod=mod entgo.io/ent/cmd/ent generate ./internal/ent/schema

# サーバー起動（スキーマが自動で作成される）
go run cmd/server/main.go
```

## 主要機能

### MVP機能
1. **🌍 多言語対応**: UIとガイドを完全英語化
2. **📖 ガイド機能**: 御朱印のいただき方・マナーを解説
3. **📍 GPS寺社サジェスト**: 現在地から近くの寺社を候補表示
4. **⛩️ 寺社の公式情報**: 公式サイトやSNSへのリンク表示
5. **📸 御朱印の登録**: 写真とメモを追加
6. **🗺️ 地図ビュー**: 集めた場所を日本地図上にマッピング

## 学習ポイント

このプロジェクトを通じて以下の技術を学習できます：

### バックエンド（Go）
- **標準ライブラリ**: net/httpを使用したWebサーバー構築
- **Ent**: ORMとスキーマ管理
- **MySQL**: データベース設計
- **Docker**: コンテナ化

### フロントエンド（Next.js）
- **App Router**: 新しいNext.jsルーティング
- **TypeScript**: 型安全な開発
- **Tailwind CSS**: モダンなスタイリング
- **Framer Motion**: アニメーション

### インフラ
- **Docker Compose**: マルチコンテナ環境
- **環境変数管理**: セキュアな設定管理

## ライセンス

MIT License

## 貢献

プルリクエストやイシューの報告を歓迎します。
# Goshuin-stamp
