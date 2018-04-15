package padArchiver

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//GetPad gets the pad from the link, and stores it into local directory
func (repo *Repo) GetPad(link string, extension string, directory string, title string) (string, error) {
	if extension != "md" && extension != "txt" && extension != "html" && extension != "pdf" && extension != "odt" {
		return "", errors.New("No valid extension")
	}
	format := extension
	if extension == "md" {
		format = "markdown"
		extension = "md"
	}

	//create the pads directory
	_ = os.Mkdir(repo.Dir+"/"+directory, os.ModePerm)

	completeLink := link + "/export/" + format

	//get the content from the url
	r, err := http.Get(completeLink)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer r.Body.Close()
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s", err)
		return "", err
	}

	//save the content into a file
	err = ioutil.WriteFile(repo.Dir+"/"+directory+"/"+title+"."+extension, content, 0644)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return repo.Dir + "/" + directory + "/" + title + "." + extension, nil
}
