-- +migrate Up

ALTER TABLE ONLY history_accounts ADD COLUMN accounttype bigint NOT NULL;

-- +migrate Down

ALTER TABLE ONLY history_accounts DROP COLUMN accounttype;
