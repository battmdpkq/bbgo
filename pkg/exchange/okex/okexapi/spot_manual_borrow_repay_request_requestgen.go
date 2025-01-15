// Code generated by "requestgen -method POST -responseType .APIResponse -responseDataField Data -url /api/v5/account/spot-manual-borrow-repay -type SpotManualBorrowRepayRequest -responseDataType []SpotBorrowRepayResponse -rateLimiter 1+20/2s"; DO NOT EDIT.

package okexapi

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/time/rate"
	"net/url"
	"reflect"
	"regexp"
)

var SpotManualBorrowRepayRequestLimiter = rate.NewLimiter(10, 1)

func (s *SpotManualBorrowRepayRequest) Currency(currency string) *SpotManualBorrowRepayRequest {
	s.currency = currency
	return s
}

func (s *SpotManualBorrowRepayRequest) Side(side MarginSide) *SpotManualBorrowRepayRequest {
	s.side = side
	return s
}

func (s *SpotManualBorrowRepayRequest) Amount(amount string) *SpotManualBorrowRepayRequest {
	s.amount = amount
	return s
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (s *SpotManualBorrowRepayRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (s *SpotManualBorrowRepayRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check currency field -> json key ccy
	currency := s.currency

	// assign parameter of currency
	params["ccy"] = currency
	// check side field -> json key side
	side := s.side

	// TEMPLATE check-valid-values
	switch side {
	case MarginSideBorrow, MarginSideRepay:
		params["side"] = side

	default:
		return nil, fmt.Errorf("side value %v is invalid", side)

	}
	// END TEMPLATE check-valid-values

	// assign parameter of side
	params["side"] = side
	// check amount field -> json key amount
	amount := s.amount

	// assign parameter of amount
	params["amount"] = amount

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (s *SpotManualBorrowRepayRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := s.GetParameters()
	if err != nil {
		return query, err
	}

	for _k, _v := range params {
		if s.isVarSlice(_v) {
			s.iterateSlice(_v, func(it interface{}) {
				query.Add(_k+"[]", fmt.Sprintf("%v", it))
			})
		} else {
			query.Add(_k, fmt.Sprintf("%v", _v))
		}
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (s *SpotManualBorrowRepayRequest) GetParametersJSON() ([]byte, error) {
	params, err := s.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (s *SpotManualBorrowRepayRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

func (s *SpotManualBorrowRepayRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (s *SpotManualBorrowRepayRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (s *SpotManualBorrowRepayRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (s *SpotManualBorrowRepayRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := s.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for _k, _v := range params {
		slugs[_k] = fmt.Sprintf("%v", _v)
	}

	return slugs, nil
}

// GetPath returns the request path of the API
func (s *SpotManualBorrowRepayRequest) GetPath() string {
	return "/api/v5/account/spot-manual-borrow-repay"
}

// Do generates the request object and send the request object to the API endpoint
func (s *SpotManualBorrowRepayRequest) Do(ctx context.Context) ([]SpotBorrowRepayResponse, error) {
	if err := SpotManualBorrowRepayRequestLimiter.Wait(ctx); err != nil {
		return nil, err
	}

	params, err := s.GetParameters()
	if err != nil {
		return nil, err
	}
	query := url.Values{}

	var apiURL string

	apiURL = s.GetPath()

	req, err := s.client.NewAuthenticatedRequest(ctx, "POST", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := s.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse APIResponse

	type responseUnmarshaler interface {
		Unmarshal(data []byte) error
	}

	if unmarshaler, ok := interface{}(&apiResponse).(responseUnmarshaler); ok {
		if err := unmarshaler.Unmarshal(response.Body); err != nil {
			return nil, err
		}
	} else {
		// The line below checks the content type, however, some API server might not send the correct content type header,
		// Hence, this is commented for backward compatibility
		// response.IsJSON()
		if err := response.DecodeJSON(&apiResponse); err != nil {
			return nil, err
		}
	}

	type responseValidator interface {
		Validate() error
	}

	if validator, ok := interface{}(&apiResponse).(responseValidator); ok {
		if err := validator.Validate(); err != nil {
			return nil, err
		}
	}
	var data []SpotBorrowRepayResponse
	if err := json.Unmarshal(apiResponse.Data, &data); err != nil {
		return nil, err
	}
	return data, nil
}
