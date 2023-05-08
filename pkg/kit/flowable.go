package kit

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type Flowable struct {
	baseURL  string
	username string
	password string
	client   *http.Client
}

type Task struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Assignee    string `json:"assignee"`
}

func NewFlowable(baseURL, username, password string) *Flowable {
	return &Flowable{
		baseURL:  baseURL,
		username: username,
		password: password,
		client:   &http.Client{},
	}
}

func (f *Flowable) DoRequest(method, path string, body interface{}) (*http.Response, error) {
	url := f.baseURL + path

	var buf bytes.Buffer
	if body != nil {
		err := json.NewEncoder(&buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, &buf)
	if err != nil {
		return nil, err
	}

	auth := f.username + ":" + f.password
	base64Auth := base64.StdEncoding.EncodeToString([]byte(auth))
	authHeader := "Basic " + base64Auth
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/json")

	resp, err := f.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}

	return resp, nil
}

