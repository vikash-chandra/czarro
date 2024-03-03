CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "cz_vendors" (
  "id" bigserial PRIMARY KEY,
  "vendor_id" uuid NOT NULL DEFAULT uuid_generate_v4() UNIQUE,
  "vendor_name" varchar NOT NULL,
  "registration_number" varchar NOT NULL,
  "website_url" varchar NOT NULL,
  "contact_number" varchar NOT NULL UNIQUE,
  "contact_email" varchar NOT NULL,
  "country_code" integer NOT NULL,
  "salt" varchar NOT NULL,
  "password" varchar NOT NULL,
  "password_modifed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00+00',
  "status_id" integer NOT NULL,
  "create_user" bigint NOT NULL DEFAULT 0,
  "modify_user" bigint NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00+00'
);

CREATE TABLE "cz_vendors_shops" (
  "id" bigserial PRIMARY KEY,
  "vendor_id" bigserial NOT NULL,
  "shop_name" varchar(200) NOT NULL,
  "address" varchar(250) NOT NULL,
  "city" varchar(100) NOT NULL,
  "state" varchar(100) NOT NULL,
  "postal_code" varchar(20) NOT NULL,
  "country_code" integer NOT NULL,
  "location" varchar(100) NOT NULL,
  "create_user" bigint NOT NULL DEFAULT 0,
  "modify_user" bigint NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00+00'
);

CREATE TABLE "cz_vendors_workers" (
  "id" bigserial PRIMARY KEY,
  "worker_id" uuid NOT NULL DEFAULT uuid_generate_v4() UNIQUE,
  "shop_id" bigserial NOT NULL,
  "first_name" varchar NOT NULL,
  "middle_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "designation" varchar NOT NULL,
  "contact_number" varchar NOT NULL UNIQUE,
  "contact_email" varchar NOT NULL, 
  "dob" timestamptz NOT NULL DEFAULT '0001-01-01',
  "create_user" bigint NOT NULL DEFAULT 0,
  "modify_user" bigint NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00+00'
);

CREATE INDEX ON "cz_vendors" ("vendor_id");

CREATE INDEX ON "cz_vendors" ("vendor_name");

CREATE INDEX ON "cz_vendors_shops" ("shop_name");

CREATE INDEX ON "cz_vendors_shops" ("city");

CREATE INDEX ON "cz_vendors_workers" ("worker_id");

CREATE INDEX ON "cz_vendors_workers" ("contact_number");

ALTER TABLE "cz_vendors_shops" ADD FOREIGN KEY ("vendor_id") REFERENCES "cz_vendors" ("id");

ALTER TABLE "cz_vendors_workers" ADD FOREIGN KEY ("shop_id") REFERENCES "cz_vendors_shops" ("id");
