package main

import (
	"context"
	"log"
	"fmt"
	// "github.com/kr/pretty"

	"googlemaps.github.io/maps"
)

func main(){
	c, err := maps.NewClient(maps.WithAPIKey("api_key"))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.DirectionsRequest{
		Origin:      "Yokohana",
		Destination: "Nagoya",
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