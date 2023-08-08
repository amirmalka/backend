package example

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/amirmalka/backend"
)

type MyBackend struct{}

// Implement the BackendProvider interface.
func (m *MyBackend) GetBackendAPIVersion() string {
	return "v1.0.0"
}

func (m *MyBackend) GetBackendAPI() backend.BackendAPI {
	return backend.BackendAPI{
		Version: m.GetBackendAPIVersion(),
		Endpoints: []backend.Endpoint{
			{
				Name: "event-receiver-rest",
				URL:  "https://report.domain.com",
			},
			{
				Name: "event-receiver-websocket",
				URL:  "wss://report.domain.com",
			},
			// ... Add other endpoints as needed.
		},
	}
}

func main() {
	myBackend := &MyBackend{}

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		response := myBackend.GetBackendAPI()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
