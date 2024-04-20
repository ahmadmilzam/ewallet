ALTER TABLE "transfers"
ADD COLUMN "updated_at" timestamptz NOT NULL DEFAULT 'now()';