CREATE TABLE users
(
    id         INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    username   VARCHAR(255) NOT NULL UNIQUE,
    password   VARCHAR(255) NOT NULL,
    nickname   VARCHAR(255) NOT NULL DEFAULT 'no nickname',
    self_description VARCHAR(255) NOT NULL DEFAULT 'no self_description',
    avatar VARCHAR(255) COMMENT 'URL of the avatar',
    created_at DATETIME(3),
    updated_at DATETIME(3),
    deleted_at DATETIME(3) NULL
) CHARACTER SET utf8mb4;