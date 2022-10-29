package view

import (
	"encoding/json"
	"final-project/adaptor"
	"final-project/server/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AllUsers struct {
	Id       string      `json:"id"`
	Fullname string      `json:"fullname"`
	Address  interface{} `json:"address"`
	Auth     interface{} `json:"auth"`
}

type Address struct {
	City     AddressDetail `json:"city"`
	Province AddressDetail `json:"province"`
	Street   string        `json:"street"`
}

type AddressDetail struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewAllUsers(users *[]model.User, rajaOngkir *adaptor.RajaOngkirAdaptor) ([]AllUsers, error) {
	var newUsers []AllUsers

	for _, user := range *users {
		detailAddress, err := getDetailAddress(&user, rajaOngkir)
		if err != nil {
			return nil, err
		}
		newUsers = append(newUsers, AllUsers{
			Id:       user.Id,
			Fullname: user.Fullname,
			Auth:     gin.H{"email": user.Email},
			Address:  detailAddress,
		})
	}

	return newUsers, nil
}

func getDetailAddress(user *model.User, rajaOngkir *adaptor.RajaOngkirAdaptor) (*Address, error) {
	city, err := rajaOngkir.GetCity(user.CityId)
	if err != nil {
		log.Printf("Error when get city from adaptor %v\n", err)
		return nil, err
	}

	var jsonCity map[string]interface{}
	err = json.Unmarshal(city, &jsonCity)
	if err != nil {
		log.Printf("Error when unmarshal city from adaptor %v\n", err)
		return nil, err
	}

	resp := jsonCity["rajaongkir"].(map[string]interface{})
	result := resp["results"].(map[string]interface{})

	return &Address{
		Street: user.Street,
		City: AddressDetail{
			Id:   result["city_id"].(string),
			Name: result["city_name"].(string),
		},
		Province: AddressDetail{
			Id:   result["province_id"].(string),
			Name: result["province"].(string),
		},
	}, nil
}

func SuccessAllUsersResponse(query *model.Pagination, payload interface{}) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: "GET_ALL_USERS_SUCCESS",
		Payload: payload,
		Query:   query,
	}
}
