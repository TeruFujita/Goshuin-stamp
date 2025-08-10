-- 御朱印アプリ用データベース初期化スクリプト

-- データベースの作成（既に環境変数で作成済み）
-- CREATE DATABASE IF NOT EXISTS stamp_db;

-- 文字セットの設定
SET NAMES utf8mb4;
SET CHARACTER SET utf8mb4;
SET collation_connection = 'utf8mb4_unicode_ci';

-- データベースの選択
USE stamp_db;

-- テーブル作成の前に既存テーブルを削除（開発環境用）
-- DROP TABLE IF EXISTS goshuin_collections;
-- DROP TABLE IF EXISTS temples;
-- DROP TABLE IF EXISTS users;

-- 基本的なテーブル構造（Ent ORMで管理されるため、ここでは最小限）
-- 実際のテーブルはGoのEntスキーマから自動生成されます

-- 開発用のサンプルデータ（オプション）
-- ここに初期データを追加できます
