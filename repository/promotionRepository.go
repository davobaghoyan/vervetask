package repository

import (
	"awesomeProject/model"
	"context"
	"github.com/go-sql-driver/mysql"
	"log"
	"time"
)

const bulkInsertQuery = `LOAD DATA LOCAL INFILE 'promotions.csv' INTO TABLE promotion
		    FIELDS TERMINATED BY ","
		    ENCLOSED BY '"'
		    (id, price, expiration_date)`
const getByIdQuery = "select * from promotion where id = ?"
const fileName = "promotions.csv"

func BulkInsertFromCsv() {
	db := GetDb()
	mysql.RegisterLocalFile(fileName)
	_, err := db.ExecContext(context.Background(), bulkInsertQuery)

	if err != nil {
		log.Fatal(err)
	}
}

func GetPromotionById(promotionId string) model.Promotion {
	db := GetDb()
	rows, err := db.Query(getByIdQuery, promotionId)
	if err != nil {
		log.Fatal(err)
	}
	var (
		id              string
		price           float64
		expiration_date time.Time
	)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &price, &expiration_date)
		if err != nil {
			log.Fatal(err)
		}
	}

	return model.Promotion{
		Id: id, Price: price, Expiration_date: expiration_date,
	}
}
