package skip

import (
	"bytes"
	"encoding/json"
	"net/url"
)

func GetFungibleRoutes(req *FungibleRouteRequest) (*FungibleRouteResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return &FungibleRouteResponse{}, err
	}

	res, err := PostRequest("/fungible/route", bytes.NewBuffer(body))
	if err != nil {
		return &FungibleRouteResponse{}, err
	}

	var response FungibleRouteResponse

	err = json.Unmarshal(res, &response)
	if err != nil {
		return &FungibleRouteResponse{}, err
	}

	return &response, nil
}

func GetChainAssets(chainIDs []string) (map[string]ChainAssets, error) {
	queryParams := url.Values{}
	for _, chainID := range chainIDs {
		queryParams.Add("chain_ids", chainID)
	}

	res, err := GetRequest("/fungible/assets", queryParams)
	if err != nil {
		return map[string]ChainAssets{}, err
	}

	var response ChainToAssetsResponse

	err = json.Unmarshal(res, &response)
	if err != nil {
		return map[string]ChainAssets{}, err
	}

	return response.ChainToAssetsMap, nil
}
