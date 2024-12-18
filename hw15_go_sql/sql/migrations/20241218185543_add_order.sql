-- +goose Up
-- +goose StatementBegin
insert into general.users (id, name, email, password_hash, password_salt)
values ('0191db0a-4380-20e2-c82e-b2d8ac16da3d', 'john', '0193db0a-4380-20e2-c82e-b2d8ac16da3d@mail.ru', '123','123');

insert into general.products (id, name, price)
values ('0191db0a-4380-20e2-c82e-b2d8ac16da3d', 'p1', 100), ('0191db0a-4381-2ac3-8688-017722d58fda', 'p2', 150);

insert into general.orders (id, user_id, order_date, total_amount)
values ('0191db0a-4380-20e2-c82e-b2d8ac16da3d', '0191db0a-4380-20e2-c82e-b2d8ac16da3d', now(), 250);

insert into general.order_products(order_id, product_id, price)
values('0191db0a-4380-20e2-c82e-b2d8ac16da3d', '0191db0a-4380-20e2-c82e-b2d8ac16da3d', 100), 
('0191db0a-4380-20e2-c82e-b2d8ac16da3d', '0191db0a-4381-2ac3-8688-017722d58fda', 100);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

