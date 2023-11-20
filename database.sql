CREATE TABLE IF NOT EXISTS `user`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NULL DEFAULT NULL,
  `password` varchar(255) NULL DEFAULT NULL,
  `name` varchar(255) NULL DEFAULT NULL,
  `phone` varchar(255) NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB;

INSERT INTO `user` VALUES (1, 'dirga@gmail.com', '$2a$10$TGH1JvcjgszNEbtkzu.EteFsBB21dIJ00mqFoUVdiVlgkAgp3dvBq', 'Dirga Meligo', '85319076822', '2023-02-12 23:53:15', '2023-02-12 23:53:15');

CREATE TABLE IF NOT EXISTS `task` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `title` VARCHAR(255) NOT NULL,
    `description` TEXT,
    `duedate` timestamp NULL DEFAULT NULL,
    `status` ENUM('pending', 'completed') DEFAULT 'pending',
    `user_id` INT,
    FOREIGN KEY (`user_id`) REFERENCES `user`(`id`)
) ENGINE = InnoDB;