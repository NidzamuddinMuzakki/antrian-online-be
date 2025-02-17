-- +goose Up
-- +goose StatementBegin
-- insurance_cofi.partner definition

CREATE TABLE IF NOT EXISTS `master_antrian`
(
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `number`       int DEFAULT 0 NOT NULL ,
    `tipe_pasien_id`       int DEFAULT 0 NOT NULL ,
    `loket_id`       int DEFAULT 0 NOT NULL ,
    `status`       TEXT DEFAULT '' NOT NULL ,
    `created_at` timestamp    NULL DEFAULT CURRENT_TIMESTAMP,
    `created_by` varchar(200)     DEFAULT '' NOT NULL,
    `updated_at` timestamp    NULL,
    `updated_by` varchar(200)  DEFAULT ''  NOT NULL ,
    PRIMARY KEY (`id`),
    UNIQUE KEY `master_antrian` (`number`,`created_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `master_antrian`;
-- +goose StatementEnd