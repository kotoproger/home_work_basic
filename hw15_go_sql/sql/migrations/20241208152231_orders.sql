-- +goose Up
-- +goose StatementBegin
CREATE TABLE general.orders (
	id bigserial NOT NULL,
	user_id bigint NOT NULL,
	order_date timestamp NOT NULL,
	total_amount integer NOT null,
	CONSTRAINT orders_pk PRIMARY KEY (id)
);
comment on column general.orders.total_amount is 'в копейках';
CREATE INDEX orders_user_id_idx ON general.orders (user_id);
ALTER TABLE general.orders ADD CONSTRAINT orders_users_fk FOREIGN KEY (user_id) REFERENCES general.users(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table general.orders;
-- +goose StatementEnd
