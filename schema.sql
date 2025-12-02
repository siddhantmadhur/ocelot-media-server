CREATE SCHEMA IF NOT EXISTS auth
    AUTHORIZATION postgres;

CREATE TABLE IF NOT EXISTS auth.users
(
    uid uuid NOT NULL,
    username text COLLATE pg_catalog."default" NOT NULL,
    encrypted_passworrd text NOT NULL,
    photo_url text NOT NULL,
    created_at timestamp without time zone NOT NULL,
    deleted_at timestamp without time zone,
    CONSTRAINT users_pkey PRIMARY KEY (uid),
    CONSTRAINT users_username_key UNIQUE (username)
);

CREATE TABLE IF NOT EXISTS auth.permissions
(
    id BIGINT PRIMARY KEY,
    created_at timestamp without time zone NOT NULL,
);
