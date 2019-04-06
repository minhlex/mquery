package mquery

import (
	"fmt"
	"strings"
)

// TODO: binding map insert
type InsertQueryBuilder interface {
	Value(value map[string]interface{}) InsertQueryBuilder

	IToQuery
}

type insertQueryBuilder struct {
	mapColValue map[string]interface{}

	qb *queryBuilder
}

func newInsertBuilder(qb *queryBuilder) InsertQueryBuilder {
	return &insertQueryBuilder{
		qb: qb,
	}
}
func (iqb *insertQueryBuilder) Value(mapValue map[string]interface{}) InsertQueryBuilder {
	for k, v := range mapValue {
		if !iqb.qb.colValid(k) {
			panic("column not exist. Please check " + iqb.qb.tableName + " QueryBuilder")
		}
		if value, ok := iqb.mapColValue[k]; ok && value != nil {
			iqb.mapColValue[k] = v
		} else {
			panic("cant find column : " + k)
		}
	}
	return iqb
}
func (iqb *insertQueryBuilder) ToQuery() string {
	listCol := []string{}
	listValue := []string{}
	for k, v := range iqb.mapColValue {
		listCol = append(listCol, k)
		listValue = append(listValue, interfaceToString(v))
	}
	return fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s)", iqb.qb.tableName, strings.Join(listCol, ","), strings.Join(listValue, ","))
}
