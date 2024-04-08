INSERT INTO "accounts" ("id", "phone", "name", "email", "role", "status")
VALUES ('001', '001', 'INTERNAL_COA_IN', 'coa-in@mail.com', 'INTERNAL_COA', 'ACTIVE');

INSERT INTO "accounts" ("id", "phone", "name", "email", "role", "status")
VALUES ('002', '002', 'INTERNAL_COA_OUT', 'coa-out@mail.com', 'INTERNAL_COA', 'ACTIVE');

INSERT INTO "wallets" ("id", "account_id", "balance", "type")
VALUES('001', '001', 0.00, 'CASH');

INSERT INTO "wallets" ("id", "account_id", "balance", "type")
VALUES('002', '002', 0.00, 'CASH');