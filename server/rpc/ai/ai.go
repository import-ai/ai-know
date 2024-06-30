package ai

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/rs/zerolog/log"
	"github.com/ycdzj/shuinotes/server/config"
)

type Doc struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func PostDoc(docID string, doc *Doc) error {
	u, err := url.Parse(config.AIServerAddr() + "/api/index")
	if err != nil {
		return err
	}
	u = u.JoinPath("/" + docID)

	reqBody, err := json.Marshal(doc)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, u.String(), bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Info().
		Str("url", u.String()).
		Str("resp_body", string(respBody)).
		Int("status", resp.StatusCode).
		Send()
	return nil
}
