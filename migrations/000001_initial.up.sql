CREATE TABLE IF NOT EXISTS users(
     id           varchar PRIMARY KEY,
     name         varchar not null ,
     password     varchar not null ,
     email        varchar not null unique ,
     created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE table IF NOT EXISTS courses(
    id            varchar primary key,
    name          varchar not null,
    description   varchar not null,
    imgURL   varchar not null
);

CREATE table if not exists orders(
    order_id      varchar PRIMARY KEY ,
    user_id       varchar NOT NULL ,
    course_id     varchar not null,
    created_at    timestamp DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "orders" add foreign key ("course_id") references  "courses" ("id");

CREATE INDEX ON "users" ("id");
CREATE INDEX ON "courses" ("id");
CREATE INDEX ON "orders" ("user_id");
