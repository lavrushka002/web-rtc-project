package client

type Client struct {
	ip       string
	port     string
	username string
	password string
}

func New(ip, port, username, password string) *Client {
	return &Client{
		ip:       ip,
		port:     port,
		username: username,
		password: password,
	}
}

func (c *Client) IP() string {
	return c.ip
}

func (c *Client) Port() string {
	return c.port
}

func (c *Client) Username() string {
	return c.username
}

func (c *Client) Password() string {
	return c.password
}
