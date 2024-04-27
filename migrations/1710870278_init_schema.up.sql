CREATE TABLE "wallets" (
  "id" varchar PRIMARY KEY,
  "account_phone" varchar NOT NULL,
  "balance" bigint NOT NULL,
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
  "coa_type" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "transfers" (
  "id" varchar PRIMARY KEY,
  "src_wallet_id" varchar NOT NULL,
  "dst_wallet_id" varchar NOT NULL,
  "amount" integer NOT NULL,
  "type" varchar NOT NULL,
  "reference" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "entries" (
  "id" varchar PRIMARY KEY,
  "wallet_id" varchar NOT NULL,
  "credit_amount" integer NOT NULL,
  "debit_amount" integer NOT NULL,
  "balance_before" bigint NOT NULL,
  "balance_after" bigint NOT NULL,
  "correlation_id" varchar NOT NULL,
  "transfer_id" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "update_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE INDEX IF NOT EXISTS wallets_account_phone_idx ON wallets("account_phone");

CREATE INDEX IF NOT EXISTS wallets_account_phone_type_idx ON wallets("account_phone", "type");

CREATE INDEX IF NOT EXISTS transfers_src_wallet_id_idx ON transfers("src_wallet_id");

CREATE INDEX IF NOT EXISTS transfers_dst_wallet_id_idx ON transfers("dst_wallet_id");

CREATE INDEX IF NOT EXISTS transfers_src_dst_wallet_id_idx ON transfers("src_wallet_id", "dst_wallet_id");

CREATE INDEX IF NOT EXISTS entries_transfer_id_idx ON entries("transfer_id");

CREATE INDEX IF NOT EXISTS entries_wallet_id_idx ON entries("wallet_id");

CREATE INDEX IF NOT EXISTS entries_wallet_id_created_at_idx ON entries("wallet_id", "created_at");
