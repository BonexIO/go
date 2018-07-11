package resourceadapter

import (
	"context"

	"github.com/stivens13/go/xdr"
	. "github.com/stivens13/go/protocols/horizon"

)

func PopulateAsset(ctx context.Context, dest *Asset, asset xdr.Asset) error {
	return asset.Extract(&dest.Type, &dest.Code, &dest.Issuer)
}
