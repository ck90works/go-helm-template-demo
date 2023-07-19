package main

import (
	"log"
	"os"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

var tpl *template.Template

type schiff struct {
	Name         string
	Namespace    string
	Image        string
	ImageName    string
	ImageVersion string
}

type metadaten struct {
	Name        string
	Namespace   string
	Annotations string
}

type secrets struct {
	Type string
	Data string
}

type schiff_config struct {
	Metadaten []metadaten
	Secrets   []secrets
}

func init() {
	tpl = template.Must(template.New("").Funcs(sprig.FuncMap()).ParseGlob("templates/*"))
}

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "slice" {
			range_over_slice()
		} else if os.Args[1] == "map" {
			range_over_map()
		} else if os.Args[1] == "struct" {
			range_over_struct()
		} else if os.Args[1] == "slice-struct" {
			range_over_slice_struct()
		} else if os.Args[1] == "struct-from-tpl" {
			execute_struct_from_tpl()
		}
	}

}

func range_over_slice() {
	namespaces := []string{
		"   namespace1",
		"namespace2",
		"namespace3",
		"namespace4",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl1.goyaml", namespaces)
	if err != nil {
		log.Fatalln(err)
	}
}

func range_over_map() {
	eine_map := map[string]string{
		"certificate":    "irgenein base64 wert",
		"user":           "password",
		"vault_seal_key": "ganz geheim",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl2.goyaml", eine_map)
	if err != nil {
		log.Fatalln(err)
	}
}

func range_over_struct() {
	eine_struct := schiff{
		Name:         "aid_test_1",
		Namespace:    "aid_test_1_namespace",
		Image:        "nginx",
		ImageName:    "aid_nginx",
		ImageVersion: "1.14.2",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl3.goyaml", eine_struct)
	if err != nil {
		log.Fatalln(err)
	}
}

func range_over_slice_struct() {
	aid := metadaten{
		Name:        "aid_testname",
		Namespace:   "aid_test_namespace",
		Annotations: "aid_test_annotation",
	}

	zabbix := metadaten{
		Name:        "zabbix_testname",
		Namespace:   "zabbix_test_namespace",
		Annotations: "zabbix_test_annotation",
	}

	aid_secrets := secrets{
		Type: "Sensitive",
		Data: "certs: ",
	}

	zabbix_secrets := secrets{
		Type: "Opaque",
		Data: "vaults: ",
	}

	schiff_konfigurationen := schiff_config{
		[]metadaten{aid, zabbix},
		[]secrets{aid_secrets, zabbix_secrets},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl4.goyaml", schiff_konfigurationen)
	if err != nil {
		log.Fatalln(err)
	}
}

func execute_struct_from_tpl() {
	eine_struct := schiff{
		Name:         "aid_test_1",
		Namespace:    "aid_test_1_namespace",
		Image:        "nginx",
		ImageName:    "aid_nginx",
		ImageVersion: "1.14.2",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl_execute_struct_from_tpl.goyaml", eine_struct)
	if err != nil {
		log.Fatalln(err)
	}
}
