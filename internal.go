package env

import "fmt"

func errorForKey(key string) error {
	return fmt.Errorf("failed to load key from environment: %s", key)
}
