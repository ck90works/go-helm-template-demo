package main

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

var tpl *template.Template

type pod struct {
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
	Type      string
	DataKey   string
	DataValue string
}

func init() {
	tpl = template.Must(template.New("").Funcs(sprig.FuncMap()).ParseGlob("templates/*"))
}

func main() {
	help := `
	Info:

		Diese Go-Anwendung demonstriert die Go Template Anwendung
		mithilfe von fünf gängigen (teils komplexen) Datenstrukturen, 
		die in Helm Charts verwendet werden.

		Dieses Tool ist so aufgebaut, dass man dies auch als Sandbox
		verwenden kann.


	Kommandos:

		Für die Ausführung dieser Anwendung muss ein Argument mitgeteilt werden,
		folgende fünf Argumente sind valide:

			slice
			map
			struct
			slice-struct
			struct-of-structs


		Beispielhaft würde ein Aufruf dieser Anwendung in dev so aussehen:

		> go run yaml_generator.go slice-struct

		
		Oder, sollte die Anwendung kompiliert sein, würde sie so aussehen:

		> yaml_generator slice-struct 
	`

	if len(os.Args) > 1 {
		if os.Args[1] == "slice" {
			range_over_slice()
		} else if os.Args[1] == "map" {
			range_over_map()
		} else if os.Args[1] == "struct" {
			range_over_struct()
		} else if os.Args[1] == "slice-struct" {
			range_over_slice_struct()
		} else if os.Args[1] == "struct-of-structs" {
			execute_struct_from_tpl()
		} else {
			fmt.Println(help)
		}
	} else {
		fmt.Println(help)
	}
}

func range_over_slice() {
	names := []string{
		"   secret-sa-sample",
		"secret_Sa    ",
		" another _secret-sa",
		"ordentlicher-sa-sample",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl_slice.goyaml", names)
	if err != nil {
		log.Fatalln(err)
	}
}

func range_over_map() {
	multiline_string := `|
	TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNldGV0dXIgc2FkaXBzY2luZyBlbGl0ciwgc
	2VkIGRpYW0gbm9udW15IGVpcm1vZCB0ZW1wb3IgaW52aWR1bnQgdXQgbGFib3JlIGV0IGRvbG9yZS
	BtYWduYSBhbGlxdXlhbSBlcmF0LCBzZWQgZGlhbSB2b2x1cHR1YS4gQXQgdmVybyBlb3MgZXQgYWN
	jdXNhbSBldCBqdXN0byBkdW8gZG9sb3JlcyBldCBlYSByZWJ1bS4gU3RldCBjbGl0YSBrYXNkIGd1
	YmVyZ3Jlbiwgbm8gc2VhIHRha2ltYXRhIHNhbmN0dXMgZXN0IExvcmVtIGlwc3VtIGRvbG9yIHNpd
	CBhbWV0LiBMb3JlbSBpcHN1bSBkb2xvciBzaXQgYW1ldCwgY29uc2V0ZXR1ciBzYWRpcHNjaW5nIG`

	eine_map := map[string]string{
		"random.cfg": multiline_string,
		"user":       "password",
		"vault_seal": "ganz geheim",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl_map.goyaml", eine_map)
	if err != nil {
		log.Fatalln(err)
	}
}

func range_over_struct() {
	eine_struct := pod{
		Name:         "test-app",
		Namespace:    "production",
		Image:        "nginx",
		ImageName:    "go_nginx",
		ImageVersion: "1.14.2",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl_struct.goyaml", eine_struct)
	if err != nil {
		log.Fatalln(err)
	}
}

func range_over_slice_struct() {
	eine_slice_aus_structs := []pod{
		{
			Name:         "test-new-feature-app",
			Namespace:    "dev",
			Image:        "nginx",
			ImageName:    "go_nginx",
			ImageVersion: "1.14.2",
		},
		{
			Name:         "test-staging-app",
			Namespace:    "staging",
			Image:        "nginx",
			ImageName:    "go_nginx",
			ImageVersion: "1.14.2",
		},
		{
			Name:         "test-app",
			Namespace:    "production",
			Image:        "nginx",
			ImageName:    "go_nginx",
			ImageVersion: "1.14.2",
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl_slice_struct.goyaml", eine_slice_aus_structs)
	if err != nil {
		log.Fatalln(err)
	}
}

func execute_struct_from_tpl() {

	eine_struct_aus_structs := struct {
		Metadaten []metadaten
		Secrets   []secrets
	}{
		[]metadaten{
			{
				Name:        "team-1-app",
				Namespace:   "team-1",
				Annotations: `kubernetes.io/service-account.name: "team-1-sa-name"`,
			},
			{
				Name:        "team-2-app",
				Namespace:   "team-2",
				Annotations: `kubernetes.io/service-account.name: "team-2-sa-name"`,
			},
		},
		[]secrets{
			{Type: "Sensitive",
				DataKey:   "random.cfg",
				DataValue: "path/to/rome"},
			{Type: "Opaque",
				DataKey: "some.cfg",
				DataValue: `| 
				ein weiterer Multiline Text 
				der irgendwelche zufaelligen 
				Saetze enthaelt`},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl_struct_from_tpl.goyaml", eine_struct_aus_structs)
	if err != nil {
		log.Fatalln(err)
	}
}
