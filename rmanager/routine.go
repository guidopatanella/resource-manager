package rmanager

import "strings"
import "strconv"

type StackRow struct {
	Function string
	Program string
}
type Routine struct {
	Number int
	State string
	Stack []StackRow
}

func (r *Routine) ParseInit(row string) {
	parts := strings.Split(row, " ") 
	number, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	r.Number = number
	r.State = parts[2]
	if len(parts)>3 {
		r.State += " " + parts[3]
	}
	r.State = r.State[1:len(r.State)-2]
}

func (r *Routine) ParseStackRow(row1 string, row2 string) {
	row1 = strings.TrimSpace(row1)
	row2 = strings.TrimSpace(row2)
	// parse function name
	funcParts := strings.Split(row1, "(")
	functionParsed := funcParts[0]
	if len(funcParts)>2 {
		// scenario of method i.e.: "func (x *blah) methodname()"
		functionParsed += "(" + funcParts[1]
	}
	if strings.Index(row1, "created by")==0 {
		functionParsed = strings.Replace(row1,"created by","", -1)
	}
	// parse program
	programParsed := strings.Split(row2, " ")[0]
	r.Stack = append(r.Stack, StackRow{
		Function: functionParsed,
		Program: programParsed,
	})
}

