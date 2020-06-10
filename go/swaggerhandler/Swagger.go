package swaggerhandler

import "net/http"

// GetSwaggerFS Returns FS access to the APIs swagger documentation
func GetSwaggerFS() http.FileSystem {
	fs := http.FileSystem(http.Dir("www/swagger-ui"))
	return fs
}
