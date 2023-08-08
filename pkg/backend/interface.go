package backend

// BackendProvider is the interface that any backend implementation must satisfy.
type BackendProvider interface {
	GetBackendAPIVersion() string
	GetBackendAPI() BackendAPI
}
