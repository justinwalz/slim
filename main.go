package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/justinwalz/golang-soundcloud/soundcloud"
)

func main() {
	var client = flag.String("c", "", "soundcloud client id")
	var uid = flag.Uint64("u", 0, "soundcloud user id")
	var debug = flag.Bool("d", false, "debug")

	flag.Parse()

	if *client == "" {
		error("client id required")
	}

	if *uid == 0 {
		error("user id required")
	}

	unauthenticatedAPI := &soundcloud.Api{
		ClientId: *client,
	}

	tracks, err := unauthenticatedAPI.User(*uid).AllFavorites()
	if err != nil {
		if *debug {
			panic(err)
		}
		os.Exit(1)
	}

	for _, track := range tracks {
		fmt.Println(track.PermalinkUrl)
	}
}

func printUsage() {
	fmt.Println("usage: slim -c <client-id> -u <user-id>")
}

func error(i interface{}) {
	printUsage()
	fmt.Printf("error: %v\n", i)
	os.Exit(1)
}
