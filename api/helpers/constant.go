package helpers

import "context"

var Ctx = context.Background()

const (
	// API_QUOTA string = "API_QUOTA"
	// DOMAIN           string = "DOMAIN"
	// DOMAIN_RETURN    string = "DOMAIN_RETURN"
	CANNOT_PARSE_ENV string = "cannot parse ENV keys"
)

const (
	HTTP_PREFIX string = "http"
	HTTP        string = "http://"
	HTTPS       string = "https://"
	WWW         string = "www."
)

const (
	COUNTER string = "counter"
	INDEX   string = "index:"

	CANNOT_CONNECT_TO_DB           string = "cannot connect to database"
	URL_NOT_FOUND                  string = "URL not found in database"
	CANNOT_RETRIEVE_TO_DB          string = "cannot retrieve to database"
	CANNOT_RETRIEVE_CLIENT         string = "cannot retrieve client"
	CANNOT_RETRIEVE_URLS_OF_CLIENT string = "cannot retrieve URLs of client"
	CANNOT_RETRIEVE_ALL_URLS       string = "cannot retrieve all URLs"
	INVALID_URL                    string = "invalid URL"
	CANNOT_SHORTEN_URL             string = "cannot shorten this URL"
	SHORT_LINK_EXISTS              string = "short link already exists"
)
