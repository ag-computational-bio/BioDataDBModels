package swaggerhandler

import (
	"log"
	"net/http"
	"os"
)

// GetSwaggerFS Returns FS access to the APIs swagger documentation
func GetSwaggerFS() *http.FileSystem {
	log.Println(os.Getwd())
	fs := http.FileSystem(http.Dir("www/swagger-ui"))
	return &fs
}
