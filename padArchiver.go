package padArchiver

import (
	"os"

	"github.com/fatih/color"
)

//Storage is the directory where are stored the repos
const Storage = "reposStorage"

//Repo is the directory where is placed the repository of pads
type Repo struct {
	Dir string
}

//OpenRepo opens a repo from the directory
func OpenRepo(directory string) Repo {
	//if not exist create the repos directory
	_ = os.Mkdir(Storage, os.ModePerm)

	var repo Repo
	repo.Dir = Storage + "/" + directory
	//create the repo directory
	_ = os.Mkdir(repo.Dir, os.ModePerm)
	return repo
}

//StorePad gets a pad from the link, and stores it into local directory. Then also, adds the file to IPFS.
func (repo *Repo) StorePad(link string, directory string, title string, ipfsActive bool) (string, error) {
	path, err := repo.GetPad(link, "md", directory, title)
	if err != nil {
		color.Red(err.Error())
		return "", err
	}

	if !ipfsActive {
		return "", nil
	}
	hash, err := IpfsAdd(path)
	if err != nil {
		color.Red(err.Error())
		return hash, err
	}

	err = AddLineToFile(path, "IPFS hash of this document: "+hash)
	if err != nil {
		color.Red(err.Error())
		return hash, err
	}
	// TODO
	// err = repo.GitUpdate("update commit")
	// if err != nil {
	// 	color.Red(err.Error())
	// 	return hash, err
	// }

	return hash, nil
}
