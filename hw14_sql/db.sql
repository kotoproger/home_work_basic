/* структура бд */
create database hw14;

CREATE TABLE public.users (
	id bigserial NOT NULL,
	"name" varchar NULL,
	email varchar NOT NULL,
	password_hash varchar NOT NULL,
	password_salt varchar NOT NULL,
	CONSTRAINT users_pk PRIMARY KEY (id),
	CONSTRAINT users_unique_email UNIQUE (email)
);

CREATE TABLE public.orders (
	id bigserial NOT NULL,
	user_id bigint NOT NULL,
	order_date timestamp NOT NULL,
	total_amount integer NOT null,
	CONSTRAINT orders_pk PRIMARY KEY (id)
);
comment on column orders.total_amount is 'в копейках';
ALTER TABLE public.orders ADD CONSTRAINT orders_users_fk FOREIGN KEY (user_id) REFERENCES public.users(id);

create table public.products (
	id bigserial not null,
	"name" varchar not null,
	price integer not null,
	CONSTRAINT products_pk PRIMARY KEY (id)
);
comment on column products.price is 'в копейках';

create table order_products(
	id bigserial not null,
	order_id bigint NOT NULL,
	product_id bigint NOT NULL,
	price integer not null,
	CONSTRAINT order_products_pk PRIMARY KEY (id)
);
comment on column order_products.price is 'в копейках';
ALTER TABLE public.order_products ADD CONSTRAINT order_products_order_fk FOREIGN KEY (order_id) REFERENCES public.orders(id);
ALTER TABLE public.order_products ADD CONSTRAINT order_products_product_fk FOREIGN KEY (product_id) REFERENCES public.products(id);

/* функция для создания заказа */
CREATE OR REPLACE FUNCTION create_order(inemail varchar, inproducts bigint[]) returns bigint as $$
DECLARE 
 new_order_id bigint;
 new_product_id bigint;
 new_total_amount int;
begin
	insert into orders (user_id, order_date, total_amount) (select id, now(), 0 from users where email = inemail)
	returning id into new_order_id;

	foreach new_product_id in ARRAY inproducts loop
		insert into order_products (order_id, product_id, price) (select new_order_id, id, price from products where id = new_product_id);
	end loop;

	update orders set total_amount = sq.ta from (select sum(price) as ta from order_products where order_id = new_order_id) as sq where id = new_order_id;
	
	return new_order_id;
end;
$$ LANGUAGE plpgsql;

/* функция для удаления заказа */
CREATE OR REPLACE FUNCTION remove_order(remove_order_id int) returns bool as $$
begin
	delete from order_products where order_id = remove_order_id;
	delete from orders where id = remove_order_id;
	return true;
end;
$$ LANGUAGE plpgsql;

/* работа с пользователями */
insert into users values
(DEFAULT, 'иван', 'email1@mail.ru', 'o38947jy5djy43t', '3yt'),
(DEFAULT, 'маша', 'email2@mail.ru', 'mo8e4u5o9d8jyow', '3yt'),
(DEFAULT, 'даша', 'email3@mail.ru', 'mpoqu94850fq98y', '3yt'),
(DEFAULT, 'дима', 'email4@mail.ru', 'p8u5oihuxoqnydo', '3yt'),
(DEFAULT, 'настя', 'email5@mail.ru', 'p93845pmxoqp983', '3yt');
update users set name='кирилл' where email='email3@mail.ru';
delete from users where email='email3@mail.ru';

/* работа с родуктами */
insert into products values
(DEFAULT, 'спички', 143),
(DEFAULT, 'гречка', 54634),
(DEFAULT, 'соль', 56341),
(DEFAULT, 'туалетная бумага 3-х слойная', 6345123),
(DEFAULT, 'вода питьевая', 572343523);

update products set price = 344 where id = 1;
delete from products where id = 5;

/* выборка пользователей */

select * from users;

/* выборка токаров */

select * from products;

/* выборка заказов по пользователю */

select * from orders where user_id = <user_id>

/* статистика по пользователю */

select 
	o.user_id, sum(op.price), avg(op.price) 
from 
	orders as o 
	inner join order_products as op on (op.order_id = o.id)
where
	o.user_id = <user_id>
group by o.user_id;