CREATE TABLE "customers" (
  "id" bigserial PRIMARY KEY,
  "role_id" int NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "dob" date DEFAULT NULL,
  "mobile" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "status" varchar(20) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
