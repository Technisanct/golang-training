package clients

type client struct {
}

var (
	c = &client{}
)

// New ... creates all clients
func New() {
}

// Get ... all clients
func Get() *client {
	return c
}
