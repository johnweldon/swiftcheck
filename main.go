package main

import (
	"html/template"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/ncw/swift"
)

var (
	OS_VARIABLES = []string{
		"OS_API_KEY",
		"OS_AUTH_URL",
		"OS_REGION_NAME",
		"OS_TENANT_ID",
		"OS_TENANT_NAME",
		"OS_USERNAME",
		"OS_PASSWORD",
		"OS_CONTAINER_NAME",
	}
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/swifttest", swifttest)
	mux.HandleFunc("/", index)
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":8888")
}

type indexData struct {
	Title     string
	Heading   string
	Action    string
	Variables []string
}

func index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("index").Parse(indexTemplate))
	d := indexData{
		Title:     "OpenStack test",
		Heading:   "OpenStack swift variables",
		Action:    "swifttest",
		Variables: OS_VARIABLES,
	}
	t.Execute(w, d)
}

type swiftData struct {
	Success bool
	Error   error
	Items   []string
	Vars    map[string]string
}

func swifttest(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("swifttest").Parse(swifttestTemplate))
	d := swiftData{}

	vars := make(map[string]string)
	for _, k := range OS_VARIABLES {
		vars[k] = r.PostFormValue(k)
	}
	d.Vars = vars

	if conn, err := getSwiftConnection(vars); err != nil {
		d.Error = err
	} else {
		if items, err := conn.ObjectNames(vars["OS_CONTAINER_NAME"], nil); err != nil {
			d.Error = err
		} else {
			d.Items = items
		}
	}
	t.Execute(w, d)
}

func getSwiftConnection(vars map[string]string) (*swift.Connection, error) {
	c := swift.Connection{
		UserName: vars["OS_USERNAME"],
		ApiKey:   vars["OS_API_KEY"],
		AuthUrl:  vars["OS_AUTH_URL"],
		Region:   vars["OS_REGION_NAME"],
		Tenant:   vars["OS_TENANT_NAME"],
		TenantId: vars["OS_TENANT_ID"],
	}
	if err := c.Authenticate(); err != nil {
		return nil, err
	}
	return &c, nil
}
