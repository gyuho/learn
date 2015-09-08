package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

func main() {
	tagName := "{{.BranchName}}_{{.Type}}"
	tagStruct := struct {
		BranchName string
		Type       string
	}{
		"gyuho",
		"prod",
	}
	buf := new(bytes.Buffer)
	if err := template.Must(template.New("tmpl").Parse(tagName)).Execute(buf, tagStruct); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
	// gyuho_prod

	queryStruct := struct {
		SchemaName string
		TableName  string
		Slice      []map[string]string
		LastIndex  int
	}{
		"my",
		"table",
		[]map[string]string{
			map[string]string{"key": "VARCHAR(100) PRIMARY KEY NOT NULL"},
			map[string]string{"value1": "INTEGER"},
			map[string]string{"value2": "INTEGER"},
		},
		2,
	}
	var queryTmpl = `CREATE TABLE IF NOT EXISTS {{.SchemaName}}.{{.TableName}}  ({{$lastIndex := .LastIndex}}
{{range $index, $valueMap := .Slice}}{{if ne $lastIndex $index}}{{range $key, $value := $valueMap}}	{{$key}} {{$value}},{{end}}
{{else}}{{range $key, $value := $valueMap}}	{{$key}} {{$value}}{{end}}
{{end}}{{end}});`
	tb := new(bytes.Buffer)
	if err := template.Must(template.New("tmpl").Parse(queryTmpl)).Execute(tb, queryStruct); err != nil {
		log.Fatal(err)
	}
	fmt.Println(tb.String())
	/*
	   CREATE TABLE IF NOT EXISTS my.table  (
	   	key VARCHAR(100) PRIMARY KEY NOT NULL,
	   	value1 INTEGER,
	   	value2 INTEGER
	   );
	*/
}
