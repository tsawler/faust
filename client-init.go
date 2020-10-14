package clienthandlers

import (
	"github.com/tsawler/goblender/client/clienthandlers/clientdb"
	"github.com/tsawler/goblender/pkg/config"
	"github.com/tsawler/goblender/pkg/driver"
	"github.com/tsawler/goblender/pkg/handlers"
	"log"
)

var app config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger
var parentDB *driver.DB
var repo *handlers.DBRepo

var dbModel *clientdb.DBModel

// ClientInit gives us access to site values for client code.
func ClientInit(c config.AppConfig, p *driver.DB, rep *handlers.DBRepo) {
	// c is the application config, from goblender
	app = c
	repo = rep

	// loggers
	infoLog = app.InfoLog
	errorLog = app.ErrorLog

	// in case we need it, we get the db connection from goblender and save it in a variable
	parentDB = p

	// we can access handlers from goblender, but need to initialize them first
	if app.Database == "postgresql" {
		handlers.NewPostgresqlHandlers(parentDB, app.ServerName, app.InProduction, &app)
	} else {
		handlers.NewMysqlHandlers(parentDB, app.ServerName, app.InProduction, &app)
	}

	dbModel = &clientdb.DBModel{DB: p.SQL}

	// create client middleware
	NewClientMiddleware(app)
}
