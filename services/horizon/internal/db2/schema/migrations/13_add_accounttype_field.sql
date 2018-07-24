-- +migrate Up

ALTER TABLE ONLY history_accounts ADD COLUMN accounttype bigint DEFAULT NULL;

-- +migrate Down

ALTER TABLE ONLY history_accounts DROP COLUMN accounttype;
