CREATE TABLE "notifications" (
  "id" integer,
  "status" varchar NOT NULL DEFAULT 'Suspended',
  "create_user" integer NOT NULL DEFAULT 0,
  "modify_user" integer,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz
);

CREATE TABLE "services" (
  "id" integer PRIMARY KEY,
  "title" varchar(50),
  "short_name" varchar(30),
  "description" text,
  "send_notification" integer,
  "status" varchar NOT NULL DEFAULT 'Suspended',
  "create_user" integer NOT NULL DEFAULT 0,
  "modify_user" integer,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz
);
