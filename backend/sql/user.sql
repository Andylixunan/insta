CREATE TABLE users
(
    id         INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    username   VARCHAR(20) NOT NULL UNIQUE,
    password   VARCHAR(100) NOT NULL,
    created_at DATETIME(3),
    updated_at DATETIME(3),
    deleted_at DATETIME(3) NULL
) CHARACTER SET utf8mb4;