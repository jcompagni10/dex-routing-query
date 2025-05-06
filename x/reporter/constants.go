package reporter

import (
	"os"
	"strings"

	"github.com/jcompagni10/skip-router-data/x/skip"
	log "github.com/sirupsen/logrus"
)

var (
	Amounts    = []int{1, 50, 100, 1000, 5000}
	Pairs      = [][]string{}
	ChainIds   = []string{}
	Exclusions = map[string][]string{}
)

var ChainData = map[string]skip.ChainAssets{}

func ParsePairsFromEnv() {
	pairs := strings.Split(os.Getenv("PAIRS"), ";")
	if len(pairs) == 0 {
		panic("PAIRS is not set")
	}

	for _, pair := range pairs {
		split := strings.Split(pair, ",")
		Pairs = append(Pairs, split)
	}

	log.Info("Parsed pairs from env: ", Pairs)

}

func ParseChainIdsFromEnv() {
	chainIds := strings.Split(os.Getenv("CHAIN_IDS"), ";")
	if len(chainIds) == 0 {
		panic("CHAIN_IDS is not set")
	}

	ChainIds = chainIds

	log.Info("Parsed chain ids from env: ", ChainIds)
}

func ParseExclusionsFromEnv() {
	exclusions := strings.Split(os.Getenv("PAIR_EXCLUSIONS"), ";")
	if len(exclusions) == 0 {
		log.Info("No exclusions set")
		return
	}

	for _, exclusion := range exclusions {
		split := strings.Split(exclusion, ":")
		Exclusions[split[0]] = strings.Split(split[1], ",")
	}

	log.Info("Parsed exclusions from env: ", Exclusions)
}
