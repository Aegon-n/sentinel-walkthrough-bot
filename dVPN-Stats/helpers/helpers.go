package helpers

import(
	"encoding/json"
	"log"
	"net/http"
	"github.com/Aegon-n/sentinel-bot/dVPN-Stats/models"
	"github.com/Aegon-n/sentinel-bot/dVPN-Stats/constants"
)

func GetCount(statType, field string) (int, int, error) {
	var active models.Active
	var average models.Average
	var url string
	if statType == "active" {
		if field == "nodes" {
			url = constants.MasterNodestatsUrl + "nodes/active?interval=current"
		} else {
			url = constants.MasterNodestatsUrl + "sessions/active?filter=lastday&format=count"
		}
		resp, err := http.Get(url)
		if err != nil {
			log.Println("Unable get active nodes")
			return 0, 0, err
		}
		if err := json.NewDecoder(resp.Body).Decode(&active); err != nil {
			log.Println("unable to decode response")
			return 0, 0, err
		}
		defer resp.Body.Close()
		return active.Count, average.Average, nil

	} 
	if field == "nodes" {
		url = constants.MasterNodestatsUrl + "nodes/active?interval=day&format=count"
	} else {
		url = constants.MasterNodestatsUrl + "sessions/average?interval=day&format=count"
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Unable get active nodes")
		return 0, 0, err
	}
	if err := json.NewDecoder(resp.Body).Decode(&average); err != nil {
		log.Println("unable to decode response")
		return 0, 0, err
	}
	defer resp.Body.Close()
	return active.Count, average.Average, nil
}

func GetUsedBandwidth(filter string) (float64, error){
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
		return bandwidth.Stats, err
	}
	if err := json.NewDecoder(resp.Body).Decode(&bandwidth); err != nil {
		log.Println("unable to decode response")
		return bandwidth.Stats, err
	}
	defer resp.Body.Close()
	return bandwidth.Stats, nil
}
