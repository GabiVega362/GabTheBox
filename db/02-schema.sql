
CREATE TABLE IF NOT EXISTS "users" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "username" USERNAME NOT NULL,
  "email" EMAIL NOT NULL,
  "password" VARCHAR(32) NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT "pk_users" PRIMARY KEY ("id"),
  CONSTRAINT "uq_users_username" UNIQUE ("username"),
  CONSTRAINT "uq_users_email" UNIQUE ("email")
);


CREATE TABLE IF NOT EXISTS "labs" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "org" VARCHAR(255) NOT NULL,
  "environment" VARCHAR(255) NOT NULL,
  "release" VARCHAR(255) DEFAULT 'latest',
  "port" PORT NOT NULL DEFAULT 80,
  "description" VARCHAR(255) NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT "pk_labs" PRIMARY KEY ("id"),
  CONSTRAINT "uq_labs_image" UNIQUE ("org","environment","release")
);

CREATE TABLE IF NOT EXISTS "users_labs"(
  "user" uuid NOT NULL,
  "lab" uuid NOT NULL,
  "container" VARCHAR(255) NOT NULL,
  "port" PORT NOT NULL,
  "started_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT "pk_users_labs" PRIMARY KEY ("user"),

  CONSTRAINT "fk_users_labs_user" FOREIGN KEY ("user") REFERENCES "users" ("id") ON DELETE CASCADE,
  CONSTRAINT "fk_users_labs_lab" FOREIGN KEY ("lab") REFERENCES "labs" ("id") ON DELETE CASCADE
);