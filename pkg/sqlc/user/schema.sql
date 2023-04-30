CREATE TABLE users (
  `id` char(36) NOT NULL,
  `username` varchar(200) NOT NULL,
	`password` binary(32) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB CHARSET utf8mb4 COLLATE utf8mb4_bin;

ALTER TABLE `users` ADD UNIQUE INDEX `user_username` (`usernmae`);

