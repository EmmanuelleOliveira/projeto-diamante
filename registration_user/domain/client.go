package domain

type Client struct {
	Id             int
	Name           string
	Email          string
	DocumentNumber string
	PhoneNumber    string
	Address        AddressClient
}

type AddressClient struct {
	Cep    string
	Street string
	City   string
	UF     string
}
