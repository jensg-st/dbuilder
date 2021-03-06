package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/direktiv/direktiv-apps/pkg/reusable"
	"github.com/santhosh-tekuri/jsonschema"

	_ "embed"
)

//go:embed schema.json
var schema []byte

func handler(w http.ResponseWriter, r *http.Request, ri *reusable.RequestInfo) {

	// if err := c.AddResource("schema.json", bytes.NewReader(group.Schema)); err != nil {
	// 	t.Fatal(err)
	// }
	c := jsonschema.NewCompiler()
	if err := c.AddResource("schema.json", bytes.NewReader(schema)); err != nil {
		t.Fatal(err)
	}

	sch, err := jsonschema.Compile("testdata/person_schema.json")
	if err != nil {
		log.Fatalf("%#v", err)
	}

	data, err := ioutil.ReadFile("testdata/person.json")
	if err != nil {
		log.Fatal(err)
	}

	// obj := new(Input)
	// err := reusable.Unmarshal(obj, true, r)
	// if err != nil {
	// 	reusable.ReportError(w, reusable.UnmarshallError, err)
	// 	return
	// }

	// json schema test

}

func main() {
	reusable.StartServer(handler, nil)
}
