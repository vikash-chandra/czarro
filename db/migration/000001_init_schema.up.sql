CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "cz_country" (
  "id" serial PRIMARY KEY,
  "iso" varchar(2) NOT NULL,
  "name" varchar(80) NOT NULL,
  "nicename" varchar(80) NOT NULL,
  "iso3" varchar(3) NOT NULL,
  "numcode" integer NOT NULL DEFAULT 0,
  "phone_code" integer NOT NULL DEFAULT 0
);

CREATE TABLE "cz_roles" (
  "id" serial PRIMARY KEY,
  "role_name" varchar(20) NOT NULL,
  "status_id" integer NOT NULL,
  "create_user" bigserial NOT NULL,
  "modify_user" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00+00',
  "visible" boolean NOT NULL DEFAULT true
);

CREATE TABLE "cz_users" (
  "id" bigserial NOT NULL,
  "unique_id" uuid PRIMARY KEY DEFAULT uuid_generate_v4() UNIQUE,
  "role_id" integer NOT NULL,
  "first_name" varchar NOT NULL,
  "middle_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "dob" timestamptz NOT NULL DEFAULT '0001-01-01',
  "country_code" integer NOT NULL,
  "phone" varchar NOT NULL,
  "email" varchar NOT NULL,
  "salt" varchar NOT NULL,
  "password" varchar NOT NULL,
  "status_id" integer NOT NULL,
  "create_user" bigserial NOT NULL,
  "modify_user" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00+00'
);

CREATE TABLE "cz_users_address" (
  "id" bigserial PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "country_code" integer NOT NULL,
  "address1" varchar(100) NOT NULL,
  "address2" varchar(100) NOT NULL,
  "address3" varchar(100) NOT NULL,
  "address4" varchar(100) NOT NULL,
  "location" varchar(100) NOT NULL,
  "status_id" integer NOT NULL,
  "create_user" bigserial NOT NULL,
  "modify_user" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00+00'
);

CREATE TABLE "cz_notification" (
  "id" bigserial PRIMARY KEY,
  "status_id" integer NOT NULL DEFAULT 0,
  "create_user" bigserial NOT NULL,
  "modify_user" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00+00'
);

CREATE TABLE "cz_services" (
  "id" serial PRIMARY KEY,
  "title" varchar(50) NOT NULL,
  "short_name" varchar(30) NOT NULL,
  "description" text NOT NULL,
  "send_notification" integer NOT NULL,
  "status_id" integer NOT NULL,
  "create_user" bigserial NOT NULL,
  "modify_user" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00+00'
);

CREATE INDEX ON "cz_users" ("unique_id");

CREATE INDEX ON "cz_users" ("phone");

CREATE INDEX ON "cz_users_address" ("user_id");

ALTER TABLE "cz_users" ADD FOREIGN KEY ("role_id") REFERENCES "cz_roles" ("id");

ALTER TABLE "cz_users" ADD FOREIGN KEY ("country_code") REFERENCES "cz_country" ("id");

ALTER TABLE "cz_users_address" ADD FOREIGN KEY ("user_id") REFERENCES "cz_users" ("unique_id");

ALTER TABLE "cz_users_address" ADD FOREIGN KEY ("country_code") REFERENCES "cz_country" ("id");

INSERT INTO cz_roles (id, role_name, status_id, visible) 
VALUES (100, 'user', 1, true);

INSERT INTO cz_country (id, iso, name, nicename, iso3, numcode, phone_code) 
VALUES (1, 'in', 'India', 'IND', 'iso', 91, 91);
