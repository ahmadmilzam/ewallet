CREATE TABLE "wallets" (
  "id" varchar PRIMARY KEY,
  "account_id" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "type" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "accounts" (
  "id" varchar PRIMARY KEY,
  "phone" int UNIQUE NOT NULL,
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
  "partner_reference" varchar,
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

CREATE INDEX ON "wallets" ("account_id");

CREATE INDEX ON "wallets" ("account_id", "type");

CREATE INDEX ON "journals" ("src_wallet_id");

CREATE INDEX ON "journals" ("dst_wallet_id");

CREATE INDEX ON "journals" ("src_wallet_id", "dst_wallet_id");

CREATE INDEX ON "transfers" ("wallet_id");

CREATE INDEX ON "transfers" ("wallet_id", "created_at");

COMMENT ON COLUMN "transfers"."credit_amount" IS 'must be positive';

COMMENT ON COLUMN "transfers"."debit_amount" IS 'must be positive';

ALTER TABLE "wallets" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("wallet_id") REFERENCES "wallets" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("journal_id") REFERENCES "journals" ("id");

ALTER TABLE "journals" ADD FOREIGN KEY ("src_wallet_id") REFERENCES "wallets" ("id");

ALTER TABLE "journals" ADD FOREIGN KEY ("dst_wallet_id") REFERENCES "wallets" ("id");
