-- +goose Up
-- +goose StatementBegin
insert into general.products (name, price)
values ('tea', 100), ('coffe', 100), ('sugar', 150);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
