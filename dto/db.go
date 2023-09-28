package dto

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/xpuls-com/xpuls-ml/common/utils"
	"github.com/xpuls-com/xpuls-ml/config"
	"gorm.io/driver/postgres"
	"log"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	db         *gorm.DB
	dbOpenOnce sync.Once
)

type DbCtxKeyType string

const DbSessionKey DbCtxKeyType = "session"

type GormLogger struct{}

func (l *GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return l.getLogger(context.Background()).LogMode(level)
}

func (l *GormLogger) getLogger(ctx context.Context) gormlogger.Interface {
	prefix := "\r\n"

	return gormlogger.New(log.New(os.Stdout, prefix, log.LstdFlags), gormlogger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      gormlogger.Warn,
		Colorful:      true,
	})
}

func (l *GormLogger) Info(ctx context.Context, format string, args ...interface{}) {
	l.getLogger(ctx).Info(ctx, format, args...)
}

func (l *GormLogger) Warn(ctx context.Context, format string, args ...interface{}) {
	l.getLogger(ctx).Warn(ctx, format, args...)
}

func (l *GormLogger) Error(ctx context.Context, format string, args ...interface{}) {
	l.getLogger(ctx).Error(ctx, format, args...)
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	l.getLogger(ctx).Trace(ctx, begin, fc, err)
}

func getPgHost() string {
	return config.ServerConfig.Postgresql.Host
}

// nolint: unparam
func getDBURI() (string, error) {
	user := url.QueryEscape(config.ServerConfig.Postgresql.User)
	password := url.QueryEscape(config.ServerConfig.Postgresql.Password)
	database := url.QueryEscape(config.ServerConfig.Postgresql.Database)
	sslMode := "disable"
	if config.ServerConfig.Postgresql.SSLMode != "" {
		sslMode = config.ServerConfig.Postgresql.SSLMode
	}
	uri := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		user,
		password,
		getPgHost(),
		config.ServerConfig.Postgresql.Port,
		database,
		sslMode)
	return uri, nil
}

func openDB() (*gorm.DB, error) {
	uri, err := getDBURI()
	if err != nil {
		return nil, errors.Wrap(err, "cannot get db uri")
	}

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		PrepareStmt:    false,
	})
	if err != nil {
		return nil, errors.Wrap(err, "open db")
	}
	var rawDb *sql.DB
	rawDb, err = db.DB()
	if err != nil {
		return nil, err
	}
	rawDb.SetMaxOpenConns(config.ServerConfig.Postgresql.MaxOpenConns)
	rawDb.SetMaxIdleConns(config.ServerConfig.Postgresql.MaxIdleConns)
	rawDb.SetConnMaxLifetime(config.ServerConfig.Postgresql.ConnMaxLifetime)
	logrus.Infof("pg max open connections: %d", config.ServerConfig.Postgresql.MaxOpenConns)
	logrus.Infof("pg max idle connections: %d", config.ServerConfig.Postgresql.MaxIdleConns)
	logrus.Infof("pg connection max lifetime: %s", config.ServerConfig.Postgresql.ConnMaxLifetime.String())
	return db, nil
}

func getDB() (*gorm.DB, error) {
	var err error
	dbOpenOnce.Do(func() {
		db, err = openDB()
	})
	if err != nil {
		return nil, err
	}
	if config.GlobalCommandOption.Debug {
		return db.Debug(), nil
	}

	return db, nil
}

func StartTransaction(ctx *gin.Context) (*gorm.DB, context.Context, func(error), error) {
	return startTransaction(ctx)
}

type TransactionDBWrapper struct {
	orig     *gorm.DB
	released bool
}

// nolint: unparam
func startTransaction(ctx context.Context) (*gorm.DB, context.Context, func(error), error) {
	// FIXME: pq: unexpected Parse response 'D'
	defaultCb := func(err error) {}
	// return mustGetDB(), ctx, defaultCb, nil
	session_ := ctx.Value(DbSessionKey)
	if session_ != nil {
		db_ := session_.(*TransactionDBWrapper)
		if !db_.released {
			return db_.orig, ctx, defaultCb, nil
		}
	}
	db := mustGetDB(ctx)
	tx := db.Begin()
	if tx.Error != nil {
		return nil, ctx, defaultCb, tx.Error
	}
	db_ := &TransactionDBWrapper{orig: tx}
	ctx = context.WithValue(ctx, DbSessionKey, db_)
	return tx, ctx, func(err error) {
		select {
		case <-ctx.Done():
			return
		default:
		}
		if db_.released {
			return
		}
		db_.released = true
		// nolint: gocritic
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}, nil
}

func mustGetSession(ctx context.Context) *gorm.DB {
	session_ := ctx.Value(DbSessionKey)
	if session_ != nil {
		db_ := session_.(*TransactionDBWrapper)
		if !db_.released {
			return db_.orig
		}
	}
	return mustGetDB(ctx)
}

func mustGetDB(ctx context.Context) *gorm.DB {
	db, err := getDB()
	if err != nil {
		panic(fmt.Sprintf("cannot get db: %s", err.Error()))
	}
	db = db.WithContext(ctx)
	return db
}

type MigrateLog struct{}

func (*MigrateLog) Printf(format string, v ...interface{}) {
	logrus.Infof(fmt.Sprintf("[%s] %s", time.Now(), format), v...)
}

func (*MigrateLog) Verbose() bool {
	return false
}

func MigrateUp() error {
	uri, err := getDBURI()
	if err != nil {
		return errors.Wrap(err, "cannot get db uri")
	}

	logrus.Debugf("db uri: %s", uri)
	migrationDir := config.ServerConfig.MigrationDir

	exists, err := utils.PathExists(migrationDir)
	if err != nil {
		return errors.Wrapf(err, "check migration dir exists: %s", migrationDir)
	}
	if !exists {
		return errors.Errorf("migration dir is not exists: %s", migrationDir)
	}

	logrus.Debugf("migration dir: %s", migrationDir)
	m, err := migrate.New(
		fmt.Sprintf("file://%s", migrationDir),
		uri,
	)
	if err != nil {
		return errors.Wrap(err, "cannot create migrate")
	}
	defer m.Close()

	m.Log = &MigrateLog{}

	logrus.Info("migrate up...")
	if err := m.Up(); err != nil && !strings.Contains(err.Error(), "no change") {
		return errors.Wrap(err, "cannot migrate up")
	}
	logrus.Info("[DONE] migrate up")
	return nil
}
