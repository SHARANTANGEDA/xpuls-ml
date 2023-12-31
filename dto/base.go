package dto

import (
	"context"
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BaseListOption struct {
	Start             *uint     `query:"start"`
	Count             *uint     `query:"count"`
	Search            *string   `query:"search"`
	Keywords          *[]string `query:"keywords"`
	KeywordFieldNames *[]string `query:"keyword_field_names"`
}

func (opt BaseListOption) BindQueryWithLimit(query *gorm.DB) *gorm.DB {
	if opt.Count != nil {
		query = query.Limit(int(*opt.Count))
	}
	if opt.Start != nil {
		query = query.Offset(int(*opt.Start))
	}
	return query
}

func (opt BaseListOption) BindQueryWithKeywords(query *gorm.DB, tableName string) *gorm.DB {
	tableName = query.Statement.Quote(tableName)
	keywordFieldNames := []string{"name"}
	if opt.KeywordFieldNames != nil {
		keywordFieldNames = *opt.KeywordFieldNames
	}
	if opt.Search != nil && *opt.Search != "" {
		sqlPieces := make([]string, 0, len(keywordFieldNames))
		args := make([]interface{}, 0, len(keywordFieldNames))
		for _, keywordFieldName := range keywordFieldNames {
			keywordFieldName = query.Statement.Quote(keywordFieldName)
			sqlPieces = append(sqlPieces, fmt.Sprintf("%s.%s LIKE ?", tableName, keywordFieldName))
			args = append(args, fmt.Sprintf("%%%s%%", *opt.Search))
		}
		query = query.Where(fmt.Sprintf("(%s)", strings.Join(sqlPieces, " OR ")), args...)
	}
	if opt.Keywords != nil {
		sqlPieces := make([]string, 0, len(keywordFieldNames))
		args := make([]interface{}, 0, len(keywordFieldNames))
		for _, keywordFieldName := range keywordFieldNames {
			keywordFieldName = query.Statement.Quote(keywordFieldName)
			sqlPieces_ := make([]string, 0, len(*opt.Keywords))
			for _, keyword := range *opt.Keywords {
				sqlPieces_ = append(sqlPieces_, fmt.Sprintf("%s.%s LIKE ?", tableName, keywordFieldName))
				args = append(args, fmt.Sprintf("%%%s%%", keyword))
			}
			sqlPieces = append(sqlPieces, fmt.Sprintf("(%s)", strings.Join(sqlPieces_, " AND ")))
		}
		query = query.Where(fmt.Sprintf("(%s)", strings.Join(sqlPieces, " OR ")), args...)
	}
	return query
}

type IDBService interface {
	getBaseDB(ctx context.Context) *gorm.DB
}

func getBaseQuery(ctx context.Context, service IDBService) *gorm.DB {
	return service.getBaseDB(ctx).Preload(clause.Associations)
}
