
-- +migrate Up
CREATE TABLE IF NOT EXISTS `samples`(
  `id` INTEGER UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `user_id` INT UNSIGNED NOT NULL,
  `title` VARCHAR(20) NOT NULL,
  `content` VARCHAR(100) NULL,

  CONSTRAINT `fk_samples_user_id` FOREIGN KEY(`user_id`) REFERENCES `users`(`id`),
  CONSTRAINT `uk_samples_user_id_title` UNIQUE(`user_id`, `title`)
);

-- +migrate Down
DROP TABLE IF EXISTS `samples`;
