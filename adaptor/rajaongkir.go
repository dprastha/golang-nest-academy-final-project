package adaptor

import (
	"final-project/pkg/httpclient"
	"fmt"
)

type RajaOngkirAdaptor struct {
	client *httpclient.Client
}

type RajaOngkirResponse struct {
	Rajaongkir ResultResponse `json:"rajaongkir"`
}

type ResultResponse struct {
	Result CourierService `json:"result"`
}

type CourierService struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	CourierCost *[]CourierCost
}

type CourierCost struct {
	Service     string `json:"service"`
	Description string `json:"description"`
	ServiceCost *[]ServiceCost
}

type ServiceCost struct {
	Value int    `json:"value"`
	Etd   string `json:"etd"`
	Note  string `json:"note"`
}

func NewRajaOngkirAdaptor(baseUrl string, apiKey string) *RajaOngkirAdaptor {
	client := httpclient.NewHttpClient(baseUrl, apiKey)

	return &RajaOngkirAdaptor{
		client: client,
	}
}

func (r *RajaOngkirAdaptor) GetCity(citycode string) ([]byte, error) {
	path := fmt.Sprintf("city?id=%v", citycode)
	data, err := r.client.Get(path)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (r *RajaOngkirAdaptor) GetProvince(provincecode string) ([]byte, error) {
	path := fmt.Sprintf("province?id=%v", provincecode)
	data, err := r.client.Get(path)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (r *RajaOngkirAdaptor) PostCost(payload interface{}) ([]byte, error) {
	path := "cost"
	data, err := r.client.Post(path, payload)

	if err != nil {
		return nil, err
	}

	return data, err
}
