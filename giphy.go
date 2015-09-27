package giphy

import (
  "fmt"
  "net/http"
  "log"
  "encoding/json"
  "io/ioutil"
  "net/url"
)

// Endpoint
const (
  Endpoint = "http://api.giphy.com"
)

// a Giphy client
type Client struct {
  Verbose bool
  Key     string
}

// New returns a new Giphy client with the given API `key`.
func New(key string) *Client {
  c := &Client{ Key: key }
  return c
}

// Search the Giphy API for `query`.
func (c *Client) Search(query string) ([]Gif, error) {
  u := fmt.Sprintf("%s/v1/gifs/search?q=%s&api_key=%s", Endpoint, url.QueryEscape(query), c.Key)
  c.log("search '%s' -> GET %s", query, u)
  res, err := http.Get(u)
  if err != nil {
    return nil, err
  }
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return nil, err
  }
  if res.StatusCode != 200 {
    c.log("error response '%d' -> %s", res.StatusCode, string(body))
  }
  var s SearchResponse
  if err := json.Unmarshal(body, &s); err != nil {
    return nil, err
  }
  return s.Data, nil
}

// Log a `msg`.
func (c *Client) log(msg string, args ...interface{}) {
  if c.Verbose {
    log.Printf("giphy: " + msg, args)
  }
}