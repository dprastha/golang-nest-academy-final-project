package httpclient

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type JSONBaseGet struct {
	Rajaongkir JSONGet
}

type JSONGet struct {
	Query   interface{}
	Status  interface{}
	Results JSONResultsGet
}

type JSONResultsGet struct {
	City_id     string
	Province_id string
	Province    string
	Kabupaten   string
	City_name   string
	Postal_code string
}

type JSONBasePost struct {
	Rajaongkir JSONPost
}

type JSONPost struct {
	Query               interface{}
	Status              interface{}
	Origin_details      interface{}
	Destination_details interface{}
	Results             interface{}
}

func TestGet(t *testing.T) {
	client := NewHttpClient("https://api.rajaongkir.com/starter", "6d4d26125ea1c2991b801880cf3842f7")
	resp, err := client.Get("city?id=39")
	assert.Nil(t, err)

	var datas JSONBaseGet

	err = json.Unmarshal(resp, &datas)
	assert.Nil(t, err)
	fmt.Println(datas.Rajaongkir.Results)
}

func TestPost(t *testing.T) {
	client := NewHttpClient("https://api.rajaongkir.com/starter", "6d4d26125ea1c2991b801880cf3842f7")
	resp, err := client.Post("cost", map[string]interface{}{
		"origin": "10",
		"destination":  "15",
		"weight":   1700,
		"courier": "jne",
	})
	assert.Nil(t, err)

	var datas JSONBasePost

	err = json.Unmarshal(resp, &datas)
	assert.Nil(t, err)

	fmt.Println(datas.Rajaongkir)
}
