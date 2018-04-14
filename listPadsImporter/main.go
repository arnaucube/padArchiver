package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	padArchiver ".."
	"github.com/fatih/color"
)

type PadModel struct {
	Link  string `json:"link"`
	Dir   string `json:"dir"`
	Title string `json:"title"`
}
type ListModel struct {
	RepoID string     `json:"repoid"`
	Pads   []PadModel `json:"pads"`
}

func readList(path string) ListModel {
	file, err := ioutil.ReadFile(path)
	check(err)
	content := string(file)
	var list ListModel
	json.Unmarshal([]byte(content), &list)
	return list
}

func main() {
	asciiart := `
.
.                _                   _      _
.               | |   /\            | |    (_)
 _ __   __ _  __| |  /  \   _ __ ___| |__  ___   _____ _ __
| '_ \ / _  |/ _  | / /\ \ | '__/ __| '_ \| \ \ / / _ \ '__|
| |_) | (_| | (_| |/ ____ \| | | (__| | | | |\ V /  __/ |
| .__/ \__,_|\__,_/_/    \_\_|  \___|_| |_|_| \_/ \___|_|		- listPadsImporter
| |
|_|

	`
	color.Blue(asciiart)
	fmt.Println("							v0.0.1")
	color.Blue("https://github.com/arnaucode/padArchiver")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	list := readList("list.json")

	//open the repo
	repo := padArchiver.OpenRepo(list.RepoID)
	fmt.Println("repo opened")
	for _, pad := range list.Pads {
		fmt.Println("importing pad:")
		fmt.Println("	link: " + pad.Link)
		fmt.Println("	dir: " + pad.Dir)
		fmt.Println("	title: " + pad.Title)
		ipfsHash, err := repo.StorePad(pad.Link, pad.Dir, pad.Title)
		check(err)
		fmt.Println("	ipfs hash: " + ipfsHash)
	}
	color.Green("listPadsImporter finished")
}
