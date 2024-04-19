INSERT INTO "accounts" ("phone", "name", "email", "role", "status", "coa_type")
VALUES ('+62000000001', 'INTERNAL_COA_IN', 'coa-in@mail.com', 'INTERNAL_COA', 'ACTIVE', 'ASSETS');

INSERT INTO "accounts" ("phone", "name", "email", "role", "status", "coa_type")
VALUES ('+62000000002', 'INTERNAL_COA_OUT', 'coa-out@mail.com', 'INTERNAL_COA', 'ACTIVE', 'LIABILITIES');

INSERT INTO "wallets" ("id", "account_phone", "balance", "type")
VALUES('001', '+62000000001', 0.00, 'CASH');

INSERT INTO "wallets" ("id", "account_phone", "balance", "type")
VALUES('002', '+62000000002', 0.00, 'CASH');