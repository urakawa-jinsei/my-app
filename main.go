package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

func main() {
	t := template.New("").
		Funcs(template.FuncMap{
			"FormatDateTime": func(format string, d time.Time) string {
				if d.IsZero() {
					return ""
				}
				return d.Format(format)
			}})
	tmpl := `{{FormatDateTime "2006 年 01 月 02 日" .}}` // 実行するテンプレート
	t = template.Must(t.Parse(tmpl))
	err := t.Execute(os.Stdout, time.Now())
	if err != nil {
		log.Fatal(err)
	}
}
