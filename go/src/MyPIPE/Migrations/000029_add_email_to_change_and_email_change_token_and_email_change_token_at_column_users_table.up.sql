ALTER TABLE users ADD COLUMN email_change_token LONGTEXT NULL AFTER password_remember_token_at;
ALTER TABLE users ADD COLUMN email_change_token_at INTEGER UNSIGNED NULL AFTER email_change_token;
ALTER TABLE users ADD COLUMN email_to_change LONGTEXT NULL AFTER email_change_token_at;