-- +goose Up
-- +goose StatementBegin
create table general.products (
	id uuid default general.new_uuid(),
	"name" varchar not null,
	price integer not null,
	CONSTRAINT products_pk PRIMARY KEY (id)
);
comment on column general.products.price is 'в копейках';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table general.products
-- +goose StatementEnd
