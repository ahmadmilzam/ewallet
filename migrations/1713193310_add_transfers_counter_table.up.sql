CREATE TABLE "transfer_counters" (
  "wallet_id" varchar PRIMARY KEY,
  "credit_count_daily" int NOT NULL DEFAULT 0,
  "credit_count_monthly" int NOT NULL DEFAULT 0,
  "credit_amount_daily" bigint NOT NULL DEFAULT 0,
  "credit_amount_monthly" bigint NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);