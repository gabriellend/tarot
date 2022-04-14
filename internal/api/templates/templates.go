package templates

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const fileType = ".gohtml"

var funcs = template.FuncMap{
	"AddOne": func(n int) int {
		return n + 1
	},
	"Minus64": func(a, b int64) int64 {
		return a - b
	},
	"USHours": func(s string) string {
		t, err := time.Parse("15:04", s)
		if err != nil {
			return ""
		}

		return t.Format("3:04 PM")
	},
	"PrintStringsLine": func(s []string) template.HTML {
		return template.HTML(
			fmt.Sprintf("%s", strings.Join(s, "<br>")),
		)
	},
	"IsNextPoint": func(sn int, np *int) bool {
		if np == nil {
			return false
		}
		return sn == *np
	},
	"FormatTel": func(tel string) string {
		if len(tel) < 10 {
			return ""
		}

		var str strings.Builder
		str.WriteString(tel[:3])
		str.WriteString("-")
		str.WriteString(tel[3:6])
		str.WriteString("-")
		str.WriteString(tel[6:])

		return str.String()
	},
}

func Parse(dir string) (*template.Template, error) {
	tmpl := template.New("").Funcs(funcs)
	if err := filepath.Walk(
		dir,
		func(path string, info os.FileInfo, err error) error {
			if strings.Contains(path, fileType) {
				_, err = tmpl.ParseFiles(path)
				if err != nil {
					log.Println(err)
				}
			}

			return err
		},
	); err != nil {
		return nil, err
	}

	return tmpl, nil
}
