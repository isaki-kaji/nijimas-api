CREATE TABLE "users" (
  "uid" char(28) PRIMARY KEY,
  "username" varchar(255) NOT NULL,
  "self_intro" text,
  "profile_image_url" varchar(2000),
  "country_code" char(2),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "posts" (
  "post_id" uuid PRIMARY KEY,
  "uid" char(28) NOT NULL,
  "main_category" varchar(20) NOT NULL,
  "post_text" text,
  "photo_url" varchar(2000),
  "expense" numeric(15,2) NOT NULL DEFAULT 0,
  "location" varchar(2000),
  "public_type_no" char(1) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "post_subcategories" (
  "post_id" uuid,
  "category_no" char(1),
  "category_id" uuid NOT NULL,
  PRIMARY KEY ("post_id", "category_no")
);

CREATE TABLE "favorites" (
  "post_id" uuid,
  "uid" varchar(255),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("post_id", "uid")
);

CREATE TABLE "main_categories" (
  "category_name" varchar(20) PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "sub_categories" (
  "category_id" uuid PRIMARY KEY,
  "category_name" varchar(50) UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "follows" (
  "uid" char(28),
  "follow_uid" char(28),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("uid", "follow_uid")
);

CREATE TABLE "expense_summaries" (
  "uid" char(28),
  "year" integer NOT NULL,
  "month" integer NOT NULL,
  "main_category" varchar(20) NOT NULL,
  "amount" numeric(15,2) NOT NULL DEFAULT 0,
  PRIMARY KEY ("uid", "year", "month")
);

CREATE TABLE "subcategory_summaries" (
  "uid" char(28),
  "year" integer NOT NULL,
  "month" integer NOT NULL,
  "category_id" uuid NOT NULL,
  "amount" numeric(15,2) NOT NULL DEFAULT 0,
  PRIMARY KEY ("uid", "year", "month")
);

CREATE TABLE "daily_activity_summaries" (
  "uid" char(28),
  "year" integer NOT NULL,
  "month" integer NOT NULL,
  "day" integer NOT NULL,
  "number" integer NOT NULL DEFAULT 0,
  "amount" numeric(15,2) NOT NULL DEFAULT 0,
  PRIMARY KEY ("uid", "year", "month", "day")
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "posts" ("uid");

CREATE INDEX ON "favorites" ("post_id", "uid");

CREATE INDEX ON "follows" ("uid", "follow_uid");

COMMENT ON COLUMN "posts"."public_type_no" IS '0:公開、1:フォロワーにのみ公開、2:非公開';

ALTER TABLE "posts" ADD FOREIGN KEY ("uid") REFERENCES "users" ("uid");

ALTER TABLE "posts" ADD FOREIGN KEY ("main_category") REFERENCES "main_categories" ("category_name");

ALTER TABLE "post_subcategories" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("post_id");

ALTER TABLE "post_subcategories" ADD FOREIGN KEY ("category_id") REFERENCES "sub_categories" ("category_id");

ALTER TABLE "favorites" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("post_id");

ALTER TABLE "favorites" ADD FOREIGN KEY ("uid") REFERENCES "users" ("uid");

ALTER TABLE "follows" ADD FOREIGN KEY ("uid") REFERENCES "users" ("uid");

ALTER TABLE "follows" ADD FOREIGN KEY ("follow_uid") REFERENCES "users" ("uid");

ALTER TABLE "expense_summaries" ADD FOREIGN KEY ("uid") REFERENCES "users" ("uid");

ALTER TABLE "expense_summaries" ADD FOREIGN KEY ("main_category") REFERENCES "main_categories" ("category_name");

ALTER TABLE "subcategory_summaries" ADD FOREIGN KEY ("uid") REFERENCES "users" ("uid");

ALTER TABLE "subcategory_summaries" ADD FOREIGN KEY ("category_id") REFERENCES "sub_categories" ("category_id");

ALTER TABLE "daily_activity_summaries" ADD FOREIGN KEY ("uid") REFERENCES "users" ("uid");

INSERT INTO "main_categories" ("category_name") VALUES ('food');
INSERT INTO "main_categories" ("category_name") VALUES ('hobbies');
INSERT INTO "main_categories" ("category_name") VALUES ('fashion');
INSERT INTO "main_categories" ("category_name") VALUES ('goods');
INSERT INTO "main_categories" ("category_name") VALUES ('essentials');
INSERT INTO "main_categories" ("category_name") VALUES ('travel');
INSERT INTO "main_categories" ("category_name") VALUES ('entertainment');
INSERT INTO "main_categories" ("category_name") VALUES ('transport');
INSERT INTO "main_categories" ("category_name") VALUES ('other');
