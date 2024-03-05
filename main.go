package main

import "github.com/AmicieDeMonfourt/sqlparser/parser"

func main() {
	var parser parser.Parser

	parser.Accept(`DROP TABLE customers;
				CREATE TABLE customers;
				SELECT * FROM table1 WHERE person_id < 10;`)

	parser.Prepare()

	parser.Execute()
}
