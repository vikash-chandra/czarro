DROP TABLE IF EXISTS cz_users_address;
DROP TABLE IF EXISTS cz_users;
DROP TABLE IF EXISTS cz_roles;
DROP TABLE IF EXISTS cz_country;

ALTER TABLE "cz_users" ADD FOREIGN KEY ("role_id") REFERENCES "cz_roles" ("role_id");

ALTER TABLE "cz_users" ADD FOREIGN KEY ("country_code") REFERENCES "cz_country" ("id");

ALTER TABLE "cz_users_address" ADD FOREIGN KEY ("user_id") REFERENCES "cz_users" ("unique_id");

ALTER TABLE "cz_users_address" ADD FOREIGN KEY ("country_code") REFERENCES "cz_country" ("id");
