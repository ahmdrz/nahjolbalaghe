package main

import (
	"io/ioutil"
	"strings"
)

func Decode(filename string) (map[string]string, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return map[string]string{}, err
	}
	out := make(map[string]string, 0)
	result := string(bytes)
	result = strings.TrimSpace(result)
	lines := strings.Split(result, "\n")
	var flag bool
	var item string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "[[") && strings.HasSuffix(line, "]]") {
			cmd := line[2 : len(line)-2]
			flag = cmd != "end"
			item = cmd
			continue
		}
		if flag {
			out[item] += line + "\n"
		}
	}
	for key, value := range out {
		if strings.HasSuffix(value, "\n") {
			out[key] = value[:len(value)-1]
		}
	}
	return out, nil
}
