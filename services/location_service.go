package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/syohex/go-ipinfo"
)

type locationService struct {
}

// LocationService : Service For Location
func LocationService() *locationService {
	return &locationService{}
}

// GetCountryCode : Get Country Code of current location by using ipinfo api
func (locationService *locationService) GetCountryCode() string {
	ip := locationService.getOutboundIP()
	info := ipinfo.IPInfo(net.ParseIP(ip))

	return info.Country
}

type httpbinResponse struct {
	Origin string `json:"origin"`
}

func (locationService *locationService) getOutboundIP() string {
	res, err := http.Get("http://httpbin.org/ip")
	if err != nil {
		log.Fatal(err)
		return ""
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Printf("StatusCode=%d", res.StatusCode)
		return ""
	}

	var response httpbinResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&response)
	if err != nil {
		log.Fatal(err)
	}

	return response.Origin
}
