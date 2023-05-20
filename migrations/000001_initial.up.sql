CREATE TABLE IF NOT EXISTS abay
  (
     created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     id           SERIAL PRIMARY KEY,
     full_name    VARCHAR NOT NULL,
     pseudonym    VARCHAR NOT NULL,
     specialty    VARCHAR NOT NULL
  );