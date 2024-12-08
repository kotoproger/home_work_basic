-- +goose Up
-- +goose StatementBegin
create table general.order_products(
	id uuid  default general.new_uuid(),
	order_id uuid NOT NULL,
	product_id uuid NOT NULL,
	price integer not null,
	CONSTRAINT order_products_pk PRIMARY KEY (id)
);
comment on column general.order_products.price is 'в копейках';
CREATE INDEX order_products_order_id_idx ON general.order_products (order_id);
CREATE INDEX order_products_product_id_idx ON general.order_products (product_id);
ALTER TABLE general.order_products ADD CONSTRAINT order_products_order_fk FOREIGN KEY (order_id) REFERENCES general.orders(id);
ALTER TABLE general.order_products ADD CONSTRAINT order_products_product_fk FOREIGN KEY (product_id) REFERENCES general.products(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table general.order_products;
-- +goose StatementEnd
