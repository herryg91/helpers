package hUrls

import (
	"net/url"
)

//IsURL check if the string is url or not
func IsURL(urlString string) bool {
	parsedURL, err := url.ParseRequestURI(urlString)
	return (err == nil) &&
		(parsedURL.Scheme == "http" || parsedURL.Scheme == "https") &&
		(parsedURL.Host != "")
}

//SecureURL make url scheme into https
func SecureURL(urlString string) (secureUrl string) {
	if !IsURL(urlString) {
		return urlString
	}
	parsedURL, _ := url.ParseRequestURI(urlString)
	parsedURL.Scheme = "https"
	secureUrl = parsedURL.String()
	return
}
