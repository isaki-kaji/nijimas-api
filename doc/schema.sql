-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2024-03-20T10:48:20.324Z

CREATE TABLE "user" (
  "user_id" bigserial PRIMARY KEY,
  "uid" varchar(255) NOT NULL,
  "username" varchar(255) NOT NULL,
  "currency" varchar(3) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "post" (
  "post_id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "main_category" varchar(255) NOT NULL,
  "post_text" text,
  "photo_url" varchar(2000),
  "location" geometry,
  "meal_flag" boolean NOT NULL DEFAULT false,
  "public_type_no" char(1) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "post_subcategory" (
  "post_subcategory_id" bigserial PRIMARY KEY,
  "post_id" bigint NOT NULL,
  "subcategory_no" char(1) NOT NULL,
  "sub_category" varchar(255) NOT NULL
);

CREATE TABLE "comment" (
  "comment_id" bigserial PRIMARY KEY,
  "post_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "comment_text" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "main_category" (
  "category_name" varchar(255) PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "sub_category" (
  "category_name" varchar(255) PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "follow_user" (
  "follow_id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "follow_user_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "currency" (
  "currency" varchar(3) PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "meal" (
  "meal_id" bigserial PRIMARY KEY,
  "post_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "calorie" bigint NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "user" ("uid");

CREATE INDEX ON "user" ("username");

CREATE INDEX ON "post" ("user_id");

CREATE INDEX ON "post" ("created_at");

CREATE INDEX ON "post_subcategory" ("post_id", "sub_category");

COMMENT ON COLUMN "post"."public_type_no" IS '1:公開、2:フォロワーにのみ公開、3:非公開';

ALTER TABLE "user" ADD FOREIGN KEY ("currency") REFERENCES "currency" ("currency");

ALTER TABLE "post" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");

ALTER TABLE "post" ADD FOREIGN KEY ("main_category") REFERENCES "main_category" ("category_name");

ALTER TABLE "post_subcategory" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("post_id");

ALTER TABLE "post_subcategory" ADD FOREIGN KEY ("sub_category") REFERENCES "sub_category" ("category_name");

ALTER TABLE "comment" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("post_id");

ALTER TABLE "comment" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");

ALTER TABLE "follow_user" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");

ALTER TABLE "follow_user" ADD FOREIGN KEY ("follow_user_id") REFERENCES "user" ("user_id");

ALTER TABLE "meal" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("post_id");

ALTER TABLE "meal" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");
