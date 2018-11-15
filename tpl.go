package kob

import (
	"html/template"
	"strings"
)

var cache = make(map[string]*template.Template)

func GetTpl(files ...string) (t *template.Template, err error) {
	key := strings.Join(files, "")
	t, ok := cache[key]
	if ok {
		return t, nil
	}
	if t, err = template.ParseFiles(files...); err == nil {
		cache[key] = t
	}
	return
}

func GetTplGlob(pattern string) (t *template.Template, err error) {
	t, ok := cache[pattern]
	if ok {
		return t, nil
	}

	if t, err = template.ParseGlob(pattern); err == nil {
		cache[pattern] = t
	}
	return
}
