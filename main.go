package main

import (
    "fmt"
    "github.com/suren-m/go-demos-models/music"
)

func main() {
    var albums = []music.AlbumV2{
        {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
        {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
        {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
fmt.Println(albums)
}
