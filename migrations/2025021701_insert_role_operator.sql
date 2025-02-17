-- +goose Up
-- +goose StatementBegin
-- insurance_cofi.partner definition

insert into `master_role` (`name`,`role_uuid`,`status`, `created_by`) value ('operator','6137715f-0895-4f03-9a4a-daf46acefd99','active','system');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `master_role`;
-- +goose StatementEnd