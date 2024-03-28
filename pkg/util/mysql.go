package util

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLUtil struct {
	DB *sql.DB
}

// NewMySQLUtil 创建一个 MySQLUtil 实例
func NewMySQLUtil() (*MySQLUtil, error) {
	host := os.Getenv("OPSHELPER_RDS_HOST")
	database := os.Getenv("OPSHELPER_RDS_DBNAME")
	username := os.Getenv("OPSHELPER_RDS_USERNAME")
	password := os.Getenv("OPSHELPER_RDS_PASSWORD")
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		username,
		password,
		host,
		"3306",
		database,
	)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MySQL database")

	return &MySQLUtil{
		DB: db,
	}, nil
}

// Close 关闭 MySQL 数据库连接
func (mu *MySQLUtil) Close() {
	err := mu.DB.Close()
	if err != nil {
		log.Println("Error closing MySQL connection:", err)
	}
}

// QueryTable 执行查询并返回结果集
func (mu *MySQLUtil) QueryTable(tableName string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT cloud_type,account_name,ak,sk FROM %s", tableName)
	rows, err := mu.DB.Query(query)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}

	return rows, nil
}
