CREATE TABLE "campaign" (
  "id" serial PRIMARY KEY,
  "campaign_name" varchar UNIQUE NOT NULL ,
  "start" timestamptz NOT NULL,
  "end" timestamptz NOT NULL,
  "active" boolean NOT NULL DEFAULT (false)
);

CREATE INDEX ON "campaign" ("campaign_name");