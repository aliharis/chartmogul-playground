package chartmogul

import (
	"os"

	cm "github.com/chartmogul/chartmogul-go/v3"
)

var api cm.API

func init() {
	api = cm.API{
		ApiKey: os.Getenv("CHARTMOGUL_API_KEY"),
	}
}

func GetAPI() cm.API {
	return api
}
