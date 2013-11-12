package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/golang/groupcache"
)

func main() {
	me := "127.0.0.1:8080"
	peers := groupcache.NewHTTPPool("http://" + me)

	// Whenever peers change:
	//peers.Set("http://10.0.0.1", "http://10.0.0.2", "http://10.0.0.3")

	peers.SetBasePath("/cache/")

	getter := groupcache.GetterFunc(func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
		// key == gopher.png
		dest.SetString(strconv.Itoa(0) + ":" + key)
		return nil
	})
	groupcache.NewGroup("thumnail", 1<<20, getter)

	log.Fatal(http.ListenAndServe(me, peers))
}
