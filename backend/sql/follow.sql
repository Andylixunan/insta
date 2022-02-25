CREATE TABLE follows
(
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    follower_id INT UNSIGNED NOT NULL,
    followee_id INT UNSIGNED NOT NULL,
    created_at DATETIME(3),
    updated_at DATETIME(3),
    deleted_at DATETIME(3) NULL,
    UNIQUE KEY (follower_id, followee_id) 
) CHARACTER SET utf8mb4;