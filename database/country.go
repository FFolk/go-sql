package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var dsn = "landmark:landmark@csmsu@tcp(202.28.34.197)/landmark"

type Country struct {
	idx  int
	name string
}

type Countryx struct {
	Idx  int
	Name string
}

func (c *Country) SetIdx(id int) {
	c.idx = id
}

func (c *Country) SetName(name string) {
	c.name = name
}

func AddCountry(country *Country) (int64, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return -1, err
	}
	sql := "insert into country (name) values (?)"
	result, err := db.Exec(sql, country.name)
	if err != nil {
		return -1, err
	}
	affeted, err := result.RowsAffected()
	if affeted >= 0 {
		return affeted, nil
	} else {
		return -1, err
	}
}

func UpdateCountry(Country *Country) (int64, error) {
	db, err := getConnection()
	if err != nil {
		return -1, err
	}
	sql := "update country set name = ? where idx = ?"
	result, err := db.Exec(sql, Country.name, Country.idx)
	if err != nil {
		return -1, nil
	}
	affected, err := result.RowsAffected()
	if affected >= 0 {
		return affected, nil
	} else {
		return -1, err
	}
}

func DeleteCountry(idx int) (int64, error) {
	db, err := getConnection()
	if err != nil {
		return -1, err
	}
	sql := "delete from country where idx = ?"
	result, err := db.Exec(sql, idx)
	if err != nil {
		return +1, err
	}
	affected, err := result.RowsAffected()
	if affected >= 0 {
		return affected, nil
	} else {
		return -1, err
	}
}

func getConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	println("connection success")

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	return db, nil
}
func GetCountry() ([]Country, error) {
	db, err := getConnection()
	if err != nil {
		return nil, err
	}

	query := "select * from country"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	cuntries := []Country{}
	for rows.Next() {
		contry := Country{}
		rows.Scan(&contry.idx, &contry.name)
		cuntries = append(cuntries, contry)
	}
	defer rows.Close()
	return cuntries, nil
}

func GetCountryByID(idx int) (*Country, error) {
	db, err := getConnection()
	if err != nil {
		return nil, err
	}

	query := "select * from country where idx = ?"
	rows := db.QueryRow(query, idx)
	contry := Country{}
	err = rows.Scan(&contry.idx, &contry.name)
	if err != nil {
		return nil, err
	}
	return &contry, err

}

func getConnectionx() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	println("connection success")

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	return db, nil
}
func GetCountryx() ([]Countryx, error) {
	db, err := getConnectionx()
	if err != nil {
		return nil, err
	}

	cuntries := []Countryx{}
	query := "select * from country"
	err = db.Select(&cuntries, query)
	if err != nil {
		return nil, err
	} else {
		return cuntries, nil
	}
}
