package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	padArchiver ".."
	"github.com/fatih/color"
)

const checkIcon = "\xE2\x9C\x94 "

func main() {
	asciiart := `
.
.                _                   _      _
.               | |   /\            | |    (_)
 _ __   __ _  __| |  /  \   _ __ ___| |__  ___   _____ _ __
| '_ \ / _  |/ _  | / /\ \ | '__/ __| '_ \| \ \ / / _ \ '__|
| |_) | (_| | (_| |/ ____ \| | | (__| | | | |\ V /  __/ |
| .__/ \__,_|\__,_/_/    \_\_|  \___|_| |_|_| \_/ \___|_|		- cli
| |
|_|

		`
	color.Blue(asciiart)
	fmt.Println("							v0.0.1")
	color.Blue("https://github.com/arnaucode/padArchiver")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	newcommand := bufio.NewReader(os.Stdin)
	fmt.Print("Please select command number")
	options := `
	1 - Store Pad (to local directory and IPFS)
	2 - IPFS hash to file
	0 - Exit cli
option to select: `
	for {
		fmt.Print(options)

		option, _ := newcommand.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			fmt.Println("selected 1 - Store Pad (to local directory and IPFS)")
			option1()
			break
		case "2":
			fmt.Println("selected 2 - IPFS hash to file")
			option2()
			break
		case "0":
			fmt.Println("selected 0 - exit cli")
			os.Exit(3)
			break
		default:
			fmt.Println("Invalid option")
			break
		}
	}
}
func option1() {
	newcommand := bufio.NewReader(os.Stdin)
	fmt.Print("	Enter the repo ID (name): ")
	repoID, _ := newcommand.ReadString('\n')
	repoID = strings.Replace(repoID, "\n", "", -1)

	newcommand = bufio.NewReader(os.Stdin)
	fmt.Print("	Enter the pad link: ")
	link, _ := newcommand.ReadString('\n')
	link = strings.Replace(link, "\n", "", -1)

	newcommand = bufio.NewReader(os.Stdin)
	fmt.Print("	Enter the subdirectory: ")
	subdirectory, _ := newcommand.ReadString('\n')
	subdirectory = strings.Replace(subdirectory, "\n", "", -1)

	newcommand = bufio.NewReader(os.Stdin)
	fmt.Print("	Enter the pad Title: ")
	title, _ := newcommand.ReadString('\n')
	title = strings.Replace(title, "\n", "", -1)

	repo := padArchiver.OpenRepo(repoID)

	ipfsHash, err := repo.StorePad(link, subdirectory, title, true)
	if err != nil {
		color.Red(err.Error())
	} else {
		fmt.Println("IPFS hash: " + ipfsHash)
		color.Green(checkIcon + "Pad stored in IPFS and Git")
	}
}
func option2() {
	newcommand := bufio.NewReader(os.Stdin)
	fmt.Print("	Enter the IPFS hash: ")
	hash, _ := newcommand.ReadString('\n')
	hash = strings.Replace(hash, "\n", "", -1)
	err := padArchiver.IpfsGet(hash, hash+".md")
	if err != nil {
		color.Red(err.Error())
	} else {

		color.Green(checkIcon + "File downloaded from IPFS network")
		fmt.Print("File stored in: ")
		color.Blue(padArchiver.GettedPads + "/" + hash + ".md")
	}
}
