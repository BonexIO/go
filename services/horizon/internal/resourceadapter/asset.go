package resourceadapter

import (
	"context"

	"github.com/BonexIO/go/xdr"
	. "github.com/BonexIO/go/protocols/horizon"

)

func PopulateAsset(ctx context.Context, dest *Asset, asset xdr.Asset) error {
	return asset.Extract(&dest.Type, &dest.Code, &dest.Issuer)
}
