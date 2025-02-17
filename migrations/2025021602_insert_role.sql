-- +goose Up
-- +goose StatementBegin
-- insurance_cofi.partner definition

insert into `master_role` (`name`,`role_uuid`,`status`, `created_by`) value ('superadmin','849c9eee-e30f-4dc5-9816-9b395b0121b7','active','system');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `master_role`;
-- +goose StatementEnd