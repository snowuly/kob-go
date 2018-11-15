package kob

import "regexp"

var rkey = regexp.MustCompile(`:[\w]+`)

const rvalue = `([^/]+)`

func PathToReg(path string) (reg *regexp.Regexp, keys []string, err error) {
	path = regexp.QuoteMeta(path)
	reg, err = regexp.Compile("^" + rkey.ReplaceAllStringFunc(path, func(m string) string {
		keys = append(keys, m[1:])
		return rvalue
	}) + "$")
	return
}
