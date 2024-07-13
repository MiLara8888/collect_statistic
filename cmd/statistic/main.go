package main

import (
	"context"
	"log"
	"os"
	"runtime"

	"github.com/urfave/cli"
	"github.com/milara8888/collect_statistic/pkg/settings"
	"github.com/milara8888/collect_statistic/pkg/storage/statisticdb"
	"github.com/milara8888/collect_statistic/internal/staticstic"
)

//использование настроек из .env
var (
	config *settings.Config
)


func init() {
	var err error
	runtime.GOMAXPROCS(runtime.NumCPU() - 2)
	config, err = settings.InitEnv()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := &cli.App{
		Name:  "Collect statictic rest server",
		Usage: "",
		Commands: []cli.Command{
			{
				Name:  "statistic",
				Usage: "statistic rest http",
				Action: func(*cli.Context) error {
					return StatisticRun()
				},
			},
			{
				Name:  "migrate",
				Usage: "Migrate db",
				Action: func(*cli.Context) error {
					return Migrate()
				},
			},
		},
	}
	app.Run(os.Args)
}


//запускает сервис
func StatisticRun() error {
	r, err := staticstic.New(config)
	if err != nil {
		return err
	}
	err = r.Start()
	if err != nil {
		return err
	}
	return nil
}


//запускает миграцию в бд
func Migrate() error {
	db, err := statisticdb.New(config)
	if err != nil {
		return err
	}
	err = db.Migrate(context.TODO())
	if err != nil {
		return err
	}
	return nil
}
