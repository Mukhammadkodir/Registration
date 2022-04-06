package models


type User struct {
	Id    string `json:"id"`
	Name    string `json:"name"`
	Password      string `json:"password"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`

}

type Login struct {
	Name    string `json:"name"`
	Password string `json:"password"`
}

type ById struct {
	Id    string `json:"id"`
}



