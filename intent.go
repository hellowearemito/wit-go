package witai

import (
	"encoding/json"
	"net/http"
)

// Intent - https://wit.ai/docs/http/20200513#post__intents_link
type Intent struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c *Client) GetIntents() ([]Intent, error) {
	resp, err := c.request(http.MethodGet, "/intents", "application/json", nil)
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var intents []Intent
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&intents)
	return intents, err

}
