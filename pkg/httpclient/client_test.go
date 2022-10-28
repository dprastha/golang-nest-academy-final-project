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
	Results interface{}
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
	resp, err := client.Get("/city?id=12")
	assert.Nil(t, err)

	var datas JSONBaseGet

	err = json.Unmarshal(resp, &datas)
	assert.Nil(t, err)
	fmt.Println(datas.Rajaongkir)
}
