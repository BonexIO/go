package resourceadapter

import (
	"context"

	"github.com/BonexIO/go/amount"
	"github.com/BonexIO/go/services/horizon/internal/db2/assets"
	"github.com/BonexIO/go/xdr"
	. "github.com/BonexIO/go/protocols/horizon"
	"github.com/BonexIO/go/support/render/hal"
)

// PopulateAssetStat fills out the details
//func PopulateAssetStat(
func PopulateAssetStat(
	ctx context.Context,
	res *AssetStat,
	row assets.AssetStatsR,
) (err error) {

	res.Asset.Type = row.Type
	res.Asset.Code = row.Code
	res.Asset.Issuer = row.Issuer
	res.Amount, err = amount.IntStringToAmount(row.Amount)
	if err != nil {
		return err
	}
	res.NumAccounts = row.NumAccounts
	res.Flags = AccountFlags{
		(row.Flags & int8(xdr.AccountFlagsAuthRequiredFlag)) != 0,
		(row.Flags & int8(xdr.AccountFlagsAuthRevocableFlag)) != 0,
	}
	res.PT = row.SortKey

	res.Links.Toml = hal.NewLink(row.Toml)
	return
}
