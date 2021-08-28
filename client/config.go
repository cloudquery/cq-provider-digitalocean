package client

// Provider Configuration

type Config struct {
	Token string `hcl:"token"`
}

func (c Config) Example() string {
	return `
		configuration {
			// API Token to access DigialOcean resources 
			// See https://docs.digitalocean.com/reference/api/api-reference/#section/Authentication
			token = <YOUR_API_TOKEN_HERE>
		}
`
}
