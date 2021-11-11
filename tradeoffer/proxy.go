package tradeoffer

import "net/http"

func (c *Client) SetTransport(t *http.Transport) {
	c.client.Transport = t
}
