CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS public.baskets
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    CONSTRAINT baskets_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.basket_products
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    basketid uuid,
    productid uuid,
    quantity integer,
    CONSTRAINT basket_products_pkey PRIMARY KEY (id),

    CONSTRAINT fk_basket_products_basketid
    FOREIGN KEY(basketid)
    REFERENCES baskets(id),

    CONSTRAINT fk_basket_products_productid
    FOREIGN KEY(productid)
    REFERENCES products(id)

);

CREATE TABLE IF NOT EXISTS public.products
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    name text COLLATE pg_catalog."default",
    description text COLLATE pg_catalog."default",
    price numeric,
    quantity integer,
    status text COLLATE pg_catalog."default",
    CONSTRAINT products_pkey PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public.baskets OWNER to postgres;
ALTER TABLE IF EXISTS public.basket_products OWNER to postgres;
ALTER TABLE IF EXISTS public.products OWNER to postgres;