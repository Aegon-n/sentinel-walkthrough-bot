package helpers

import(
	"encoding/json"
	"log"
	"net/http"
	"github.com/Aegon-n/sentinel-bot/dVPN-Stats/models"
	"github.com/Aegon-n/sentinel-bot/dVPN-Stats/constants"
)

func GetCount(statType, field string, ch chan<-int)  {
	var active models.Active
	var average models.Average
	var url string
	if statType == "active" {
		if field == "nodes" {
			url = constants.MasterNodestatsUrl + "nodes/active?interval=current"
		} else {
			url = constants.MasterNodestatsUrl + "sessions/active?filter=lastday&format=count"
		}
		resp, _ := http.Get(url)
		
		if err := json.NewDecoder(resp.Body).Decode(&active); err != nil {
			log.Println("unable to decode response")
			ch <- 0
		}
		defer resp.Body.Close()
		ch <- active.Count
		return
	} 
	if field == "nodes" {
		url = constants.MasterNodestatsUrl + "nodes/active?interval=day&format=count"
	} else {
		url = constants.MasterNodestatsUrl + "sessions/average?interval=day&format=count"
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Unable get active nodes")
		ch <- 0
	}
	if err := json.NewDecoder(resp.Body).Decode(&average); err != nil {
		log.Println("unable to decode response")
		ch <- 0
		return
	}
	defer resp.Body.Close()
	ch <- int(average.Average)
}

func GetUsedBandwidth(filter string, ch chan<-float64) {
	var bandwidth models.Bandwidth
	var url string
  if filter == "lastday" {
		url = constants.MasterNodestatsUrl + "bandwidth?filter=lastday&format=count"
	} else {
		url = constants.MasterNodestatsUrl + "bandwidth/all?format=count"
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Unable to get bandwidth stats")
		ch <- 0.0
	}
	if err := json.NewDecoder(resp.Body).Decode(&bandwidth); err != nil {
		log.Println("unable to decode response")
		ch <- 0.0
	}
	defer resp.Body.Close()
	ch <- bandwidth.Stats
}
