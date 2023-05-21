CREATE TABLE IF NOT EXISTS users
  (
     created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     id           SERIAL PRIMARY KEY,
     name varchar,
     password varchar,
     email varchar
  );