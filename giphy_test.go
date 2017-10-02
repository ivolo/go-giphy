package giphy

import "fmt"
import "testing" 

func setup() *Client {
  return &Client{ Key: "dc6zaTOxFJmzC", Verbose: true }
}

func TestSearch(t *testing.T) {
  c := setup()
  gifs, err := c.Search("simpsons ralph")
  if err != nil {
    t.Fatal("error searching:", err)
  }
  if len(gifs) == 0 {
    t.Fatal("no gifs returned")
  }

  t.Log(fmt.Sprintf("found %d gifs", len(gifs)))
  for i, gif := range(gifs) {
    t.Log(fmt.Sprintf("gif %d: %s", i, gif.URL))
  }
}

func TestTranslate(t *testing.T) {
  c := setup()
  gif, err := c.Translate("simpsons ralph")
  if err != nil {
    t.Fatal("error translating:", err)
  }

  t.Log(fmt.Sprintf("gif: %s", gif.URL))
}