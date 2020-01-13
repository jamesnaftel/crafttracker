package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type server struct {
	beers []untappd
}

func main() {

	list, err := readUntappdData("../data/untappd.json")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	s := server{beers: list}

	// API
	http.HandleFunc("/list", s.handlerList)
	http.ListenAndServe(":8085", nil)
}

func (s *server) handlerList(rw http.ResponseWriter, r *http.Request) {

	b, err := json.Marshal(s.beers)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}
	rw.Write(b)
}

func readUntappdData(file string) ([]untappd, error) {

	var b []untappd
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &b)

	return b, nil
}

type untappd struct {
	BeerName                  string      `json:"beer_name"`
	BreweryName               string      `json:"brewery_name"`
	BeerType                  string      `json:"beer_type"`
	BeerAbv                   string      `json:"beer_abv"`
	BeerIbu                   string      `json:"beer_ibu"`
	Comment                   string      `json:"comment"`
	VenueName                 string      `json:"venue_name"`
	VenueCity                 string      `json:"venue_city"`
	VenueState                string      `json:"venue_state"`
	VenueCountry              string      `json:"venue_country"`
	VenueLat                  string      `json:"venue_lat"`
	VenueLng                  string      `json:"venue_lng"`
	RatingScore               string      `json:"rating_score"`
	CreatedAt                 string      `json:"created_at"`
	CheckinURL                string      `json:"checkin_url"`
	BeerURL                   string      `json:"beer_url"`
	BreweryURL                string      `json:"brewery_url"`
	BreweryCountry            string      `json:"brewery_country"`
	BreweryCity               string      `json:"brewery_city"`
	BreweryState              string      `json:"brewery_state"`
	FlavorProfiles            string      `json:"flavor_profiles"`
	PurchaseVenue             string      `json:"purchase_venue"`
	ServingType               string      `json:"serving_type"`
	CheckinID                 string      `json:"checkin_id"`
	Bid                       string      `json:"bid"`
	BreweryID                 string      `json:"brewery_id"`
	PhotoURL                  interface{} `json:"photo_url"`
	GlobalRatingScore         float64     `json:"global_rating_score"`
	GlobalWeightedRatingScore float64     `json:"global_weighted_rating_score"`
	TaggedFriends             string      `json:"tagged_friends"`
	TotalToasts               string      `json:"total_toasts"`
	TotalComments             string      `json:"total_comments"`
}
