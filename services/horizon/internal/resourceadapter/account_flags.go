package resourceadapter

import (
	"github.com/stivens13/go/services/horizon/internal/db2/core"
	. "github.com/stivens13/go/protocols/horizon"
)

func PopulateAccountFlags(dest *AccountFlags, row core.Account) {
	dest.AuthRequired = row.IsAuthRequired()
	dest.AuthRevocable = row.IsAuthRevocable()
}
