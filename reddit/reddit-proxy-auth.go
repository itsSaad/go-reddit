package reddit

import "net/http"

// Sets the User-Agent header for requests.
// We need to set a custom user agent because using the one set by the
// stdlib gives us 429 Too Many Requests responses from the Reddit API.
type authorizationTransport struct {
	Bearer string
	Base   http.RoundTripper
}

func (t *authorizationTransport) setAuthorization(req *http.Request) *http.Request {
	req2 := cloneRequest(req)
	req2.Header.Set(headerAuthorization, t.Bearer)
	return req2
}

func (t *authorizationTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req2 := t.setAuthorization(req)
	return t.base().RoundTrip(req2)
}

func (t *authorizationTransport) base() http.RoundTripper {
	if t.Base != nil {
		return t.Base
	}
	return http.DefaultTransport
}
