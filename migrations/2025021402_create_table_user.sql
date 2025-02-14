-- +goose Up
-- +goose StatementBegin
-- insurance_cofi.partner definition

CREATE TABLE IF NOT EXISTS `master_user`
(
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `username`       varchar(200)  DEFAULT '' NOT NULL ,
     `password`       varchar(900)  DEFAULT '' NOT NULL ,
    `role`       TEXT DEFAULT '' NOT NULL ,
    `status`       TEXT DEFAULT '' NOT NULL ,
    `created_at` timestamp    NULL DEFAULT CURRENT_TIMESTAMP,
    `created_by` varchar(200)     DEFAULT '' NOT NULL,
    `updated_at` timestamp    NULL,
    `updated_by` varchar(200)  DEFAULT ''  NOT NULL ,
    PRIMARY KEY (`id`),
    UNIQUE KEY `master_user` (`username`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `master_user`;
-- +goose StatementEnd