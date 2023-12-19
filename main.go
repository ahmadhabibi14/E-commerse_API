package main

import (
	"e-commerse_api/app"
	"e-commerse_api/conf"
	"e-commerse_api/models"
	"os"
	"strings"
)

func main() {
	conf.LoadEnv()
	zlog := conf.InitLogger()
	validArgs := `web, migrate`

	var mode string
	if len(os.Args) < 2 {
		mode = `web`
	} else {
		mode = strings.ToLower(os.Args[1])
	}

	switch mode {
	case `web`:
		ws := &app.WebServer{
			AppName: "E-commerse Store",
			Cfg:     conf.EnvWebConf(),
		}
		ws.Start()
	case `migrate`:
		models.RunMigration()
	default:
		zlog.Fatal().Msg(`Must start with: ` + validArgs)
	}
}
