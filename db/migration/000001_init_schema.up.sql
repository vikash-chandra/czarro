CREATE TABLE "status" (
  "id" integer PRIMARY KEY,
  "status" varchar(20) NOT NULL,
  "create_user" integer NOT NULL DEFAULT 0,
  "modify_user" integer,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz
);

CREATE TABLE "roles" (
  "role_id" integer PRIMARY KEY,
  "role_name" varchar(20),
  "status_id" integer,
  "create_user" integer NOT NULL DEFAULT 0,
  "modify_user" integer,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz,
  "visible" boolean NOT NULL DEFAULT true
);

CREATE TABLE "customers" (
  "id" bigserial PRIMARY KEY,
  "unique_id" uuid NOT NULL,
  "role_id" integer,
  "first_name" varchar NOT NULL,
  "middle_name" varchar,
  "last_name" varchar NOT NULL,
  "dob" date,
  "country_code" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "email" varchar,
  "salt" varchar,
  "password" varchar,
  "status_id" integer,
  "create_user" integer NOT NULL DEFAULT 0,
  "modify_user" integer,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz
);

CREATE TABLE "customers_address" (
  "id" bigserial PRIMARY KEY,
  "customer_id" integer,
  "address1" varchar NOT NULL,
  "address2" varchar,
  "address3" varchar,
  "address4" varchar,
  "location" varchar NOT NULL,
  "status_id" integer,
  "create_user" integer NOT NULL DEFAULT 0,
  "modify_user" integer,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz
);

CREATE INDEX ON "customers" ("unique_id");

CREATE INDEX ON "customers" ("phone");

CREATE INDEX ON "customers_address" ("customer_id");

ALTER TABLE "roles" ADD FOREIGN KEY ("status_id") REFERENCES "status" ("id");

ALTER TABLE "customers" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("role_id");

ALTER TABLE "customers" ADD FOREIGN KEY ("status_id") REFERENCES "status" ("id");

ALTER TABLE "customers_address" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("unique_id");

ALTER TABLE "customers_address" ADD FOREIGN KEY ("status_id") REFERENCES "status" ("id");
