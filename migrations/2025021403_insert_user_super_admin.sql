-- +goose Up
-- +goose StatementBegin
-- insurance_cofi.partner definition

insert into `master_user` (`username`,`password`,`role`,`status`, `created_by`) value ('antoganteng','785f60702a692d4b5e32:$2a$10$N3EH8ms.b80Ss0siqZjxbOIs2sYcm4ABPgW4C4y3N33HEn8V19972','849c9eee-e30f-4dc5-9816-9b395b0121b7','active','system');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `master_user`;
-- +goose StatementEnd