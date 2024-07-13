package test

import (
	"log"
	"runtime"
	"testing"

	"github.com/milara8888/collect_statistic/pkg/settings"
)

var (
	err error
	// настройка подключения
	config *settings.Config
)


func TestMain(m *testing.M) {

	config, err = settings.InitEnv()
	if err != nil {
		log.Fatal(err)
	}
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	m.Run()
}
