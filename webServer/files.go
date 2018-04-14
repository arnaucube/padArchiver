package main

import (
	"io/ioutil"

	"github.com/fatih/color"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

func readFile(path string) string {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		color.Red(path)
	}
	check(err)
	return string(dat)
}

func fileToHTML(path string) (string, error) {
	mdcontent := readFile(path)
	htmlcontent := string(blackfriday.Run([]byte(mdcontent)))
	return htmlcontent, nil
}
