package database

import (
	"database/sql"
	"log"
)

type Customer struct {
	CustomerId   int    `json: "customer_id"`
	FirstName    string `json: "first_name"`
	LastName     string `json: "last_name"`
	NpwpId       string `json;"npwp_id"`
	Age          int    `json: "age"`
	CustomerType string `json: "customer_type"`
	Street       string `json: "street"`
	City         string `json: "city"`
	State        string `json: "state"`
	ZipCode      string `json: "zip_code"`
	PhoneNumber  string `json: "phone_number"`
}

//create new customer
func InsertCustomer(customer Customer, db *sql.DB) {
	_, err := db.Exec("insert into customers (first_name, last_name,npwp_id,age,customer_type,street,city,state,zip_code,phone_number) values (?,?,?,?,?,?,?,?,?,?)",
		customer.FirstName,
		customer.LastName,
		customer.NpwpId,
		customer.Age,
		customer.CustomerType,
		customer.Street,
		customer.City,
		customer.State,
		customer.ZipCode,
		customer.PhoneNumber)

	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("insert data success")
}

//get customer data
func GetCustomers(db *sql.DB) {
	rows, err := db.Query("select * from customers")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []Customer

	//Iterasi data per baris yang didapat dari query select
	for rows.Next() {
		var each = Customer{}
		//scan per kolom field
		var err = rows.Scan(
			&each.CustomerId,
			&each.FirstName,
			&each.LastName,
			&each.NpwpId,
			&each.Age,
			&each.CustomerType,
			&each.Street,
			&each.City,
			&each.State,
			&each.ZipCode,
			&each.PhoneNumber)
		if err != nil {
			log.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	log.Println(result)
}

func DeleteCustomer(id int, db *sql.DB) {
	_, err := db.Exec("delete from customers where customer_id = ?", id)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("success delete data")
}

func UpdateCustomer(age int, id int, db *sql.DB) {
	_, err := db.Exec("update customers set age = ? where customer_id = ?", age, id)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("success update data")
}
