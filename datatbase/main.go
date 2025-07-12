package main

import (
	"context"
	"errors"
	"fmt"

	changelog "github.com/jl-sky/grom/golangNotes/datatbase/changeLog"
	"github.com/jl-sky/grom/golangNotes/datatbase/config"
	"github.com/jl-sky/grom/golangNotes/datatbase/loger"
	"github.com/jl-sky/grom/golangNotes/datatbase/models"
	"github.com/jl-sky/grom/golangNotes/datatbase/mysql"
	log "github.com/sirupsen/logrus"
)

type ChangeLogAdmin interface {
	Changelog(ctx context.Context, req *models.ChangelogAdminReq) error
}

func ChangeLogByAdmin(ctx context.Context, req *models.ChangelogAdminReq) error {
	if req == nil {
		return fmt.Errorf("req is empty")
	}
	admin, err := GetChangelogAdmin(req.TableName)
	if err != nil {
		return err
	}
	if admin != nil {
		err := admin.Changelog(ctx, req)
		log.Debugf("changelogByAdmin error is %+v", err)
	}
	return fmt.Errorf("changelog admin is empty")
}

func GetChangelogAdmin(tableName string) (ChangeLogAdmin, error) {
	if tableName == "" {
		return nil, errors.New("table name is empty")
	}
	db, err := mysql.Conn()
	if err != nil {
		return nil, err
	}
	changelog.InitChangeLogSystem(db)
	switch tableName {
	case config.TOutput:
		outputImpl, err := changelog.NewOutput(db)
		return outputImpl, err
	default:
		return nil, nil
	}
}

func main() {
	loger.InitLogger()
	// changelogAdminTest()
	TestConcurrentUpdates()
}
