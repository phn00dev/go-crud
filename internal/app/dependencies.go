package app

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/phn00dev/go-crud/pkg/config"
	dbconfig "github.com/phn00dev/go-crud/pkg/database/db_config"
	httpclient "github.com/phn00dev/go-crud/pkg/http_client"
)

type Dependencies struct {
	Config     *config.Config
	DB         *gorm.DB
	HttpClient *http.Client
}

func GetDependencies() (*Dependencies, error) {
	// get config
	getConfig, err := config.GetConfig()
	if err != nil {
		return nil, err
	}
	// db connection
	newDB := dbconfig.NewDbConnection(getConfig)
	getDB, err := newDB.GetDB()
	if err != nil {
		return nil, err
	}
	// HTTP client
	clientHttp := httpclient.NewHttpClient()

	return &Dependencies{
		Config:     getConfig,
		DB:         getDB,
		HttpClient: clientHttp,
	}, nil
}
