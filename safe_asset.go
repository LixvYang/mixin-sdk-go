package mixin

import (
	"context"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

type SafeAsset struct {
	AssetID        string          `json:"asset_id"`
	ChainID        string          `json:"chain_id"`
	FeeAssetID     string          `json:"fee_asset_id"`
	KernalAssetID  string          `json:"kernal_asset_id,omitempty"`
	Symbol         string          `json:"symbol,omitempty"`
	Name           string          `json:"name,omitempty"`
	IconURL        string          `json:"icon_url,omitempty"`
	PriceBTC       decimal.Decimal `json:"price_btc,omitempty"`
	PriceUSD       decimal.Decimal `json:"price_usd,omitempty"`
	ChangeBTC      decimal.Decimal `json:"change_btc,omitempty"`
	ChangeUsd      decimal.Decimal `json:"change_usd,omitempty"`
	PriceUpdatedAt time.Time       `json:"price_updated_at,omitempty"`
	AssetKey       string          `json:"asset_key,omitempty"`
	Precision      int32           `json:"precision,omitempty"`
	Dust           decimal.Decimal `json:"dust,omitempty"`
	Balance        decimal.Decimal `json:"balance,omitempty"`
	Confirmations  int             `json:"confirmations,omitempty"`
	KernelAssetID  string          `json:"kernel_asset_id,omitempty"`
}

func (c *Client) SafeReadAsset(ctx context.Context, assetID string) (*SafeAsset, error) {
	uri := fmt.Sprintf("/safe/assets/%s", assetID)

	var asset SafeAsset
	if err := c.Get(ctx, uri, nil, &asset); err != nil {
		return nil, err
	}

	return &asset, nil
}

func SafeReadAsset(ctx context.Context, accessToken, assetID string) (*SafeAsset, error) {
	return NewFromAccessToken(accessToken).SafeReadAsset(ctx, assetID)
}

func (c *Client) SafeReadAssets(ctx context.Context) ([]*SafeAsset, error) {
	var assets []*SafeAsset
	if err := c.Get(ctx, "/safe/assets", nil, &assets); err != nil {
		return nil, err
	}

	return assets, nil
}

func SafeReadAssets(ctx context.Context, accessToken string) ([]*SafeAsset, error) {
	return NewFromAccessToken(accessToken).SafeReadAssets(ctx)
}
