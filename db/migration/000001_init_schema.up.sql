CREATE SCHEMA "user";

CREATE TABLE "user"."core" (
  "id" serial PRIMARY KEY,
  "email" varchar(100) UNIQUE NOT NULL,
  "hash" text NOT NULL,
  "first_name" varchar(50) NOT NULL,
  "last_name" varchar(50) NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  "deleted_at" timestamptz DEFAULT null
);

CREATE TABLE "user"."core_role" (
  "core_id" int NOT NULL,
  "role_id" int NOT NULL
);

CREATE TABLE "user"."role" (
  "id" serial PRIMARY KEY,
  "role_name" varchar(50) UNIQUE NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  "deleted_at" timestamptz DEFAULT null
);

CREATE TABLE "user"."token" (
  "id" serial PRIMARY KEY,
  "user_id" int NOT NULL,
  "hash" text NOT NULL,
  "user_agent" text NOT NULL,
  "ip_address" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "expires_at" timestamptz DEFAULT (now() + interval '3 months'),
  "revoked_at" timestamptz DEFAULT null
);

CREATE INDEX ON "user"."core" ("email");

CREATE INDEX ON "user"."core" ("updated_at");

CREATE INDEX ON "user"."core" ("deleted_at");

CREATE UNIQUE INDEX ON "user"."core_role" ("core_id", "role_id");

CREATE INDEX ON "user"."role" ("role_name");

CREATE INDEX ON "user"."role" ("updated_at");

CREATE INDEX ON "user"."role" ("deleted_at");

CREATE INDEX ON "user"."token" ("user_id");

CREATE INDEX ON "user"."token" ("revoked_at");

COMMENT ON COLUMN "user"."core"."hash" IS 'bcrypt';

COMMENT ON COLUMN "user"."token"."hash" IS 'bcrypt';

ALTER TABLE "user"."core_role" ADD FOREIGN KEY ("core_id") REFERENCES "user"."core" ("id");

ALTER TABLE "user"."core_role" ADD FOREIGN KEY ("role_id") REFERENCES "user"."role" ("id");

ALTER TABLE "user"."token" ADD FOREIGN KEY ("user_id") REFERENCES "user"."core" ("id");
