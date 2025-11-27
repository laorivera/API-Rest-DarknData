package services

//import "os"

func GetKey() string {
	key := "S@NN3N$AD!"
	if key == "" {
		// Fallback for development
		return "hashed-development-key-here"
	}
	return key
}
