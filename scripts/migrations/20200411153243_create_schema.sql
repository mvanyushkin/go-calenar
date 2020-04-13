-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE SEQUENCE public.events_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE public.events_id_seq
    OWNER TO postgres;

CREATE TABLE public.events
(
    id bigint NOT NULL DEFAULT nextval('events_id_seq'::regclass),
    title text COLLATE pg_catalog."default",
    description text COLLATE pg_catalog."default",
    date timestamp without time zone,
    reminded boolean DEFAULT false,
    CONSTRAINT events_pkey PRIMARY KEY (id)
)

    TABLESPACE pg_default;

ALTER TABLE public.events
    OWNER to postgres;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
