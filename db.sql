DROP TABLE IF EXISTS orders;

CREATE TABLE public.orders (
    order_uid VARCHAR(100) NOT NULL,
    order_data json NOT NULL
);