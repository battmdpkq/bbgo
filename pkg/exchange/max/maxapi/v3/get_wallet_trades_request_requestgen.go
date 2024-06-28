// Code generated by "requestgen -debug -method GET -url /api/v3/wallet/:walletType/new/trades -type GetWalletTradesRequest -responseType []"github.com/c9s/bbgo/pkg/exchange/max/maxapi/v3".Trade"; DO NOT EDIT.

package v3

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

func (g *GetWalletTradesRequest) Market(market string) *GetWalletTradesRequest {
	g.market = market
	return g
}

func (g *GetWalletTradesRequest) Timestamp(timestamp time.Time) *GetWalletTradesRequest {
	g.timestamp = &timestamp
	return g
}

func (g *GetWalletTradesRequest) FromID(fromID uint64) *GetWalletTradesRequest {
	g.fromID = &fromID
	return g
}

func (g *GetWalletTradesRequest) Order(order string) *GetWalletTradesRequest {
	g.order = &order
	return g
}

func (g *GetWalletTradesRequest) Limit(limit uint64) *GetWalletTradesRequest {
	g.limit = &limit
	return g
}

func (g *GetWalletTradesRequest) WalletType(walletType WalletType) *GetWalletTradesRequest {
	g.walletType = walletType
	return g
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (g *GetWalletTradesRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (g *GetWalletTradesRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check market field -> json key market
	market := g.market

	// TEMPLATE check-required
	if len(market) == 0 {
		return nil, fmt.Errorf("market is required, empty string given")
	}
	// END TEMPLATE check-required

	// assign parameter of market
	params["market"] = market
	// check timestamp field -> json key timestamp
	if g.timestamp != nil {
		timestamp := *g.timestamp

		// assign parameter of timestamp
		// convert time.Time to milliseconds time stamp
		params["timestamp"] = strconv.FormatInt(timestamp.UnixNano()/int64(time.Millisecond), 10)
	} else {
	}
	// check fromID field -> json key from_id
	if g.fromID != nil {
		fromID := *g.fromID

		// assign parameter of fromID
		params["from_id"] = fromID
	} else {
	}
	// check order field -> json key order
	if g.order != nil {
		order := *g.order

		// TEMPLATE check-valid-values
		switch order {
		case "asc", "desc":
			params["order"] = order

		default:
			return nil, fmt.Errorf("order value %v is invalid", order)

		}
		// END TEMPLATE check-valid-values

		// assign parameter of order
		params["order"] = order
	} else {
	}
	// check limit field -> json key limit
	if g.limit != nil {
		limit := *g.limit

		// assign parameter of limit
		params["limit"] = limit
	} else {
	}

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (g *GetWalletTradesRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := g.GetParameters()
	if err != nil {
		return query, err
	}

	for _k, _v := range params {
		if g.isVarSlice(_v) {
			g.iterateSlice(_v, func(it interface{}) {
				query.Add(_k+"[]", fmt.Sprintf("%v", it))
			})
		} else {
			query.Add(_k, fmt.Sprintf("%v", _v))
		}
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (g *GetWalletTradesRequest) GetParametersJSON() ([]byte, error) {
	params, err := g.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (g *GetWalletTradesRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check walletType field -> json key walletType
	walletType := g.walletType

	// TEMPLATE check-required
	if len(walletType) == 0 {
		return nil, fmt.Errorf("walletType is required, empty string given")
	}
	// END TEMPLATE check-required

	// assign parameter of walletType
	params["walletType"] = walletType

	return params, nil
}

func (g *GetWalletTradesRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (g *GetWalletTradesRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (g *GetWalletTradesRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (g *GetWalletTradesRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := g.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for _k, _v := range params {
		slugs[_k] = fmt.Sprintf("%v", _v)
	}

	return slugs, nil
}

func (g *GetWalletTradesRequest) Do(ctx context.Context) ([]Trade, error) {

	// empty params for GET operation
	var params interface{}
	query, err := g.GetParametersQuery()
	if err != nil {
		return nil, err
	}

	apiURL := "/api/v3/wallet/:walletType/new/trades"
	slugs, err := g.GetSlugsMap()
	if err != nil {
		return nil, err
	}

	apiURL = g.applySlugsToUrl(apiURL, slugs)

	req, err := g.client.NewAuthenticatedRequest(ctx, "GET", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := g.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse []Trade
	if err := response.DecodeJSON(&apiResponse); err != nil {
		return nil, err
	}
	return apiResponse, nil
}
