package resourceadapter

import (
	"github.com/BonexIO/go/services/horizon/internal/db2/core"
	. "github.com/BonexIO/go/protocols/horizon"
)

func PopulateAccountFlags(dest *AccountFlags, row core.Account) {
	dest.AuthRequired = row.IsAuthRequired()
	dest.AuthRevocable = row.IsAuthRevocable()
}
