package staticstic

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/milara8888/collect_statistic/pkg/settings"
	"github.com/milara8888/collect_statistic/pkg/storage"
	"github.com/milara8888/collect_statistic/pkg/storage/statisticdb"

)


var wait time.Duration

type Statistic struct {

	hostAlowed map[string]bool

	Routes *gin.Engine

	Config *settings.Config

	DB storage.IStatisticDB

	ctx     context.Context

	errChan chan error

}

func New(c *settings.Config) (*Statistic, error) {

	hostAlowed := []string{}

	HOST_ALLOWED, exists := os.LookupEnv("HOST_ALLOWED")
	if exists {
		hostAlowed = strings.Split(HOST_ALLOWED, " ")
	}

	// moskow, _ := time.LoadLocation("Europe/Moscow")
	// crone := cron.New(cron.WithLocation(moskow))

	hosts := make(map[string]bool, len(hostAlowed))
	for _, h := range hostAlowed {
		hosts[h] = true
	}

	db, err := statisticdb.New(c)
	if err != nil {
		return nil, err
	}

	rest := &Statistic{
		Routes:         gin.Default(),
		Config:         c,
		hostAlowed:     hosts,
		DB:             db,
	}

	rest.initializeRoutes()
	return rest, err
}

// общий слушатель ошибок для потоков
func (ms *Statistic) errorListener(ctx context.Context) chan error {
	var (
		wg  sync.WaitGroup
		out = make(chan error)
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case err, ok := <-out:
				if !ok {
					return
				}
				if err != nil {
					log.Println(err)
				}
			}
		}
	}()

	go func() {
		defer close(out)
		wg.Wait()
	}()

	return out
}

func (s *Statistic) Start() error {

	connWs := net.JoinHostPort(s.Config.Host, s.Config.Port)
	log.Printf(`Server start : %s`, connWs)

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// s.ctx = ctx

	go s.errorListener(ctx)

	srv := &http.Server{
		Addr:           connWs,
		WriteTimeout:   time.Second * 15,
		ReadTimeout:    time.Second * 15,
		IdleTimeout:    time.Second * 60,
		MaxHeaderBytes: 1 << 20,
		Handler:        s.Routes,
		// BaseContext: func(l net.Listener) context.Context {
		// 	return s.ctx
		// },
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	// https://ru.wikipedia.org/wiki/%D0%A1%D0%B8%D0%B3%D0%BD%D0%B0%D0%BB_(Unix)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGTSTP)

	<-c

	srv.Shutdown(ctx)

	s.DB.Close(ctx)

	log.Println("shutting down")

	return nil
}
