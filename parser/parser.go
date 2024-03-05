package parser

import (
	"fmt"
	"strings"

	. "github.com/AmiciaDeMonfourt/sqlparser/statement/dcl"
	. "github.com/AmiciaDeMonfourt/sqlparser/statement/ddl"
	. "github.com/AmiciaDeMonfourt/sqlparser/statement/dml"
	. "github.com/AmiciaDeMonfourt/sqlparser/statement/dql"
	. "github.com/AmiciaDeMonfourt/sqlparser/statement/tcl"
)

type Parser struct {
	originRequest string
	requestMap    map[string]string

	Truncate Truncate
	Drop     Drop
	Create   Create
	Rename   Rename
	Alter    Alter

	Select Select
	Insert Insert
	Delete Delete
	Merge  Merge
	Update Update

	Grant  Grant
	Revoke Revoke
	Deny   Deny

	BeginTC    BeginTC
	CommitTC   CommitTC
	RollbackTC RollbackTC
	SaveTc     SaveTC
}

func (p *Parser) Accept(query string) error {
	/*
	   (1) Accept user request
	   (2) Run Prepare() method
	*/
	p.originRequest = query
	return nil
}

func (p *Parser) Prepare() error {
	/*
		(1) Clear SQL-request from trash (comments, escape characters, and other...)
		(2) Run a Semantic and Syntax checkers
		(3) Smash query into small pieces (SQL statement + statement expression)
		(4) Put these pieces into requestMap
	*/
	const statements = `DENY GRANT REVOKE ALTER CREATE DROP RENAME TRUNCATE DELETE INSERT MERGE
                        UPDATE CASE FROM GROUP BY HAVING JOIN LIMIT ORDER BY SELECT USING WHERE`

	words := strings.Fields(p.originRequest)
	var query strings.Builder

	for _, word := range words {
		if strings.Contains(statements, word) {
			word = ";" + word
		}
		query.WriteString(word + " ")
		fmt.Println(word)
	}
	result := query.String()

	fmt.Print(strings.Split(result, ";"))
	return nil
}

func (p *Parser) Execute() (map[string]string, error) {
	/*
		(1) Call SQL statements handlers
		(2) Generate a result set OR
		    manipulate with data
	*/
	return nil, nil
}

func (p *Parser) PrintQuery() {
	fmt.Printf("Origin request: %s \n\n", p.originRequest)
	fmt.Printf("Request map: %s \n\n", p.requestMap)
}
