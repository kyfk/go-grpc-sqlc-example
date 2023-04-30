CREATE TABLE todos (
  `id` char(36) NOT NULL,
  `user_id` char(36) NOT NULL,
  `content` text NOT NULL,
  `due_to` timestamp,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB CHARSET utf8mb4 COLLATE utf8mb4_bin;

