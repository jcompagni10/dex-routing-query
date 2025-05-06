package skip

// FungibleRouteRequest represents a request for a fungible token transfer route
type FungibleRouteRequest struct {
	AmountIn                  string     `json:"amount_in"`
	SourceAssetDenom          string     `json:"source_asset_denom"`
	SourceAssetChainID        string     `json:"source_asset_chain_id"`
	DestAssetDenom            string     `json:"dest_asset_denom"`
	DestAssetChainID          string     `json:"dest_asset_chain_id"`
	CumulativeAffiliateFeeBPS *string    `json:"cumulative_affiliate_fee_bps"`
	AllowMultiTx              *bool      `json:"allow_multi_tx"`
	SwapVenue                 *SwapVenue `json:"swap_venue"`
	AllowUnsafe               *bool      `json:"allow_unsafe"`
}

// FungibleRouteResponse represents the complete route response
type FungibleRouteResponse struct {
	AmountIn                      string        `json:"amount_in"`
	AmountOut                     string        `json:"amount_out"`
	SourceAssetDenom              string        `json:"source_asset_denom"`
	SourceAssetChainID            string        `json:"source_asset_chain_id"`
	DestAssetDenom                string        `json:"dest_asset_denom"`
	DestAssetChainID              string        `json:"dest_asset_chain_id"`
	Operations                    []Operation   `json:"operations"`
	ChainIDs                      []string      `json:"chain_ids"`
	RequiredChainAddresses        []string      `json:"required_chain_addresses"`
	DoesSwap                      bool          `json:"does_swap"`
	EstimatedAmountOut            string        `json:"estimated_amount_out"`
	SwapVenue                     SwapVenue     `json:"swap_venue"`
	TxsRequired                   int           `json:"txs_required"`
	USDAmountIn                   string        `json:"usd_amount_in"`
	USDAmountOut                  string        `json:"usd_amount_out"`
	SwapPriceImpactPercent        string        `json:"swap_price_impact_percent"`
	EstimatedFees                 []interface{} `json:"estimated_fees"`
	EstimatedRouteDurationSeconds int           `json:"estimated_route_duration_seconds"`
}

// Operation represents a single operation in the route
type Operation struct {
	AxelarTransfer *AxelarTransfer `json:"axelar_transfer,omitempty"`
	Swap           *Swap           `json:"swap,omitempty"`
	TxIndex        int             `json:"tx_index"`
}

// AxelarTransfer represents an Axelar bridge transfer operation
type AxelarTransfer struct {
	FromChain    string `json:"from_chain"`
	FromChainID  string `json:"from_chain_id"`
	ToChain      string `json:"to_chain"`
	ToChainID    string `json:"to_chain_id"`
	Asset        string `json:"asset"`
	ShouldUnwrap bool   `json:"should_unwrap"`
	FeeAmount    string `json:"fee_amount"`
	USDFeeAmount string `json:"usd_fee_amount"`
	FeeAsset     Asset  `json:"fee_asset"`
	IsTestnet    bool   `json:"is_testnet"`
	BridgeID     string `json:"bridge_id"`
}

// Asset represents a token asset
type Asset struct {
	Denom             string `json:"denom"`
	ChainID           string `json:"chain_id"`
	OriginDenom       string `json:"origin_denom"`
	OriginChainID     string `json:"origin_chain_id"`
	Trace             string `json:"trace"`
	IsCW20            bool   `json:"is_cw20"`
	IsEVM             bool   `json:"is_evm"`
	Symbol            string `json:"symbol"`
	Name              string `json:"name"`
	LogoURI           string `json:"logo_uri"`
	Decimals          int    `json:"decimals"`
	TokenContract     string `json:"token_contract"`
	RecommendedSymbol string `json:"recommended_symbol"`
}

// Swap represents a swap operation
type Swap struct {
	SwapIn                SwapIn `json:"swap_in"`
	EstimatedAffiliateFee string `json:"estimated_affiliate_fee"`
	FromChainID           string `json:"from_chain_id"`
	ChainID               string `json:"chain_id"`
	DenomIn               string `json:"denom_in"`
	DenomOut              string `json:"denom_out"`
}

// SwapIn represents the input side of a swap operation
type SwapIn struct {
	SwapVenue          SwapVenue       `json:"swap_venue"`
	SwapOperations     []SwapOperation `json:"swap_operations"`
	SwapAmountIn       string          `json:"swap_amount_in"`
	PriceImpactPercent string          `json:"price_impact_percent"`
}

// SwapVenue represents a swap venue
type SwapVenue struct {
	Name    string `json:"name"`
	ChainID string `json:"chain_id"`
}

// SwapOperation represents a single swap operation
type SwapOperation struct {
	Pool     string `json:"pool"`
	DenomIn  string `json:"denom_in"`
	DenomOut string `json:"denom_out"`
}

type ChainToAssetsMap struct {
	ChainToAssets map[string]ChainAssets `json:"chain_to_assets"`
}

// ChainAssets represents the assets for a specific chain
type ChainAssets struct {
	Assets []Asset `json:"assets"`
}

// ChainToAssetsResponse represents the complete response structure
type ChainToAssetsResponse struct {
	ChainToAssetsMap map[string]ChainAssets `json:"chain_to_assets_map"`
}
