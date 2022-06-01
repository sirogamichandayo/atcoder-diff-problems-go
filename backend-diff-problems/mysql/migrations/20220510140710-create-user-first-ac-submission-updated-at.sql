-- +migrate Up
CREATE TABLE `user_first_ac_submission_updated_at` (
    `updated_epoch_time` bigint unsigned NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +migrate Down
DROP TABLE IF EXISTS `user_first_ac_submission_updated_at`;
