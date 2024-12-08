-- name: CreateUser :one
insert into general.users ("name", email, password_hash, password_salt)
values($1, $2, $3, $4)
returning id;

-- name: FindUser :one
select * from general.users where email =$1;