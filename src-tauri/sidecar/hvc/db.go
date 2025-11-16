package hvc

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/microsoft/go-mssqldb"
)

var db *sqlx.DB

func InitDB(dev bool) (err error) {
	host := "localhost"
	if dev {
		host = "10.50.61.135"
	}
	dsn := fmt.Sprintf("sqlserver://sa:123456@%s/?timezone=Asia/Shanghai&database=Fusion", host)
	db, err = sqlx.Connect("sqlserver", dsn)
	if err != nil {
		return err
	}
	return nil
}

func CreateDevice() error {
	if db == nil {
		return fmt.Errorf("db not initialized")
	}
	query := `CREATE TABLE IF NOT EXISTS Device (
				id INT PRIMARY KEY,
				name VARCHAR(255) NOT NULL,
				skip_length FLOAT NOT NULL,
			)`

	_, err := db.Exec(query)
	return err
}

func CreateChangeInfo() error {
	if db == nil {
		return fmt.Errorf("db not initialized")
	}
	query := `CREATE TABLE IF NOT EXISTS ChangeInfo (
				id INT PRIMARY KEY,
				device_name VARCHAR(255) NOT NULL,
				device_time DATETIME NOT NULL,
				length FLOAT NOT NULL,
				voltage FLOAT NOT NULL,
				created_at DATETIME NOT NULL
			}
			)`
	_, err := db.Exec(query)
	return err

}

func CreateDeviceInfo() error {
	if db == nil {
		return fmt.Errorf("db not initialized")
	}
	query := `CREATE TABLE IF NOT EXISTS DeviceInfo (
				id INT PRIMARY KEY,
				device_name VARCHAR(255) NOT NULL,
				status BOOLEAN NOT NULL,
				created_at DATETIME NOT NULL
			)`
	_, err := db.Exec(query)
	return err
}

func CreateTable() error {
	if db == nil {
		return fmt.Errorf("db not initialized")
	}
	err := CreateDevice()
	if err != nil {
		return err
	}
	err = CreateChangeInfo()
	if err != nil {
		return err
	}
	err = CreateDeviceInfo()
	if err != nil {
		return err
	}
	return nil

}
