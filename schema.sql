CREATE SCHEMA IF NOT EXISTS auth
    AUTHORIZATION postgres;

CREATE TABLE IF NOT EXISTS auth.users
(
    uuid uuid NOT NULL,
    username text COLLATE pg_catalog."default" NOT NULL,
    email_address text COLLATE pg_catalog."default" NOT NULL,
    metadata json,
    created_at timestamp without time zone NOT NULL,
    edited_at timestamp without time zone,
    deleted_at timestamp without time zone,
    photo_url text COLLATE pg_catalog."default",
    CONSTRAINT users_pkey PRIMARY KEY (uuid),
    CONSTRAINT users_email_address_key UNIQUE (email_address),
    CONSTRAINT users_username_key UNIQUE (username)
);