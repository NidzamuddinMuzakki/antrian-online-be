-- +goose Up
-- +goose StatementBegin
-- insurance_cofi.partner definition

CREATE TABLE IF NOT EXISTS `master_tipe_pasien`
(
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `name`       varchar(200)  DEFAULT '' NOT NULL ,
    `status`       TEXT DEFAULT '' NOT NULL ,
    `created_at` timestamp    NULL DEFAULT CURRENT_TIMESTAMP,
    `created_by` varchar(200)     DEFAULT '' NOT NULL,
    `updated_at` timestamp    NULL,
    `updated_by` varchar(200)  DEFAULT ''  NOT NULL ,
    PRIMARY KEY (`id`),
    UNIQUE KEY `master_tipe_pasien` (`name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `master_tipe_pasien`;
-- +goose StatementEnd