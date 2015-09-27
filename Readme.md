
# go-gighy

 [Giphy](http://giphy.com/gifs/reaction-the-simpsons-5WQTGtSepBela) [API](https://github.com/giphy/GiphyAPI) from golang.

![](http://i.giphy.com/5WQTGtSepBela.gif)

## Quickstart

```go
import "github.com/ivolo/go-giphy"

c := giphy.New("dc6zaTOxFJmzC")
[]gifs, err = c.Search("simpsons ralph")

// gifs[25].Url -> http://giphy.com/gifs/ApEe3sVnOcHde
```

![](http://i.giphy.com/ApEe3sVnOcHde.gif)