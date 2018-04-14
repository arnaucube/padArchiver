package main

import "github.com/fatih/color"

func check(err error) {
	if err != nil {
		color.Red(err.Error())
	}
}
