package server

import (
	"github.com/go-yaml/yaml"
	"net/http"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		path := r.URL.Path

		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error){
	var pathUrls []PathUrl
	err := yaml.Unmarshal(yamlBytes, &pathUrls)
	if err != nil {
		return nil, err
	}

	pathsToUrls := map[string]string{}
	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.URL
	}

	return MapHandler(pathsToUrls, fallback), nil
}

type PathUrl struct {
	Path string `yaml:"path"`
	URL string `yaml:"url"`
}