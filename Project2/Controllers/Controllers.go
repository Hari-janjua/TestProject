package Controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"Project2/Models"
	"Project2/db"
)

func GetPersonInfo(c *gin.Context) {
	fmt.Println("IN: GetPersonInfo")
	personId := c.Param("person_id")

	db, dberr := db.GetSQLConnection()
	if dberr != nil {
		fmt.Println("ERROR in DB ", dberr)
		// logginghelper.LogError("ERROR in DB ", dberr)
		return
	}

	// SQL Query for fetching data
	// SELECT Person.Name, Phone.number, Address.city, Address.state, Address.street1, Address.street2, Address.zip_code
	// FROM Person
	// LEFT JOIN Phone ON Person.id=Phone.person_id
	// LEFT JOIN Address_join ON Person.id=Address_join.person_id
	// LEFT JOIN Address ON Address_join.address_id=Address.id
	// where Person.id=?
	var person Models.Person
	row := db.QueryRow("SELECT Person.Name, Phone.number, Address.city, Address.state, Address.street1, Address.street2, Address.zip_code FROM Person LEFT JOIN Phone ON Person.id=Phone.person_id LEFT JOIN Address_join ON Person.id=Address_join.person_id LEFT JOIN Address ON Address_join.address_id=Address.id WHERE Person.id=?", personId)
	err := row.Scan(&person.Name, &person.PhoneNumber, &person.City, &person.State, &person.Street1, &person.Street2, &person.ZipCode)
	if err != nil {
		fmt.Println("Error while executing the query")
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": person})
	fmt.Println("OUT: GetPersonInfo")
}

func CreatePerson(c *gin.Context) {
	fmt.Println("IN: createPerson")
	var person Models.Person

	err := c.BindJSON(&person)
	if err != nil {
		fmt.Println("Error while Binding the data")
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}

	db, dberr := db.GetSQLConnection()
	if dberr != nil {
		fmt.Println("ERROR in DB ", dberr)
		// logginghelper.LogError("ERROR in DB ", dberr)
		return
	}

	// session := dbconnMySQL.NewSession(nil)

	// Assume that id is auto_increment and age is set to null if no value is procide to it
	stmt1, _ := db.Prepare("INSERT INTO Person (Name)" + " VALUES (?)")
	res, err := stmt1.Exec(person.Name)
	personId, err := res.LastInsertId()
	if err != nil {
		fmt.Println("ERROR in Query ", err)
		return
	}
	fmt.Println("lid: ", personId)

	// id is primary key and auto increment
	sqlstm2 := fmt.Sprintf("INSERT INTO Phone (person_id, number)"+" VALUES ('%s', '%s')", personId, person.PhoneNumber)
	_, err = db.Query(sqlstm2)
	if err != nil {
		fmt.Println("ERROR in Query 2", err)
		return
	}

	// Insert into Address table
	stmt3, _ := db.Prepare("INSERT INTO Address (city, state, street1, street2, zip_code)" + " VALUES (?,?,?,?,?)")
	res, err = stmt3.Exec(person.City, person.State, person.Street1, person.Street2, person.ZipCode)
	addressId, err := res.LastInsertId()
	if err != nil {
		fmt.Println("ERROR in Query 3", err)
		return
	}
	fmt.Println("addressId: ", addressId)

	// Insert into address_join table the person_id and address_id
	sqlstm4 := fmt.Sprintf("INSERT INTO Phone (person_id, address_id)"+" VALUES ('%d', '%d')", personId, addressId)
	_, err = db.Query(sqlstm4)
	if err != nil {
		fmt.Println("ERROR in Query 2", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Data saved successfully"})
	fmt.Println("OUT: createPerson")
}
