package reporter

import (
	"database/sql"
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/jcompagni10/skip-router-data/x/skip"
)

func GetTokenPrice(db *sql.DB, denom string) (float64, error) {
	if denom == "USDC" {
		return 1, nil
	}
	query := `
		SELECT winning_price
		FROM swap_routes
		WHERE token_in = "USDC" AND token_out = $1
		ORDER BY time DESC
		LIMIT 1
	`
	var price float64
	err := db.QueryRow(query, denom).Scan(&price)
	if err != nil {
		return 0, err
	}
	return price, nil
}

type cachedPrice struct {
	price     float64
	timestamp time.Time
}

var priceCache = make(map[string]cachedPrice)
var priceCacheMutex sync.RWMutex

func GetTokenPriceCached(db *sql.DB, denom string) (float64, error) {
	if denom == "USDC" {
		return 1, nil
	}

	now := time.Now()
	cacheExpiry := 5 * time.Minute

	// Check cache first
	priceCacheMutex.RLock()
	if cached, exists := priceCache[denom]; exists {
		if now.Sub(cached.timestamp) < cacheExpiry {
			priceCacheMutex.RUnlock()
			return cached.price, nil
		}
	}
	priceCacheMutex.RUnlock()

	// Not in cache or expired, look up in DB
	price, err := GetTokenPrice(db, denom)
	if err != nil {
		return 0, fmt.Errorf("error getting token price for %v: %v", denom, err)
	}

	// Store in cache
	priceCacheMutex.Lock()
	priceCache[denom] = cachedPrice{
		price:     price,
		timestamp: now,
	}
	priceCacheMutex.Unlock()

	return price, nil
}

func SeedPriceCache(symbols []string) {

	for _, symbol := range symbols {
		denoms, err := GetDenomsForChain("neutron-1", []string{symbol, "USDC"})
		if err != nil {
			panic(fmt.Sprintf("error seeding price cache for token %v: %v", symbol, err))

		}
		amountIn := math.Pow(10, float64(denoms[0].Decimals))
		result, err := GetSwapRoute(big.NewInt(int64(amountIn)), "neutron-1", denoms[0], denoms[1])
		if err != nil {
			panic(fmt.Sprintf("error seeding price cache for token %v: %v", symbol, err))
		}
		priceCache[symbol] = cachedPrice{
			price:     result.Price,
			timestamp: time.Now(),
		}
	}
}

func CalcAmountIn(usdAmountIn int, price float64, denom skip.ChainDenom) *big.Int {

	amountUSDBig := big.NewFloat(float64(usdAmountIn))
	amountBigToken := new(big.Float).Quo(
		amountUSDBig,
		big.NewFloat(price),
	)
	amountIn, _ := new(big.Float).Mul(
		amountBigToken,
		big.NewFloat(math.Pow(10, float64(denom.Decimals))),
	).Int(nil)
	return amountIn
}
