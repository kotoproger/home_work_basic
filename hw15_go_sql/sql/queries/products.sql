
-- name: GetProductById :one
select * from general.products where id = $1;

-- name: CreateProduct :one
insert into general.products ("name", price)
values($1, $2)
returning id;

-- name: GetProducts :many
select * from general.products;

-- name: UpdateProductPrice :exec
update general.products set price = $2 where id = $1;

-- name: DeleteProduct :exec
delete from general.products where id = $1;