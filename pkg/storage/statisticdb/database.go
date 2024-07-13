package statisticdb

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/milara8888/collect_statistic/pkg/settings"
	"github.com/milara8888/collect_statistic/pkg/storage"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/godror/godror"
)


type StatisticDB struct {
	*gorm.DB
	Config *settings.Config
}

func New(c *settings.Config) (storage.IStatisticDB, error) {

	var (
		dbConn = c.Conn.Db
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 2 * time.Second, // Slow SQL threshold
			LogLevel:      logger.Info,     // Log level
			//   IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
			//   ParameterizedQueries:      true,           // Don't include params in the SQL log
			//   Colorful:                  false,          // Disable color
		},
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		dbConn.Host, dbConn.User, dbConn.Passw, dbConn.Sid, strconv.Itoa(int(dbConn.Port)))

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: newLogger,
	})


	if err != nil {
		return nil, err
	}

	return &StatisticDB{
		Config: c,
		DB:     db,
	}, err
}

func (s *StatisticDB) Migrate(ctx context.Context) error {
	for _, t := range Tables() {
		err := s.DB.AutoMigrate(t)
		if err != nil {
			return err
		}
	}
	return nil
}

//TODO это плохо
func (s *StatisticDB) Close(ctx context.Context) error {
	return nil
}

