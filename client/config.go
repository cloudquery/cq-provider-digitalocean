package client

// Provider Configuration

type Config struct {
	Token string `hcl:"token"`

	SpacesRegions     []string `hcl:"spaces_regions,optional"`
	SpacesAccessKey   string   `hcl:"spaces_access_key,optional"`
	SpacesAccessKeyId string   `hcl:"spaces_access_key_id,optional"`
}

func (c Config) Example() string {
	return `
		configuration {
			// API Token to access DigialOcean resources 
			// See https://docs.digitalocean.com/reference/api/api-reference/#section/Authentication
			token = <YOUR_API_TOKEN_HERE>
			
			spacesRegions = []
			spacesKey = <YOUR_SPACES_ACCESS_KEYS>
			spacesToken = <YOUR_SPACES_ACCEsS_TOKEN>
		}
`
}
