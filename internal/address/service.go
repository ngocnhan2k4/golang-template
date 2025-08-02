package address

import (
	"Template/configs"
	"Template/pkg/log"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Service interface {
	GetProvinces(ctx context.Context) string
	GetDistricts(ctx context.Context, provinceId string, depth int) string
	GetWards(ctx context.Context, districtId string, depth int) string
	GetCountries(ctx context.Context) string
}

type service struct {
	logger log.Logger
}

type Country struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
}

type GetCountriesResponse struct {
	Name string `json:"name"`
}

func NewService(logger log.Logger) Service {
	return &service{logger: logger}
}

func (s *service) GetProvinces(ctx context.Context) string {
	resp, err := http.Get(configs.AppConfig.AddressApi.BaseUrl + "/p")
	if err != nil {
		return "Error fetching provinces"
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)

	if err != nil {
		return "Error reading response"
	}
	return string(res)
}

func (s *service) GetDistricts(ctx context.Context, provinceId string, depth int) string {
	url := fmt.Sprintf("%s/p/%s?depth=%d", configs.AppConfig.AddressApi.BaseUrl, provinceId, depth)
	fmt.Print(url)
	resp, err := http.Get(url)
	if err != nil {
		return "Error fetching districts"
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Error reading response"
	}
	return string(res)
}

func (s *service) GetWards(ctx context.Context, districtId string, depth int) string {
	url := fmt.Sprintf("%s/d/%s?depth=%d", configs.AppConfig.AddressApi.BaseUrl, districtId, depth)
	resp, err := http.Get(url)
	if err != nil {
		return "Error fetching wards"
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Error reading response"
	}
	return string(res)
}

func (s *service) GetCountries(ctx context.Context) string {
	resp, err := http.Get(configs.AppConfig.AddressApi.CountriesUrl)
	if err != nil {
		return "Error fetching countries"
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Error reading response"
	}
	var countries []Country
	if err := json.Unmarshal(res, &countries); err != nil {
		return "Error unmarshalling countries response"
	}
	var countriesResponse []GetCountriesResponse
	for _, country := range countries {
		countriesResponse = append(countriesResponse, GetCountriesResponse{Name: country.Name.Common})
	}
	jsonData, err := json.Marshal(countriesResponse)
	if err != nil {
		return "Error marshalling countriesResponse response"
	}
	return string(jsonData)
}
