-- 初期データベース設定
-- このファイルはPostgreSQLコンテナ起動時に自動実行されます

-- テスト用のテーブル作成（GORMが自動作成するので通常は不要ですが、参考用）
-- CREATE TABLE IF NOT EXISTS users (
--     id SERIAL PRIMARY KEY,
--     name VARCHAR(100) NOT NULL,
--     email VARCHAR(100) UNIQUE NOT NULL,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );

-- 初期データの挿入例
-- INSERT INTO users (name, email) VALUES 
--     ('テストユーザー1', 'test1@example.com'),
--     ('テストユーザー2', 'test2@example.com');

-- データベースの準備完了メッセージ
SELECT 'Database initialized successfully!' as message;
