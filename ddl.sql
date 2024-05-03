CREATE TABLE IF NOT EXISTS senders (
    sender_id SERIAl PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    phone_number VARCHAR(15) NOT NULL UNIQUE,
    address VARCHAR(255) NOT NULL,
    address_note VARCHAR (100)
);

CREATE TABLE IF NOT EXISTS receivers (
    receiver_id SERIAl PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    phone_number VARCHAR(15) NOT NULL UNIQUE,
    address VARCHAR(255) NOT NULL,
    address_note VARCHAR (100)
);

CREATE TABLE IF NOT EXISTS item_type (
    item_type_id SERIAL PRIMARY KEY,
    item_name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS payment_method (
    payment_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS statuses (
    status_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS delivery_types (
    delivery_type_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description VARCHAR (255) 
);

CREATE TABLE IF NOT EXISTS transactions (
    transaction_id SERIAL PRIMARY KEY,
    order_number VARCHAR(50) NOT NULL UNIQUE,
    sender_id INT REFERENCES senders(sender_id),
    receiver_id INT REFERENCES receivers(receiver_id),
    item_type_id INT REFERENCES item_type(item_type_id),
    payment_id INT REFERENCES payment_method(payment_id),
    status_id INT REFERENCES statuses(status_id),
    delivery_type_id INT REFERENCES delivery_types(delivery_type_id)
);
