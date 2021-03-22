package env

import "os"

// TryGet attempts to get environment variable key and returns error in case of
// failure.
func TryGet(key string) (val string, err error) {
	if val = os.Getenv(key); val == "" {
		err = errorForKey(key)
	}
	return
}

// MustGet attemts to get environment variable key and panics on failure.
func MustGet(key string) (val string) {
	val, err := TryGet(key)
	if err != nil {
		panic(err)
	}
	return
}
