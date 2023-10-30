package client

type ListClients struct {
	Clients []*Client
}

type Repository interface {
	Save(client *Client) error
	GetAll() (*ListClients, error)
	GetClientByDocumentNumber(docNumber string) (*Client, error)
	Update(client *Client) error
	Delete(id int) error
	CheckClientExists(docNumber, email string) (*Client, error)
}
