package main

import (
	"fmt"
	//"log"
	"sync"

	"github.com/GregorioMonari/testfix-sepago/sepa"
	"github.com/GregorioMonari/testfix-sepago/sepa/sparql"
)

func main() {
	fmt.Println("## LET'S START ##")

	//NewClient configuration. Ports used in my computer 8600 e 9600 (set in docker image):
	config := sepa.Configuration{
		Host:  "localhost",
		Ports: sepa.PortsType{Http: 8600, Ws: 9600},
		Paths: sepa.PathsType{Query: "/query", Update: "/update", Subscribe: "/subscribe"},
	}
	cli := sepa.NewClient(config) //se riesci jsap

	//Insert new data:
	err := cli.Update(`INSERT DATA
	{
		<uno> <duo> <tre>
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
	fmt.Println("QUI")
	//Subscribe. Create a websocket connection:
	var wg sync.WaitGroup
	wg.Add(1) //used to wait for multiple goroutines to finish
	fmt.Println("QUI")
	sub, _ := cli.Subscribe("Select * Where {Graph <http://example/testgraph> {?s ?p ?o} }", func(notification *sparql.Notification) {

		fmt.Println("Printing notification:")
		/*for _, solution := range notification.AddedResults.Solutions() {

			//Sorting keys to get a predictable output
			// see https://blog.golang.org/go-maps-in-action#TOC_7.
			var keys []string
			for k := range solution {
				keys = append(keys, k)
			}
			sort.Strings(keys)

			for _, key := range keys {
				fmt.Print(key, ": ", solution[key], " ")
			}
			fmt.Println(".")
		}*/
		fmt.Println("qui")

		wg.Done()
	})
	fmt.Println("## SUBSCRIPTION DONE ##")
	err = cli.Update(`INSERT DATA
	{
		<http://example/lang> <http://example/thebest> "go".
		<http://example/lang> <http://example/theworst> "visual basic"
	}`)

	//log.Println(err)
	wg.Wait()
	sub.Unsubscribe()
	fmt.Println("")
	fmt.Println("Unsubscribed")
}
