package giphy

import "fmt"
import "net/http"
import "log"
import "encoding/json"
import "io/ioutil"
import "net/url"

// Endpoint
const (
  Endpoint = "http://api.giphy.com"
)

// a Giphy client
type Client struct {
  Verbose bool
  Key     string
}

type Image struct {
  Type      string
  Url       string 
  Width     string
  Height    string
  Size      string
  Mp4       string
  Mp4_size  string
  webp      string
  webp_size string
}

type SearchResponse struct {  
    Data       []Gif      `json:"data"`
    Meta       Meta       `json:"meta"`
    Pagination Pagination `json:"pagination"`
}

type Pagination struct {
  TotalCount int `json:"total_count"`
  Count      int `json:"count"`
  Offset     int `json:"offset"`
}

type Meta struct {
  Status int    `json:"status"`
  Msg    string `json:"msg"`
}

// A Giphy GIF object
type Gif struct {
  Type              string
  Id                string
  Url               string
  Bitly_Gif_Url     string
  Bitly_Url         string
  Embed_Url         string
  Username          string
  Source            string
  Rating            string
  Caption           string
  Content_Url       string
  Import_Datetime   string
  Trending_Datetime string
  Images            map[string]Image
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