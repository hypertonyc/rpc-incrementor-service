CREATE TABLE "numbers" (
  "id" bigserial PRIMARY KEY,
  "value" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "settings" (
  "id" bigserial PRIMARY KEY,
  "number_id" bigint NOT NULL,
  "increment_step" int NOT NULL,
  "upper_limit" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "settings" ADD FOREIGN KEY ("number_id") REFERENCES "numbers" ("id");