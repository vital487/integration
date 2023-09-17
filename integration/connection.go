package integration

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

const (
	DB_TYPE_MSSQL = "MSSQL"
	DB_TYPE_HANA  = "HANA"
)

type Connection struct {
	Id          *int
	Name        *string
	CompanyName *string
	B1iSld      *string
	B1iIp       *string
	B1iUsername *string
	B1iPassword *string
	DbType      *string
	DbIp        *string
	DbPort      *int
	DbName      *string
	DbUsername  *string
	DbPassword  *string
	Active      *int
}

func GetConnectionFromCompanyName(companyName string) (*Connection, error) {
	result := DB.QueryRow(`
		select *
		from connections
		where companyName = ?
	`, companyName)
	if result.Err() != nil {
		return nil, result.Err()
	}

	c := Connection{}
	err := result.Scan(&c.Id, &c.Name, &c.CompanyName, &c.B1iSld, &c.B1iIp, &c.B1iUsername, &c.B1iPassword, &c.DbType, &c.DbIp, &c.DbPort, &c.DbName, &c.DbUsername, &c.DbPassword, &c.Active)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (c Connection) GetDb() (*sql.DB, error) {
	switch *c.DbType {
	case DB_TYPE_MSSQL:
		fallthrough
	default:
		connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", *c.DbIp, *c.DbUsername, *c.DbPassword, *c.DbPort, *c.DbName)

		// Create connection pool
		db, err := sql.Open("sqlserver", connString)

		if err != nil {
			return nil, err
		}

		return db, nil
	}
}

// func (c Connection) Query(command string) sql.Rows {
// 	db, err := c.GetDb()

// 	db.Query()
// }
