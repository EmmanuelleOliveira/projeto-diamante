package client

type Repository interface {
	Save(client *Client) error
	GetAll() ([]*Client, error)
	GetClientByDocumentNumber(docNumber string) (*Client, error)
	Update(client *Client) error
	Delete(documentNumber string) error
	CheckClientExists(docNumber string) (*Client, error)
}
