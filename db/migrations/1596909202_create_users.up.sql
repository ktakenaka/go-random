CREATE TABLE IF NOT EXISTS users(
  id INTEGER UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  email VARCHAR(255) NOT NULL,
  google_sub VARCHAR(30) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  INDEX google_sub_index(google_sub),
  UNIQUE KEY unique_google_sub(google_sub)
);
