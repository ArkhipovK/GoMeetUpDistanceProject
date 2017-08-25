// Copyright 2017 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package main contains a simple command line tool for Static Maps API
// Documentation: https://developers.google.com/maps/documentation/static-maps/
package main


import (
	"flag"
	"fmt"
	"log"
	"os"
	"net/http"
	"encoding/json"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
	"github.com/rs/cors"
)


func usageAndExit(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	fmt.Println("Flags:")
	flag.PrintDefaults()
	os.Exit(2)
}

func check(err error) {
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
}

func main() {

	mux := http.NewServeMux()
    mux.HandleFunc("/", handler)
    handler := cors.Default().Handler(mux)
    http.ListenAndServe(":8080", handler)

}
func handler(w http.ResponseWriter, req *http.Request) {
	
	keys:= req.URL.Query()
    
    if len(keys) < 2 {
        return
    }

    
	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyBfiYm6_SfxdPuH173UQQjBRutC6gyPmVA"))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.DirectionsRequest{
		Origin:      keys.Get("start"),
		Destination: keys.Get("end"),
	}
	routes, waypoints, err := c.Directions(context.Background(), r)

	resp:= map[string]interface{}{
		"geocoded_waypoints":waypoints,
		"routes":routes,
	}
	if err != nil {
		 http.Error(w, err.Error(), http.StatusBadRequest)
    		return
	}
	
	json.NewEncoder(w).Encode(resp)
}

