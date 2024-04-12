CREATE TABLE "wallets" (
  "id" varchar PRIMARY KEY,
  "account_phone" varchar NOT NULL,
  "balance" numeric(22, 2) NOT NULL,
  "type" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "accounts" (
  "phone" varchar PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "role" varchar NOT NULL,
  "status" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "journals" (
  "id" varchar PRIMARY KEY,
  "src_wallet_id" varchar NOT NULL,
  "dst_wallet_id" varchar NOT NULL,
  "amount" bigint NOT NULL,
  "reference" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "transfers" (
  "id" varchar PRIMARY KEY,
  "wallet_id" varchar NOT NULL,
  "credit_amount" bigint NOT NULL,
  "debit_amount" bigint NOT NULL,
  "correlation_id" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "journal_id" varchar NOT NULL
);

CREATE INDEX ON "wallets" ("account_phone");

CREATE INDEX ON "wallets" ("account_phone", "type");

CREATE INDEX ON "journals" ("src_wallet_id");

CREATE INDEX ON "journals" ("dst_wallet_id");

CREATE INDEX ON "journals" ("src_wallet_id", "dst_wallet_id");

CREATE INDEX ON "transfers" ("wallet_id");

CREATE INDEX ON "transfers" ("wallet_id", "created_at");