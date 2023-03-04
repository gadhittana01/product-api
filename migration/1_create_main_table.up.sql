CREATE TABLE IF NOT EXISTS users(
   id SERIAL PRIMARY KEY,
   full_name VARCHAR NOT NULL,
   first_order TIMESTAMP,
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS order_items(
   id SERIAL PRIMARY KEY,
   name VARCHAR NOT NULL,
   price NUMERIC NOT NULL,
   expired_at TIMESTAMP NOT NULL,
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS order_histories(
   id SERIAL PRIMARY KEY,
   user_id INT NOT NULL REFERENCES users(id),
   order_item_id INT NOT NULL REFERENCES order_items(id),
   descriptions VARCHAR NOT NULL,
   created_at TIMESTAMP,
   updated_at TIMESTAMP
);