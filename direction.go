package main

import (
	"context"
	"log"
	"fmt"
	"github.com/go-ini/ini"

	"googlemaps.github.io/maps"
)

var (
	apikey string
)

func init(){
	cfg,err := ini.Load("google.ini")
	if err != nil{
		log.Println("this is not ini")
	}
	apikey=cfg.Section("googleApi").Key("Api_key").String()
}

func main(){
	log.Printf("%s\n",apikey)
	GetGoogleApi("Kasaoka","Nanba")
}

func GetGoogleApi(origin ,destination string){
	c, err := maps.NewClient(maps.WithAPIKey(apikey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.DirectionsRequest{
		Origin:      origin,
		Destination: destination,
	}
	route,_, err := c.Directions(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	Print(route)
}

func Print(r []maps.Route){	
	 
	for _,rou := range r {
		for _,leg := range rou.Legs	{
			for _,step := range leg.Steps{
				fmt.Printf("%s\n",step.HTMLInstructions)
				fmt.Printf("Sistance:%s\n",step.Distance.HumanReadable)
				fmt.Printf("StartLocation:%f\n",step.StartLocation)
				fmt.Printf("EndLocation:%f\n",step.EndLocation)
			}
		}
	}
}