package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	albumId := flag.String("albumId", "1", "Id of the album")
	flag.Parse()
	fmt.Println("passed in args:", *albumId)

	id, _ := strconv.Atoi(*albumId)
	printAlbumResult(id)

	// var albums = []music.AlbumV2{
	// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	// }

	// for i := range albums {
	// 	if albums[i].ID == *albumId {
	// 		fmt.Println("Result:", albums[i])
	// 		break
	// 	}
	// }

	// var movie = movies.Movie{
	// 	Title: "Inception",
	// }
	// fmt.Println(movie)

}
