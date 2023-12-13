CREATE TABLE "roles" (
  "role_id" int NOT NULL,
  "role_name" varchar(20) NOT NULL,
  "status_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz DEFAULT NULL,
  "visible" boolean DEFAULT true
);

CREATE TABLE "status" (
  "id" int NOT NULL,
  "status" varchar(20) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz DEFAULT NULL,
  "visible" boolean DEFAULT true
);

CREATE TABLE "customers" (
  "id" bigserial PRIMARY KEY,
  "role_id" int NOT NULL,
  "first_name" varchar NOT NULL,
  "middle_name" varchar DEFAULT NULL,
  "last_name" varchar NOT NULL,
  "dob" date DEFAULT NULL,
  "mobile" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "status" varchar(20) NOT NULL,
  "create_user" int DEFAULT NULL,
  "modify_user" int DEFAULT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz DEFAULT NULL
);

CREATE TABLE "customers_address" (
  "id" bigserial PRIMARY KEY,
  "customer_id" varchar(20) NOT NULL,
  "address1" varchar NOT NULL,
  "address2" varchar NOT NULL,
  "address3" varchar NOT NULL,
  "address4" varchar NOT NULL,
  "location" varchar NOT NULL,
  "status_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz DEFAULT NULL
);