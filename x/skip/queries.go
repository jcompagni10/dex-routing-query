package skip

import (
	"bytes"
	"encoding/json"
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
