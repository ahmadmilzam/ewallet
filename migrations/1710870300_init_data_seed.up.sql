INSERT INTO "accounts" ("phone", "name", "email", "role", "status")
VALUES ('62001', 'INTERNAL_COA_IN', 'coa-in@mail.com', 'INTERNAL_COA', 'ACTIVE');

INSERT INTO "accounts" ("phone", "name", "email", "role", "status")
VALUES ('62002', 'INTERNAL_COA_OUT', 'coa-out@mail.com', 'INTERNAL_COA', 'ACTIVE');

INSERT INTO "wallets" ("id", "account_phone", "balance", "type")
VALUES('001', '62001', 0.00, 'CASH');

INSERT INTO "wallets" ("id", "account_phone", "balance", "type")
VALUES('002', '62002', 0.00, 'CASH');