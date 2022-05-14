-- +migrate Up
CREATE TABLE `user_solve_problem_difficulty_sum` (
    `user_id` VARCHAR(20) NOT NULL,
    `clip_difficulty_sum` double NOT NULL,
    PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +migrate Down
DROP TABLE IF EXISTS `user_solve_problem_difficulty_sum`;
