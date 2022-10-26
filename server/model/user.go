package model

type User struct {
	BaseModel
	FullName   string `json:"full_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Gender     string `json:"gender"`
	Contact    string `json:"contact"`
	Street     string `json:"street"`
	CityId     int32  `json:"city_id"`
	ProvinceId int32  `json:"province_id"`
	Role       string `json:"role"`
}
