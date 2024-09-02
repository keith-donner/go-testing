package core

import "errors"

// Select (CTRL+Space to change) From “table_name” and return these values [“values] where [“conditions_to_query] and save in

var things []map[string]any

// variableList: contains the names of the columns to retrieve
//
// queryString: SQL query statement
func Find(tableName string, variableList []map[string]any, queryString string) []map[string]any {

	return things

}
func Info(tableName string, variableList []map[string]any, queryString string) []map[string]any {

	return things

}

// variablePairs: map of key value pairs. key is the column name
//
// returns the error response if one thrown
func Insert(tableName string, variablePairs []map[string]any) string {

	return "texring"

}

func InsertNew(tableName string, variablePairs []map[string]any) error {

	return errors.New("error")

}
