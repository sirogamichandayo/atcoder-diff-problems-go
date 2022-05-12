-- +migrate Up
CREATE TABLE `problem_difficulties` (
    `problem_id` varchar(100) NOT NULL,
    `difficulty` DOUBLE,
    `clip_difficulty` DOUBLE,
    PRIMARY KEY (`problem_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +migrate Down
DROP TABLE IF EXISTS `problem_difficulties`;
