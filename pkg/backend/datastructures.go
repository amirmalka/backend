package backend

// Endpoint represents a single backend endpoint with its name and URL.
type Endpoint struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// BackendAPI represents the structure of the backend API response.
type BackendAPI struct {
	Version   string     `json:"version"`
	Endpoints []Endpoint `json:"endpoints"`
}
