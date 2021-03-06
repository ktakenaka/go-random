
-- +migrate Up
CREATE TABLE IF NOT EXISTS users(
  id VARCHAR(32) PRIMARY KEY,
  email VARCHAR(100) NOT NULL,
  google_sub VARCHAR(30) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  INDEX google_sub_index(google_sub),
  UNIQUE KEY unique_google_sub(google_sub)
);

-- +migrate Down
DROP TABLE IF EXISTS users;
