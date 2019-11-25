package airvisual

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	baseURL        = "https://api.airvisual.com"
	getCityDataURL = baseURL + "/v2/city"
)

func GetCityData(apiKey string, country string, state string, city string) (*CityData, error) {
	url := getCityDataURL + "?city=" + city + "&state=" + state + "&country=" + country + "&key=" + apiKey
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	cityData := CityData{}
	if err := json.Unmarshal(body, &cityData); err != nil {
		return nil, err
	}

	return &cityData, nil
}
