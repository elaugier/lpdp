package proxy

import (
	"net/http/httputil"
	"net/url"
)

//Prox ...
type Prox struct {
	target *url.URL
	proxy  *httputil.ReverseProxy
}

//NewProxy ...
func NewProxy(target string) *Prox {
	url, _ := url.Parse(target)
	return &Prox{target: url, proxy: httputil.NewSingleHostReverseProxy(url)}
}
