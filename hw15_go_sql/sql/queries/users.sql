-- name: CreateUser :one
insert into general.users ("name", email, password_hash, password_salt)
values($1, $2, $3, $4)
returning id;

-- name: FindUser :one
select * from general.users where email =$1;

-- name: UpdateUserName :exec
update general.users set name = $2 where id = $1;

-- name: DeleteUser :exec
delete from general.users where id = $1;