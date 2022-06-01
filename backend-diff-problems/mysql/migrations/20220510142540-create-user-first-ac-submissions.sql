-- +migrate Up
CREATE TABLE `user_first_ac_submissions` (
    `user_id` varchar(20) NOT NULL,
    `problem_id` varchar(100) NOT NULL,
    `first_solved_epoch_time` bigint NOT NULL,
    PRIMARY KEY (`user_id`, `problem_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +migrate Down
DROP TABLE IF EXISTS `user_first_ac_submissions`;
