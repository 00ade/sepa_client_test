package main

import (
	"fmt"

	"github.com/GregorioMonari/testfix-sepago/sepa"
)

func main() {
	fmt.Println("## LET'S START ##")

	//NewClient configuration. Ports used in my computer 8600 e 9600 (set in docker image):
	config := sepa.Configuration{
		Host:  "localhost",
		Ports: sepa.PortsType{Http: 8600, Ws: 9600},
		Paths: sepa.PathsType{Query: "/query", Update: "/update", Subscribe: "/subscribe"},
	}
	cli := sepa.NewClient(config)

	//Insert new data:
	err := cli.Update(`INSERT DATA
	{
		<io> <FUNZIONA> <AA>
	}`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("## UPDATE DONE ##")

	//Do a query:
	res, err := cli.Query("Select * Where { ?s ?p ?o}")
	if err != nil {
		fmt.Println(err)
		return
	}
	vars := res.Vars()
	fmt.Println("#### STARTING QUERY ####")
	for _, varable := range vars {
		for _, term := range res.Bindings()[varable] {
			fmt.Println(term)
		}
	}
	fmt.Println("#### QUERY STOP ####")

}
