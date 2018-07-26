#! /usr/bin/env bash
set -e

# This scripts rebuilds the latest.sql file included in the schema package.
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
GOTOP="$( cd "$DIR/../../../../../../../.." && pwd )"
go generate github.com/stivens13/go/services/horizon/internal/db2/schema
go install github.com/stivens13/go/services/horizon
dropdb horizon --if-exists
createdb horizon
DATABASE_URL=postgres://localhost/horizon?sslmode=disable $GOTOP/bin/horizon db migrate up

DUMP_OPTS="--schema=public --no-owner --no-acl --inserts"
LATEST_PATH="$DIR/../db2/schema/latest.sql"
BLANK_PATH="$DIR/../test/scenarios/blank-horizon.sql"

pg_dump postgres://localhost/horizon?sslmode=disable $DUMP_OPTS \
  | sed '/SET idle_in_transaction_session_timeout/d'  \
  | sed '/SET row_security/d' \
  > $LATEST_PATH
pg_dump postgres://localhost/horizon?sslmode=disable \
  --clean --if-exists $DUMP_OPTS \
  | sed '/SET idle_in_transaction_session_timeout/d'  \
  | sed '/SET row_security/d' \
  > $BLANK_PATH

go generate github.com/stivens13/go/services/horizon/internal/db2/schema
go generate github.com/stivens13/go/services/horizon/internal/test
go install github.com/stivens13/go/services/horizon
