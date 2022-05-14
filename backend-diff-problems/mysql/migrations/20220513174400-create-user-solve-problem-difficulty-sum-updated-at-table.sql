-- +migrate Up
CREATE TABLE `user_solve_problem_difficulty_sum_updated_at` (
    `updated_epoch_time` bigint unsigned NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +migrate Down
DROP TABLE IF EXISTS `user_solve_problem_difficulty_sum_updated_at`;
