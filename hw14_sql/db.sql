/* структура бд */
create database hw14;
CREATE SCHEMA IF NOT EXISTS hw14;

CREATE TABLE hw14.users (
	id bigserial NOT NULL,
	"name" varchar NULL,
	email varchar NOT NULL,
	password_hash varchar NOT NULL,
	password_salt varchar NOT NULL,
	CONSTRAINT users_pk PRIMARY KEY (id),
	CONSTRAINT users_unique_email UNIQUE (email)
);

CREATE TABLE hw14.orders (
	id bigserial NOT NULL,
	user_id bigint NOT NULL,
	order_date timestamp NOT NULL,
	total_amount integer NOT null,
	CONSTRAINT orders_pk PRIMARY KEY (id)
);
comment on column hw14.orders.total_amount is 'в копейках';
CREATE INDEX orders_user_id_idx ON hw14.orders (user_id);
ALTER TABLE hw14.orders ADD CONSTRAINT orders_users_fk FOREIGN KEY (user_id) REFERENCES hw14.users(id);

create table hw14.products (
	id bigserial not null,
	"name" varchar not null,
	price integer not null,
	CONSTRAINT products_pk PRIMARY KEY (id)
);
comment on column hw14.products.price is 'в копейках';

create table hw14.order_products(
	id bigserial not null,
	order_id bigint NOT NULL,
	product_id bigint NOT NULL,
	price integer not null,
	CONSTRAINT order_products_pk PRIMARY KEY (id)
);
comment on column hw14.order_products.price is 'в копейках';
CREATE INDEX order_products_order_id_idx ON hw14.order_products (order_id);
CREATE INDEX order_products_product_id_idx ON hw14.order_products (product_id);
ALTER TABLE hw14.order_products ADD CONSTRAINT order_products_order_fk FOREIGN KEY (order_id) REFERENCES hw14.orders(id);
ALTER TABLE hw14.order_products ADD CONSTRAINT order_products_product_fk FOREIGN KEY (product_id) REFERENCES hw14.products(id);

/* функция для создания заказа */
CREATE OR REPLACE FUNCTION hw14.create_order(inemail varchar, inproducts bigint[]) returns bigint as $$
DECLARE 
 new_order_id bigint;
 new_product_id bigint;
 new_total_amount int;
 new_user_id bigint;
begin
	
	if not exists (select id from hw14.users where email = inemail) then
		return 0;
	end if;

	insert into hw14.orders (user_id, order_date, total_amount) (select id, now(), 0 from hw14.users where email = inemail)
	returning id into new_order_id;

	foreach new_product_id in ARRAY inproducts loop
		insert into hw14.order_products (order_id, product_id, price) (select new_order_id, id, price from hw14.products where id = new_product_id);
	end loop;

	update hw14.orders set total_amount = sq.ta from (select sum(price) as ta from hw14.order_products where order_id = new_order_id) as sq where id = new_order_id;
	
	return new_order_id;
end;
$$ LANGUAGE plpgsql;

/* функция для удаления заказа */
CREATE OR REPLACE FUNCTION hw14.remove_order(remove_order_id int) returns bool as $$
begin
	delete from order_products where order_id = remove_order_id;
	delete from orders where id = remove_order_id;
	return true;
end;
$$ LANGUAGE plpgsql;

/* работа с пользователями */
insert into hw14.users values
(DEFAULT, 'иван', 'email1@mail.ru', 'o38947jy5djy43t', '3yt'),
(DEFAULT, 'маша', 'email2@mail.ru', 'mo8e4u5o9d8jyow', '3yt'),
(DEFAULT, 'даша', 'email3@mail.ru', 'mpoqu94850fq98y', '3yt'),
(DEFAULT, 'дима', 'email4@mail.ru', 'p8u5oihuxoqnydo', '3yt'),
(DEFAULT, 'настя', 'email5@mail.ru', 'p93845pmxoqp983', '3yt');
update hw14.users set name='кирилл' where email='email3@mail.ru';
delete from hw14.users where email='email3@mail.ru';

/* работа с родуктами */
insert into hw14.products values
(DEFAULT, 'спички', 143),
(DEFAULT, 'гречка', 54634),
(DEFAULT, 'соль', 56341),
(DEFAULT, 'туалетная бумага 3-х слойная', 6345123),
(DEFAULT, 'вода питьевая', 572343523);

update hw14.products set price = 344 where id = 1;
delete from hw14.products where id = 5;

/* выборка пользователей */

select * from hw14.users;

/* выборка токаров */

select * from hw14.products;

/* выборка заказов по пользователю */

select * from hw14.orders where user_id = <user_id>

/* статистика по пользователю */

select 
	u.id, u.name, u.email, sum(op.price), avg(op.price) 
from 
	hw14.orders as o 
	inner join hw14.order_products as op on (op.order_id = o.id)
	inner join hw14.users as u on (u.id = o.user_id)
where
	o.user_id = <user_id>
group by u.id;