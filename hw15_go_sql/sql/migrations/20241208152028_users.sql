-- +goose Up
-- +goose StatementBegin
CREATE TABLE general.users (
	id uuid default general.new_uuid(),
	"name" varchar NULL,
	email varchar NOT NULL,
	password_hash varchar NOT NULL,
	password_salt varchar NOT NULL,
	CONSTRAINT users_pk PRIMARY KEY (id),
	CONSTRAINT users_unique_email UNIQUE (email)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table general.users;
-- +goose StatementEnd
