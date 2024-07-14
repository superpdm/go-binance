package binance

import (
	"context"
	"net/http"
)

// StartMarginStreamService create listen key for user stream service
type StartMarginStreamService struct {
	c *Client
}

// Do send request
func (s *StartMarginStreamService) Do(ctx context.Context, opts ...RequestOption) (listenKey string, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/margin/listen-key",
		secType:  secTypeAPIKey,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return "", err
	}
	j, err := newJSON(data)
	if err != nil {
		return "", err
	}
	listenKey = j.Get("listenKey").MustString()
	return listenKey, nil
}

// KeepaliveMarginStreamService update listen key
type KeepaliveMarginStreamService struct {
	c         *Client
	listenKey string
}

// ListenKey set listen key
func (s *KeepaliveMarginStreamService) ListenKey(listenKey string) *KeepaliveMarginStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *KeepaliveMarginStreamService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodPut,
		endpoint: "/sapi/v1/margin/listen-key",
		secType:  secTypeAPIKey,
	}
	r.setFormParam("listenKey", s.listenKey)
	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}

// CloseMarginStreamService delete listen key
type CloseMarginStreamService struct {
	c         *Client
	listenKey string
}

// ListenKey set listen key
func (s *CloseMarginStreamService) ListenKey(listenKey string) *CloseMarginStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *CloseMarginStreamService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/sapi/v1/margin/listen-key",
		secType:  secTypeAPIKey,
	}
	r.setFormParam("listenKey", s.listenKey)
	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}
