ALTER TABLE movies ADD COLUMN store_name LONGTEXT NOT NULL AFTER name;
ALTER TABLE movies ADD COLUMN display_name LONGTEXT NOT NULL AFTER store_name;