ALTER TABLE users ADD COLUMN password_remember_token LONGTEXT NULL AFTER token;
ALTER TABLE users ADD COLUMN password_remember_token_at INTEGER UNSIGNED NULL AFTER password_remember_token;