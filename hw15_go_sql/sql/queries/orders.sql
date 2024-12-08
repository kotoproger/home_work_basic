-- name: GetOrdersByUser :many
select * from general.orders where user_id = $1;

-- name: CreateOrder :one
insert into general.orders (user_id, order_date, total_amount)
values ($1, now(), 0)
returning id;

-- name: UpdateOrderAmount :one
update general.orders set total_amount = $2
where id = $1
returning total_amount;

-- name: GetOrderById :one
select * from general.orders where id = $1;
