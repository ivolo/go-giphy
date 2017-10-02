# go-giphy

 A go client for the [Giphy API](https://github.com/giphy/GiphyAPI).


## Quickstart

### [Search](https://developers.giphy.com/docs/#operation--gifs-search-get)

```go
import "github.com/ivolo/go-giphy"

c := giphy.New("dc6zaTOxFJmzC")
[]gifs, err = c.Search("simpsons ralph")

// gifs[25].URL -> http://giphy.com/gifs/ApEe3sVnOcHde
```

### [Translate](https://developers.giphy.com/docs/#operation--gifs-translate-get)

```go
import "github.com/ivolo/go-giphy"

c := giphy.New("dc6zaTOxFJmzC")
gif, err = c.Translate("simpsons ralph")

// gif.URL -> http://giphy.com/gifs/ApEe3sVnOcHde
```
