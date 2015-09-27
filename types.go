package giphy

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