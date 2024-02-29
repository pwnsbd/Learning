package urlshort

import "net/http"

func mapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return nil
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.Handler, error) {
	return nil, nil
}
