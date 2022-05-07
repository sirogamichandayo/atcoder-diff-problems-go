-- +migrate Up
CREATE TABLE `product_difficulties` (
    `problem_id` varchar(100),
    `difficulty` mediumint,
    PRIMARY KEY (`problem_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +migrate Down
DROP TABLE IF EXIST `product_difficulties`;
