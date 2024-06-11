package output

import (
	"errors"
	"os"
)

// ReadBannerFile reads the content of a banner file and returns it as a string.
func ReadBannerFile(filepath string) (string, error) {
	// Read file content
	bannerFile, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	// If file is empty report an error
	if len(bannerFile) == 0 {
		return "", errors.New("banner file is either empty or corrupted")
	}
	return string(bannerFile), nil
}
