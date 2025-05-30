package reporter

import (
	"database/sql"
	"fmt"
	"math"
	"math/big"
	"slices"
	"strconv"
	"time"

	"github.com/jcompagni10/skip-router-data/x/skip"
	log "github.com/sirupsen/logrus"
)

type SwapResult struct {
	Venue string
	Price float64
}

type RouteResult struct {
	Winner       string
	WinningPrice float64
	NeutronPrice float64
	TokenIn      string
	TokenOut     string
	AmountIn     int
	Time         time.Time
	SourceChain  string
}

func GetSwapRoute(amount *big.Int, chain string, denomIn skip.Asset, denomOut skip.Asset, venue ...string) (SwapResult, error) {

	allowUnsafe := true
	req := &skip.FungibleRouteRequest{
		AmountIn:           amount.String(),
		SourceAssetChainID: chain,
		SourceAssetDenom:   denomIn.Denom,
		DestAssetChainID:   chain,
		DestAssetDenom:     denomOut.Denom,
		AllowUnsafe:        &allowUnsafe,
	}
	if len(venue) > 0 {
		req.SwapVenue = &skip.SwapVenue{
			Name: venue[0],
		}
	}

	// Create a channel to handle the timeout
	resultChan := make(chan struct {
		resp *skip.FungibleRouteResponse
		err  error
	}, 1)

	// Start the request in a goroutine
	go func() {
		resp, err := skip.GetFungibleRoutes(req)
		resultChan <- struct {
			resp *skip.FungibleRouteResponse
			err  error
		}{resp, err}
	}()

	// Wait for either the result or timeout
	select {
	case result := <-resultChan:
		if result.err != nil {
			return SwapResult{}, fmt.Errorf("error getting route: %v to %v (%v) venue=%v %#v: %v", denomIn.Denom, denomOut.Denom, chain, venue, req, result.err)
		}
		resp := result.resp

		amountIn, err := strconv.ParseFloat(resp.AmountIn, 64)
		if err != nil {
			return SwapResult{}, fmt.Errorf("error parsing amount in (%v): %v", resp.AmountIn, err)
		}

		amountOut, err := strconv.ParseFloat(resp.AmountOut, 64)
		if err != nil {
			return SwapResult{}, fmt.Errorf("error parsing amount out (%v): %v", resp.AmountOut, err)
		}

		decimalsDifference := float64(denomOut.Decimals) - float64(denomIn.Decimals)
		price := (amountOut / amountIn) / math.Pow(10, decimalsDifference)

		return SwapResult{
			Venue: resp.SwapVenue.Name,
			Price: price,
		}, nil
	case <-time.After(10 * time.Second):
		return SwapResult{}, fmt.Errorf("timeout getting routes")
	}
}

func ReportSwapRoutes(db *sql.DB) {
	for _, pair := range Pairs {
		price0, err := GetTokenPriceCached(db, pair[0])
		if err != nil {
			log.Error(err)
			continue
		}
		price1, err := GetTokenPriceCached(db, pair[1])
		if err != nil {
			log.Error(err)
			continue
		}

		for _, amount := range Amounts {
			for _, chainID := range ChainIds {
				// check both directions
				for _, reverseDirection := range []bool{false, true} {

					denoms, err := GetDenomsForChain(chainID, pair)
					if err != nil {
						log.Error(err)
						continue
					}

					var denomIn, denomOut skip.Asset
					var amountIn *big.Int
					if reverseDirection {
						denomIn, denomOut = denoms[1], denoms[0]
						amountIn = CalcAmountIn(amount, price1, denomIn)

					} else {
						denomIn, denomOut = denoms[0], denoms[1]
						amountIn = CalcAmountIn(amount, price0, denomIn)

					}
					log.Infof("%v Processing pair: %v to %v, $amount: %v; amountIn: %v price0: %v price1: %v", chainID, denomIn.RecommendedSymbol, denomOut.RecommendedSymbol, amount, amountIn, price0, price1)
					swapResult, err := GetSwapRoute(amountIn, chainID, denomIn, denomOut)
					if err != nil {
						log.Error(err)
						continue
					}

					routeResult := RouteResult{
						Winner:       swapResult.Venue,
						WinningPrice: swapResult.Price,
						TokenIn:      denomIn.Symbol,
						TokenOut:     denomOut.Symbol,
						AmountIn:     amount,
						Time:         time.Now(),
						SourceChain:  chainID,
					}

					if routeResult.Winner != "neutron-duality" {
						neutronSwapResult, err := GetSwapRoute(amountIn, chainID, denomIn, denomOut, "neutron-duality")
						if err != nil {
							log.Error(err)
						} else {
							routeResult.NeutronPrice = neutronSwapResult.Price
						}
					} else {
						routeResult.NeutronPrice = swapResult.Price
					}

					err = insertRouteResult(db, routeResult)
					if err != nil {
						log.Errorf("Error inserting route result: %v", err)
					}
				}
			}
		}
	}
}

func insertRouteResult(db *sql.DB, routeResult RouteResult) error {
	_, err := db.Exec("INSERT INTO swap_routes (winner, winning_price, neutron_price, token_in, token_out, amount_in, time, source_chain) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		routeResult.Winner, routeResult.WinningPrice, routeResult.NeutronPrice, routeResult.TokenIn, routeResult.TokenOut, routeResult.AmountIn, routeResult.Time, routeResult.SourceChain)
	return err
}

func GetDenomForChain(chain string, symbol string) (skip.Asset, error) {

	if _, ok := Exclusions[chain]; ok {
		if slices.Contains(Exclusions[chain], symbol) {
			return skip.Asset{}, fmt.Errorf("skipping denom %v for chain: %v", symbol, chain)
		}
	}
	for chainID, chainData := range ChainData {
		if chainID != chain {
			continue
		}
		for _, denomData := range chainData.Assets {
			if denomData.RecommendedSymbol == symbol {
				return denomData, nil
			}
		}
	}
	return skip.Asset{}, fmt.Errorf("denom: %v not found for chain: %v", symbol, chain)
}

func GetDenomsForChain(chain string, symbols []string) ([]skip.Asset, error) {
	denoms := make([]skip.Asset, len(ChainData))
	for i, symbol := range symbols {
		denom, err := GetDenomForChain(chain, symbol)
		if err != nil {
			return nil, err
		}
		denoms[i] = denom
	}
	return denoms, nil
}
