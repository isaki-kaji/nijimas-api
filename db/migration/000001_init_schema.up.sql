CREATE TABLE "user" (
  "uid" varchar(255) PRIMARY KEY NOT NULL,
  "username" varchar(255) NOT NULL,
  "self_intro" text,
  "profile_image_url" varchar(2000),
  "banner_image_url" varchar(2000),
  "country_code" char(2),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "post" (
  "post_id" uuid PRIMARY KEY,
  "uid" varchar(255) NOT NULL,
  "main_category" varchar(255) NOT NULL,
  "post_text" text,
  "photo_url" varchar(2000),
  "expense" bigint,
  "location" varchar(2000),
  "public_type_no" char(1) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "post_subcategory" (
  "post_subcategory_id" bigserial PRIMARY KEY,
  "post_id" uuid NOT NULL,
  "subcategory_no" char(1) NOT NULL,
  "sub_category" varchar(255) NOT NULL
);

CREATE TABLE "favorite" (
  "favorite_id" bigserial PRIMARY KEY,
  "post_id" uuid NOT NULL,
  "uid" varchar(255) NOT NULL,
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
  "uid" varchar(255) NOT NULL,
  "follow_user_id" varchar(255) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "user" ("uid");

CREATE INDEX ON "user" ("username");

CREATE INDEX ON "post" ("uid");

CREATE INDEX ON "post" ("created_at");

CREATE INDEX ON "post_subcategory" ("post_id", "sub_category");

COMMENT ON COLUMN "post"."public_type_no" IS '1:公開、2:フォロワーにのみ公開、3:非公開';

ALTER TABLE "post" ADD FOREIGN KEY ("uid") REFERENCES "user" ("uid");

ALTER TABLE "post" ADD FOREIGN KEY ("main_category") REFERENCES "main_category" ("category_name");

ALTER TABLE "post_subcategory" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("post_id");

ALTER TABLE "post_subcategory" ADD FOREIGN KEY ("sub_category") REFERENCES "sub_category" ("category_name");

ALTER TABLE "favorite" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("post_id");

ALTER TABLE "favorite" ADD FOREIGN KEY ("uid") REFERENCES "user" ("uid");

ALTER TABLE "follow_user" ADD FOREIGN KEY ("uid") REFERENCES "user" ("uid");

ALTER TABLE "follow_user" ADD FOREIGN KEY ("follow_user_id") REFERENCES "user" ("uid");

INSERT INTO "main_category" ("category_name") VALUES ('food');
INSERT INTO "main_category" ("category_name") VALUES ('hobbies');
INSERT INTO "main_category" ("category_name") VALUES ('fashion');
INSERT INTO "main_category" ("category_name") VALUES ('goods');
INSERT INTO "main_category" ("category_name") VALUES ('essentials');
INSERT INTO "main_category" ("category_name") VALUES ('travel');
INSERT INTO "main_category" ("category_name") VALUES ('entertainment');
INSERT INTO "main_category" ("category_name") VALUES ('transport');
INSERT INTO "main_category" ("category_name") VALUES ('other');
