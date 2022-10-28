package adaptor

import (
	"encoding/json"
	"final-project/pkg/httpclient"
	"fmt"
)

type JSONBaseGet struct {
	Rajaongkir JSONRecordGet
}

type JSONRecordGet struct {
	Query   interface{}
	Status  interface{}
	Results interface{}
}

type JSONBasePost struct {
	Rajaongkir JSONRecordPost
}

type JSONRecordPost struct {
	Query               interface{}
	Status              interface{}
	Origin_details      interface{}
	Destination_details interface{}
	Results             interface{}
}

type RajaOngkirAdaptor struct {
	client *httpclient.Client
}

func NewRajaOngkirAdaptor(baseUrl string, apiKey string) *RajaOngkirAdaptor {
	client := httpclient.NewHttpClient(baseUrl, apiKey)

	return &RajaOngkirAdaptor{
		client: client,
	}
}

func (r *RajaOngkirAdaptor) GetCity(citycode string) (*JSONRecordGet, error) {
	path := fmt.Sprintf("city?id=%v", citycode)
	data, err := r.client.Get(path)
	if err != nil {
		return nil, err
	}
	var datas JSONBaseGet

	err = json.Unmarshal(data, &datas)
	if err != nil {
		return nil, err
	}

	//fmt.Println(datas)

	return &datas.Rajaongkir, err
}

func (r *RajaOngkirAdaptor) GetProvince(provincecode string) (*JSONRecordGet, error) {
	path := fmt.Sprintf("province?id=%v", provincecode)
	data, err := r.client.Get(path)
	if err != nil {
		return nil, err
	}
	var datas JSONBaseGet

	err = json.Unmarshal(data, &datas)
	if err != nil {
		return nil, err
	}

	//fmt.Println(datas)

	return &datas.Rajaongkir, err
}
