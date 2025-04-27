package reporter

import "github.com/jcompagni10/skip-router-data/x/skip"

var (
	Amounts = []int{1, 50, 100, 1000, 5000}
	Pairs   = [][]string{
		// {"USDC", "NTRN"},
		// {"USDC", "TIA"},
		{"USDC", "DYDX"},
		// {"USDC", "OSMO"},
		// {"USDC", "ATOM"},
		// {"USDC", "WETH.axl"},
		// {"USDC", "WBTC.axl"},
	}
)

var ChainData = []skip.ChainData{

	{
		ChainID: "neutron-1",
		Denoms: []skip.ChainDenom{
			{
				IBCDenom: "ibc/B559A80D62249C8AA07A380E2A2BEA6E5CA9A6F079C912C3A9E9B494105E4F81",
				Symbol:   "USDC",
				Decimals: 6,
			},
			{
				IBCDenom: "ibc/C4CFF46FD6DE35CA4CF4CE031E643C8FDC9BA4B99AE598E9B0ED98FE3A2319F9",
				Symbol:   "ATOM",
				Decimals: 6,
			},
			{
				IBCDenom: "ibc/773B4D0A3CD667B2275D5A4A7A2F0909C0BA0F4059C0B9181E680DDF4965DCC7",
				Symbol:   "TIA",
				Decimals: 6,
			},
			{
				IBCDenom: "untrn",
				Symbol:   "NTRN",
				Decimals: 6,
			},
			{
				IBCDenom: "ibc/2CB87BCE0937B1D1DFCEE79BE4501AAF3C265E923509AEAC410AD85D27F35130",
				Symbol:   "DYDX",
				Decimals: 18,
			},
			{
				IBCDenom: "ibc/A585C2D15DCD3B010849B453A2CFCB5E213208A5AB665691792684C26274304D",
				Symbol:   "WETH.axl",
				Decimals: 18,
			},
			{
				IBCDenom: "ibc/376222D6D9DAE23092E29740E56B758580935A6D77C24C2ABD57A6A78A1F3955",
				Symbol:   "OSMO",
				Decimals: 6,
			},
			{
				IBCDenom: "ibc/DF8722298D192AAB85D86D0462E8166234A6A9A572DD4A2EA7996029DF4DB363",
				Symbol:   "WBTC.axl",
				Decimals: 8,
			},
		},
	},
	{
		ChainID: "cosmoshub-4",
		Denoms: []skip.ChainDenom{
			{
				IBCDenom: "ibc/F663521BF1836B00F5F177680F74BFB9A8B5654A694D0D2BC249E03CF2509013",
				Symbol:   "USDC",
				Decimals: 6,
			},
			{
				IBCDenom: "uatom",
				Symbol:   "ATOM",
				Decimals: 6,
			},
			{
				IBCDenom: "ibc/0025F8A87464A471E66B234C4F93AEC5B4DA3D42D7986451A059273426290DD5",
				Symbol:   "NTRN",
				Decimals: 6,
			},
			{
				IBCDenom: "ibc/14F9BC3E44B8A9C1BE1FB08980FAB87034C9905EF17CF2F5008FC085218811CC",
				Symbol:   "OSMO",
				Decimals: 6,
			},
		},
	},
	{
		ChainID: "osmosis-1",
		Denoms: []skip.ChainDenom{
			{
				IBCDenom: "ibc/498A0751C798A0D9A389AA3691123DADA57DAA4FE165D5C75894505B876BA6E4",
				Symbol:   "USDC",
				Decimals: 6,
			},
			{
				IBCDenom: "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2",
				Symbol:   "ATOM",
				Decimals: 6,
			},
			{
				IBCDenom: "ibc/D79E7D83AB399BFFF93433E54FAA480C191248FC556924A2A8351AE2638B3877",
				Symbol:   "TIA",
				Decimals: 6,
			},
			{
				IBCDenom: "ibc/126DA09104B71B164883842B769C0E9EC1486C0887D27A9999E395C2C8FB5682",
				Symbol:   "NTRN",
				Decimals: 6,
			},
			{
				IBCDenom: "uosmo",
				Symbol:   "OSMO",
				Decimals: 6,
			},
		},
	},
}
