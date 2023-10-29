CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "email" VARCHAR(255) NOT NULL UNIQUE,
  "hashed_password" VARCHAR(255) NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "strategies" (
  "id" SERIAL PRIMARY KEY,
  "user_id" INTEGER REFERENCES "users",
  "name" VARCHAR(255) NOT NULL,
  "indicators" JSON,
  "buyconditions" JSON,
  "sellconditions" JSON,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX ON "users" ("email");

