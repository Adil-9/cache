package structures

type User struct {
	Id       int           `json:"id"`
	Name     string        `json:"name"`
	Username string        `json:"username"`
	Email    string        `json:"email"`
	Address  AddressStruct `json:"address"`
	Phone    string        `json:"phone"`
	Website  string        `json:"website"`
	Company  CompanyStruct `json:"company"`
}

type CompanyStruct struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

type AddressStruct struct {
	Street  string    `json:"street"`
	Suite   string    `json:"suite"`
	City    string    `json:"city"`
	Zipcode string    `json:"zipcode"`
	Geo     GeoStruct `json:"geo"`
}

type GeoStruct struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type RedisOptions struct {
	Address  string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

//user
// {
//     "id": 1,
//     "name": "Leanne Graham",
//     "username": "Bret",
//     "email": "Sincere@april.biz",
//     "address": {
//       "street": "Kulas Light",
//       "suite": "Apt. 556",
//       "city": "Gwenborough",
//       "zipcode": "92998-3874",
//       "geo": {
//         "lat": "-37.3159",
//         "lng": "81.1496"
//       }
//     },
//     "phone": "1-770-736-8031 x56442",
//     "website": "hildegard.org",
//     "company": {
//       "name": "Romaguera-Crona",
//       "catchPhrase": "Multi-layered client-server neural-net",
//       "bs": "harness real-time e-markets"
//     }
//   },
