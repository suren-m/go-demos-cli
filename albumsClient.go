package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/suren-m/go-demos-models/music"
)

const _api = "http://localhost:8080/albums"

func printAlbumResult(id int) {
	url := fmt.Sprintf("%s/%s", _api, strconv.Itoa(id))
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var album music.Album
	// deserialize json
	json.Unmarshal(body, &album)

	// print struct with fields info
	log.Printf("%+v\n", album.ID)

	// serialize and print jsondata
	jsondata, _ := json.MarshalIndent(album, "", "    ")
	log.Printf("%s\n", jsondata)

}
