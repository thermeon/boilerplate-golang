package utils

// ConfigReader interface for configuration reading.
type ConfigReader interface {
	// Read the configurations and i should be a pointer.
	Read(i interface{}) (err error)
}
