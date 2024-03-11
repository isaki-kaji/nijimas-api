-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2024-03-11T14:43:11.954Z

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
  "sub_category1" varchar(255),
  "sub_category2" varchar(255),
  "room_id" varchar(255),
  "post_text" text,
  "photo_url" varchar(2000),
  "location" geometry,
  "meal_flag" boolean NOT NULL DEFAULT false,
  "public_type_no" char(1) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
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

CREATE TABLE "room" (
  "room_id" bigserial PRIMARY KEY,
  "room_name" varchar(255),
  "owner_id" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "follow_category" (
  "follow_category_id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "custom_category_id" bigint NOT NULL,
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

COMMENT ON COLUMN "post"."public_type_no" IS '1:公開、2:ルーム内で公開、3:非公開';

ALTER TABLE "user" ADD FOREIGN KEY ("currency") REFERENCES "currency" ("currency");

ALTER TABLE "post" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");

ALTER TABLE "post" ADD FOREIGN KEY ("main_category") REFERENCES "main_category" ("category_name");

ALTER TABLE "post" ADD FOREIGN KEY ("sub_category1") REFERENCES "sub_category" ("category_name");

ALTER TABLE "post" ADD FOREIGN KEY ("sub_category2") REFERENCES "sub_category" ("category_name");

ALTER TABLE "post" ADD FOREIGN KEY ("room_id") REFERENCES "room" ("room_id");

ALTER TABLE "comment" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("post_id");

ALTER TABLE "comment" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");

ALTER TABLE "room" ADD FOREIGN KEY ("owner_id") REFERENCES "user" ("user_id");

ALTER TABLE "follow_category" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");

ALTER TABLE "follow_category" ADD FOREIGN KEY ("custom_category_id") REFERENCES "room" ("room_id");

ALTER TABLE "meal" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("post_id");

ALTER TABLE "meal" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");
