package resourceadapter

import (
	"context"

	"github.com/stivens13/go/amount"
	"github.com/stivens13/go/services/horizon/internal/assets"
	"github.com/stivens13/go/services/horizon/internal/db2/core"
	"github.com/stivens13/go/xdr"
	. "github.com/stivens13/go/protocols/horizon"
)

func PopulateBalance(ctx context.Context, dest *Balance, row core.Trustline) (err error) {
	dest.Type, err = assets.String(row.Assettype)
	if err != nil {
		return
	}

	dest.Balance = amount.String(row.Balance)
	dest.Limit = amount.String(row.Tlimit)
	dest.Issuer = row.Issuer
	dest.Code = row.Assetcode
	return
}

func PopulateNativeBalance(dest *Balance, stroops xdr.Int64) (err error) {
	dest.Type, err = assets.String(xdr.AssetTypeAssetTypeNative)
	if err != nil {
		return
	}

	dest.Balance = amount.String(stroops)
	dest.Limit = ""
	dest.Issuer = ""
	dest.Code = ""
	return
}
