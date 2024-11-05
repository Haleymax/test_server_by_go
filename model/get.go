package model

import "fmt"

// OrderList 数据库字段
type OrderList struct {
	Id       int    `json:"id"`
	IdName   string `json:"id_name"`
	IdNumber string `json:"id_number"`
}

func (o *OrderList) List() (err error, ols []OrderList) {
	list := []OrderList{}
	row, err := MysqlDb.Query("SELECT * FROM id_card_information ORDER BY RAND() LIMIT 1")
	if err != nil {
		fmt.Println(err)
		return err, nil
	}
	fmt.Println(row)
	for row.Next() {
		item := OrderList{}
		err = row.Scan(&item.Id, &item.IdName, &item.IdNumber)
		if err != nil {
			fmt.Println(err)
			return err, nil
		}
		list = append(list, item)
	}
	return nil, list
}
