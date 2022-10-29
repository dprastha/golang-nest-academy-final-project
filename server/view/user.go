package view

import (
	"encoding/json"
	"final-project/adaptor"
	"final-project/server/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DetailUsers struct {
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

func NewAllUsers(users *[]model.User, rajaOngkir *adaptor.RajaOngkirAdaptor) ([]DetailUsers, error) {
	var newUsers []DetailUsers

	for _, user := range *users {
		detailAddress, err := getDetailAddress(&user, rajaOngkir)
		if err != nil {
			return nil, err
		}
		newUsers = append(newUsers, DetailUsers{
			Id:       user.Id,
			Fullname: user.Fullname,
			Auth:     gin.H{"email": user.Email},
			Address:  detailAddress,
		})
	}

	return newUsers, nil
}

func NewUsers(user *model.User, rajaOngkir *adaptor.RajaOngkirAdaptor) (*DetailUsers, error) {
	detailAddress, err := getDetailAddress(user, rajaOngkir)
	if err != nil {
		return nil, err
	}

	return &DetailUsers{
		Id:       user.Id,
		Fullname: user.Fullname,
		Auth:     gin.H{"email": user.Email},
		Address:  detailAddress,
	}, nil
}

func getDetailAddress(user *model.User, rajaOngkir *adaptor.RajaOngkirAdaptor) (*Address, error) {
	if user.CityId != "" && user.CityId != "0" {
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

		detailAddress := &Address{
			Street: user.Street,
			City: AddressDetail{
				Id:   result["city_id"].(string),
				Name: result["city_name"].(string),
			},
			Province: AddressDetail{
				Id:   result["province_id"].(string),
				Name: result["province"].(string),
			},
		}
		return detailAddress, nil
	} else {
		detailAddress := &Address{
			Street: user.Street,
			City: AddressDetail{
				Id:   "",
				Name: "",
			},
			Province: AddressDetail{
				Id:   "",
				Name: "",
			},
		}
		return detailAddress, nil
	}
}

func SuccessAllUsersResponse(query *model.Pagination, payload interface{}) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: "GET_ALL_USERS_SUCCESS",
		Payload: payload,
		Query:   query,
	}
}

func SuccessUserResponse(msg string, statusCode int) *Response {
	return &Response{
		Status:      statusCode,
		Message:     msg,
		GeneralInfo: "Golang-4-Shop",
	}
}
