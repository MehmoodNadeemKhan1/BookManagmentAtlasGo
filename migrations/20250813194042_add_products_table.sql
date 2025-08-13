-- Create "customers" table
CREATE TABLE "public"."customers" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "customer_name" text NOT NULL,
  "customer_email" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_customers_customer_email" UNIQUE ("customer_email")
);
-- Create index "idx_customers_deleted_at" to table: "customers"
CREATE INDEX "idx_customers_deleted_at" ON "public"."customers" ("deleted_at");
-- Create "customer_addresses" table
CREATE TABLE "public"."customer_addresses" (
  "address_id" bigserial NOT NULL,
  "customer_id" bigint NOT NULL,
  "address" text NOT NULL,
  PRIMARY KEY ("address_id"),
  CONSTRAINT "fk_customers_address" FOREIGN KEY ("customer_id") REFERENCES "public"."customers" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "customer_phone_numbers" table
CREATE TABLE "public"."customer_phone_numbers" (
  "phone_number_id" bigserial NOT NULL,
  "customer_id" bigint NOT NULL,
  "phone_number" text NOT NULL,
  PRIMARY KEY ("phone_number_id"),
  CONSTRAINT "uni_customer_phone_numbers_phone_number" UNIQUE ("phone_number"),
  CONSTRAINT "fk_customers_phone_numbers" FOREIGN KEY ("customer_id") REFERENCES "public"."customers" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
