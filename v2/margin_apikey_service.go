package binance

import (
	"context"
	"net/http"
)

// CreateMarginApiKeyService apply for a loan
type CreateMarginApiKeyService struct {
	c         *Client
	apiName   string
	symbol    *string
	ip        *string
	publicKey *string
}

func (s *CreateMarginApiKeyService) ApiName(apiName string) *CreateMarginApiKeyService {
	s.apiName = apiName
	return s
}

func (s *CreateMarginApiKeyService) Symbol(symbol string) *CreateMarginApiKeyService {
	s.symbol = &symbol
	return s
}

func (s *CreateMarginApiKeyService) Ip(ip string) *CreateMarginApiKeyService {
	s.ip = &ip
	return s
}

func (s *CreateMarginApiKeyService) PublicKey(publicKey string) *CreateMarginApiKeyService {
	s.publicKey = &publicKey
	return s
}

func (s *CreateMarginApiKeyService) Do(ctx context.Context, opts ...RequestOption) (res *CreateMarginApiKeyResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/margin/apiKey",
		secType:  secTypeSigned,
	}
	m := params{
		"apiName": s.apiName,
	}
	if s.symbol != nil {
		m["symbol"] = *s.symbol
	}
	if s.ip != nil {
		m["ip"] = *s.ip
	}
	if s.publicKey != nil {
		m["pubKey"] = *s.publicKey
	}
	r.setFormParams(m)

	res = new(CreateMarginApiKeyResponse)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type CreateMarginApiKeyResponse struct {
	ApiKey    string `json:"apiKey"`
	SecretKey string `json:"secretKey"`
	Type      string `json:"type"`
}

type DeleteMarginApiKeyService struct {
	c       *Client
	apiKey  *string
	apiName *string
	symbol  *string
}

func (s *DeleteMarginApiKeyService) ApiKey(apiKey string) *DeleteMarginApiKeyService {
	s.apiKey = &apiKey
	return s
}
func (s *DeleteMarginApiKeyService) ApiName(apiName string) *DeleteMarginApiKeyService {
	s.apiName = &apiName
	return s
}
func (s *DeleteMarginApiKeyService) Symbol(symbol string) *DeleteMarginApiKeyService {
	s.symbol = &symbol
	return s
}
func (s *DeleteMarginApiKeyService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/sapi/v1/margin/apiKey",
		secType:  secTypeSigned,
	}
	m := params{}
	if s.apiKey != nil {
		m["apiKey"] = *s.apiKey
	}
	if s.apiName != nil {
		m["apiName"] = *s.apiName
	}
	if s.symbol != nil {
		m["symbol"] = *s.symbol
	}
	r.setFormParams(m)
	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}

type MarginApiKeyEditIpService struct {
	c      *Client
	apiKey string
	symbol *string
	ip     string
}

func (s *MarginApiKeyEditIpService) ApiKey(apiKey string) *MarginApiKeyEditIpService {
	s.apiKey = apiKey
	return s
}

func (s *MarginApiKeyEditIpService) Symbol(symbol string) *MarginApiKeyEditIpService {
	s.symbol = &symbol
	return s
}
func (s *MarginApiKeyEditIpService) Ip(ip string) *MarginApiKeyEditIpService {
	s.ip = ip
	return s
}

func (s *MarginApiKeyEditIpService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodPut,
		endpoint: "/sapi/v1/margin/apiKey/ip",
		secType:  secTypeSigned,
	}
	m := params{
		"apiKey": s.apiKey,
		"ip":     s.ip,
	}
	if s.symbol != nil {
		m["symbol"] = *s.symbol
	}
	r.setFormParams(m)
	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}

type QueryMarginApiKeyService struct {
	c      *Client
	apiKey string
	symbol *string
}

func (s *QueryMarginApiKeyService) ApiKey(apiKey string) *QueryMarginApiKeyService {
	s.apiKey = apiKey
	return s
}
func (s *QueryMarginApiKeyService) Symbol(symbol string) *QueryMarginApiKeyService {
	s.symbol = &symbol
	return s
}
func (s *QueryMarginApiKeyService) Do(ctx context.Context, opts ...RequestOption) (res *QueryMarginApiKeyResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/apiKey",
		secType:  secTypeSigned,
	}
	r.setParam("apiKey", s.apiKey)
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	res = new(QueryMarginApiKeyResponse)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type QueryMarginApiKeyResponse struct {
	ApiKey  string `json:"apiKey"`
	Ip      string `json:"ip"`
	ApiName string `json:"apiName"`
	Type    string `json:"type"`
}

type ListMarginApiKeyService struct {
	c      *Client
	symbol *string
}

func (s *ListMarginApiKeyService) Symbol(symbol string) *ListMarginApiKeyService {
	s.symbol = &symbol
	return s
}
func (s *ListMarginApiKeyService) Do(ctx context.Context, opts ...RequestOption) (res []*QueryMarginApiKeyResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/api-key-list",
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
