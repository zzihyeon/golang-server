CREATE TABLE "users" (
    "uid" bigserial PRIMARY KEY,
    "email" varchar NOT NULL,
    "name" varchar NOT NULL,
    "phone" varchar NOT NULL,
    "gender" varchar NOT NULL,
    "birth_date" timestamptz NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);