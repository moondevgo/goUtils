// https://blog.logrocket.com/building-simple-app-go-postgresql/

package database

import (
	"log"

	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	basic "github.com/moondevgo/goUtils/guBasic"
)

// * Stock API 접속 정보
func GetApiSetting(serverName string) map[string]interface{} {
	path := basic.SetFilePath(`C:\MoonDev\_config`, "database", "database_conn.yaml")
	return basic.GetConfigMap(path, serverName)
}

const (
	YAML_NAME = "database_conn"
)

// DB접속정보를 가지고 있는 객체를 정의합니다.
type DBMS struct {
	ServerName string  `yaml:"server_name"`
	DbName     string  `yaml:"db_name"`
	User       string  `yaml:"user"`
	Passwd     string  `yaml:"password"`
	Host       string  `yaml:"host"`
	Port       int     `yaml:"port"` // TODO: string -> int
	Conn       *sql.DB `yaml:"conn"`
	// db      string // TODO: string -> int
	// Charset string `yaml:"charset"`
}

// Section: MySQL
func InitMysql(serverName, dbName string) *DBMS {
	config := GetApiSetting(serverName) // ex) "mysql_HLS_local"
	// log.Printf("\nconfig: %v, dbName: %v\n", config, dbName)

	format := "%s:%s@tcp(%s:%d)/%s" //
	dsn := fmt.Sprintf(format, config["user"], config["password"], config["host"], config["port"], dbName)
	conn, err := sql.Open("mysql", dsn)
	dbms := &DBMS{ServerName: serverName, DbName: dbName, User: config["user"].(string), Passwd: config["password"].(string), Host: config["host"].(string), Port: config["port"].(int), Conn: conn}
	if err != nil {
		log.Fatal(err)
	}

	return dbms
}

// DBMS method
func (dbms *DBMS) SQLExecQuery(query string) ([]map[string]interface{}, error) {
	// defer dbms.Conn.Close()

	//db를 통해 sql문을 실행 시킨다.
	rows, err := dbms.Conn.Query(query)
	// 함수가 종료되면 rows도 Close한다.
	defer rows.Close()

	//컬럼을 받아온다.
	cols, err := rows.Columns()
	//err발생했는지 확인한다.
	if err != nil {
		return nil, err
	}

	data := make([]interface{}, len(cols))

	for i, _ := range data {
		var d []byte
		data[i] = &d
	}

	results := make([]map[string]interface{}, 0)

	for rows.Next() {
		err := rows.Scan(data...)
		if err != nil {
			return nil, err
		}
		result := make(map[string]interface{})
		for i, item := range data {
			result[cols[i]] = string(*(item.(*[]byte)))
		}
		results = append(results, result)
	}

	return results, nil
}

// SQL을 실행합니다.
// func (dbms DBMS) SQLExec(query string, db *sql.DB) (int64, int64, error) {
func (dbms *DBMS) SQLExec(query string) (int64, int64, error) {
	// defer dbms.Conn.Close()

	//SQL을 실행합니다.
	result, err := dbms.Conn.Exec(query)
	if err != nil {
		fmt.Println(result, err)
		return -1, -1, err
	}

	// 변경된 row의 갯수를 가져옵니다.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return -1, -1, err
	}

	// 변경된 id를 가져옵니다.
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return -1, -1, err
	}

	return rowsAffected, lastInsertId, nil
}

func (dbms *DBMS) Find(table string, fields []string, added ...string) []map[string]interface{} {
	results, _ := dbms.SQLExecQuery(SetSqlSelect(table, fields, added...))
	return results
}

// func (dbms *DBMS) InsertOneDict(data map[string]interface{}, table string) {
// 	dbms.SQLExec(SetSqlInsertOneDict(data, table))
// }

// func (dbms *DBMS) InsertDict(data []map[string]interface{}, table string) {
// 	sql := SetSqlInsert(data, table)
// 	log.Printf("\n****sql: %v\n", sql)
// 	dbms.SQLExec(SetSqlInsertDict(data, table))
// }

func (dbms *DBMS) InsertOne(table string, fields []string, values []interface{}) {
	dbms.SQLExec(SetSqlInsertOne(table, fields, values))
}

func (dbms *DBMS) Insert(table string, fields []string, data [][]interface{}) {
	dbms.SQLExec(SetSqlInsert(table, fields, data))
}

// func Update(data map[string]interface{}, table string, db *sql.DB, dbms *DBMS) {
// 	dbms.SQLExec(SetSqlUpdate(data, table), db)
// }

// func UpsertOne(data map[string]interface{}, table string, db *sql.DB, dbms *DBMS) {
// 	dbms.SQLExec(SetSqlUpsertOne(data, table), db)
// }

// func Upsert(data []map[string]interface{}, table string, db *sql.DB, dbms *DBMS) {
// 	dbms.SQLExec(SetSqlUpsert(data, table), db)
// }

func (dbms *DBMS) CreateTable(table string, schema string) {
	dbms.SQLExec("CREATE TABLE " + table + " (" + schema + ") ENGINE=MYISAM CHARSET=utf8;")
}

func (dbms *DBMS) AddColumn(table string, schema string) {
	// ALTER TABLE krx_items ADD COLUMN warning varchar(8);
	dbms.SQLExec("ALTER TABLE " + table + " ADD COLUMN " + schema)
}

// func (dbms *DBMS) FindOne() {
// 	var name string
// 	err := dbms.Conn.QueryRow("SELECT mb_id FROM member_table WHERE mb_id='monwater'").Scan(&name)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(name)
// }
