
# go-giphy

 A go client for the [Giphy](http://giphy.com/gifs/reaction-the-simpsons-5WQTGtSepBela) [API](https://github.com/giphy/GiphyAPI).

![](http://i.giphy.com/5WQTGtSepBela.gif)

It's also my first Go library. :octopus:

## Quickstart

#### Search

```go
import "github.com/ivolo/go-giphy"

c := giphy.New("dc6zaTOxFJmzC")
[]gifs, err = c.Search("simpsons ralph")

// gifs[25].URL -> http://giphy.com/gifs/ApEe3sVnOcHde
```

![](http://i.giphy.com/ApEe3sVnOcHde.gif)