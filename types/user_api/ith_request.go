package userapi

type Profile struct {
	Id                int     `json:"id"`
	FirstName         *string `json:"firstName"`
	LastName          *string `json:"lastName"`
	PreferredCurrency *string `json:"preferredCurrency"`
	Country           *string `json:"country"`
	Language          *string `json:"language"`
	BirthDate         *string `json:"birthDate"`
	TimeLimit         int64   `json:"timeLimit"`
	TransactionLimit  int64   `json:"transactionLimit"`
	Phone             *string `json:"phone"`
	Email             *string `json:"email"`
	EmailVerified     bool    `json:"emailVerified"`
	Password          *string `json:"password"`
	Created int64 `json:"created"`
	Updated int64 `json:"updated"`
	Addresses [] *UserAddress `json:"addresses"`
}

type UserAddress struct {
	Id                int    `json:"id"`
	UserId            int    `json:"userId"`
	CountryCode       string `json:"countryCode"`
	City              string `json:"city"`
	FirstAddressLine  string `json:"firstAddressLine"`
	SecondAddressLine string `json:"secondAddressLine"`
	State             string `json:"state"`
	PostalCode        string `json:"postalCode"`
}
