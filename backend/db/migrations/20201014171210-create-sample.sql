
-- +migrate Up
CREATE TABLE IF NOT EXISTS `samples`(
  `id` VARCHAR(32) PRIMARY KEY,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `user_id` VARCHAR(32) NOT NULL,
  `title` VARCHAR(20) NOT NULL,
  `content` VARCHAR(100) NULL,

  FULLTEXT INDEX `ft_samples`(`title`, `content`) WITH PARSER ngram,
  CONSTRAINT `fk_samples_user_id` FOREIGN KEY(`user_id`) REFERENCES `users`(`id`),
  CONSTRAINT `uk_samples_user_id_title` UNIQUE(`user_id`, `title`)
);

-- +migrate Down
DROP TABLE IF EXISTS `samples`;
