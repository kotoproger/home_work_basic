-- name: AddProductToOrder :one
insert into general.order_products (order_id, product_id, price)
values ($1, $2, $3)
returning id;

-- name: RemoveProductFromOrder :exec
delete from general.order_products where id = $1;

-- name: GetOrderProducts :many
select * from general.order_products where order_id = $1;