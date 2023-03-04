package database

import (
	"fmt"
	"strings"
)

func wrapValue[T any](val T) string {
	return `'` + fmt.Sprintf("%v", val) + `'`
}

func wrapWithQuotationMark[T any](val T) string {
	// val.(type)
	ret := fmt.Sprintf("%v", val)
	if fmt.Sprintf("%T", val) == "string" {
		ret = `'` + ret + `'`
	}
	return ret
}

// func KeysOfMap(m interface{}) (keys []interface{}) {
// 	v := reflect.ValueOf(m)
// 	if v.Kind() != reflect.Map {
// 		fmt.Errorf("input type not a map: %v", v)
// 	}

// 	for _, k := range v.MapKeys() {
// 		keys = append(keys, k.Interface())
// 	}
// 	return keys
// }

// TODO: char(string) 타입이 아닌 경우, `'`로 감싸는 것은 에러를 유발하는 경우가 있는지 확인
// 에러가 없다면 무조건 `'`로 감싸면 됨
func WrappedValues(values []interface{}) (vals []string) {
	for _, value := range values {
		vals = append(vals, wrapValue(value))
	}
	return vals
}

func ConvertValues2(values []interface{}) (vals []string) {
	for _, value := range values {
		vals = append(vals, wrapWithQuotationMark(value))
	}
	return vals
}

func KeysValuesOfMap(data map[string]interface{}) (keys []string, vals []string) {
	for key, val := range data {
		keys = append(keys, key)
		vals = append(vals, wrapWithQuotationMark(val))
	}
	return keys, vals
}

func ValuesStringByKeys(data map[string]interface{}, keys []string) string {
	vals := []string{}
	for _, key := range keys {
		vals = append(vals, wrapWithQuotationMark(data[key]))
	}
	return "(" + strings.Join(vals, ", ") + ")"
}

func SetSqlSelect(table string, fields []string, added ...string) string {
	keys := "*"
	if len(fields) > 0 {
		keys = strings.Join(fields, ", ")
	}
	return "SELECT " + keys + " FROM " + table + " " + strings.Join(added, " ") + ";"
}

func SetSqlInsertOne(table string, keys []string, vals []interface{}) string {
	return "INSERT INTO " + table + " (" + strings.Join(keys, ", ") + ")" + " VALUES (" + strings.Join(WrappedValues(vals), ", ") + ");"
}

func SetSqlInsert(table string, keys []string, data [][]interface{}) string {
	valuesStr := ""
	for _, vals := range data {
		valuesStr += "(" + strings.Join(WrappedValues(vals), ", ") + "),\n"
	}
	valuesStr = valuesStr[:len(valuesStr)-2]
	return "INSERT INTO " + table + " (" + strings.Join(keys, ", ") + ")" + " VALUES " + valuesStr + ";"
}

// func SetSqlInsertOneDict(table string, data map[string]interface{}) string {
// 	keys, vals := KeysValuesOfMap(data)
// 	str := "INSERT INTO " + table + " (" + strings.Join(keys, ", ") + ")" + " VALUES (" + strings.Join(vals, ", ") + ");"
// 	println(str)
// 	return "INSERT INTO " + table + " (" + strings.Join(keys, ", ") + ")" + " VALUES (" + strings.Join(vals, ", ") + ");"
// }

// func SetSqlInsertDict(table string, data []map[string]interface{}) string {
// 	keys, _ := KeysValuesOfMap(data[0])
// 	vals := []string{}
// 	for _, d := range data {
// 		vals = append(vals, ValuesStringByKeys(d, keys))
// 	}
// 	return "INSERT INTO " + table + " (" + strings.Join(keys, ", ") + ")" + " VALUES " + strings.Join(vals, ", ") + ";"
// }

func SetSqlUpdate(table string, data map[string]interface{}) string {
	keys, vals := KeysValuesOfMap(data)
	query := "UPDATE " + table + " SET "
	for i := range keys {
		query += keys[i] + " = " + vals[i] + ", "
	}
	return query[:len(query)-2]
}

func SetSqlUpsert(table string, data map[string]interface{}) string {
	// keys, vals := KeysValuesOfMap(data)
	// query := "UPDATE " + table + " SET "
	// for i := range keys {
	// 	query += keys[i] + " = " + vals[i] + ", "
	// }
	return ""
}

// func SetSqlCreateTable(table string, schema string) string {
// 	return "CREATE TABLE " + table + " " + schema
// }

// // func StrInsertKeyValues(data map[string]interface{}) string {
// // 	keys, vals := KeysValuesOfMap(data)
// // 	return "(" + strings.Join(keys, ", ") + ")"  strings.Join(vals, ", ")
// // }

// // func setInsertKeyValsFromMaps(data []map[string]interface{}) (string, string) {
// // 	keys := keySlice(data[0])
// // 	vals := []string{}
// // 	for _, d := range data {
// // 		vals = append(vals, valSql(d, keys))
// // 	}

// // 	return "(" + strings.Join(keys, ", ") + ")", strings.Join(vals, ", ")
// // }
