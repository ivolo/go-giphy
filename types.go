package giphy

// SearchResponse is the response on the Giphy API search
type SearchResponse struct {
  Data       []Gif      `json:"data"`
  Meta       Meta       `json:"meta"`
  Pagination Pagination `json:"pagination"`
}

// TranslateResponse is the response on the Giphy API search
type TranslateResponse struct {
  Data       Gif  `json:"data"`
  Meta       Meta `json:"meta"`
}

// Meta represents the API responds
type Meta struct {
  Status int    `json:"status"`
  Msg    string `json:"msg"`
}

// Pagination allows you to paginate
type Pagination struct {
  TotalCount int `json:"total_count"`
  Count      int `json:"count"`
  Offset     int `json:"offset"`
}

// Gif is the standard Giphy Gif object
type Gif struct {
  Type              string `json:"type"`
  ID                string `json:"id"`
  URL               string `json:"url"`
  BitlyGifURL       string `json:"bitly_gif_url"`
  BitlyURL          string `json:"bitly_url"`
  EmbedURL          string `json:"embed_url"`
  Username          string `json:"username"`
  Source            string `json:"source"`
  Rating            string `json:"rating"`
  Caption           string `json:"caption"`
  ContentURL        string `json:"content_url"`
  ImportDatetime    string `json:"import_datetime"`
  TrendingDatetime  string `json:"trending_datetime"`
  Images            map[string]Image `json:"images"`
}

// Image is a specifically sized gif
type Image struct {
  Type      string `json:"type"`
  URL       string `json:"url"`
  Width     string `json:"width"`
  Height    string `json:"height"`
  Size      string `json:"size"`
  Mp4       string `json:"mp4"`
  Mp4Size   string `json:"mp4_size"`
  Webp      string `json:"webp"`
  WebpSize  string `json:"webp_size"`
}