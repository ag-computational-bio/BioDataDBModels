package biodatadbapiswagger

import "net/http"

func GetSwaggerFS() http.FileSystem {
	fs := http.FileSystem(http.Dir("www/swagger-ui"))
	return fs
}
