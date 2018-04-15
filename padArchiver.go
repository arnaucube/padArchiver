package padArchiver

import (
	"os"

	"github.com/fatih/color"
)

const Storage = "reposStorage"

type Repo struct {
	Dir string
}

func OpenRepo(directory string) Repo {
	//if not exist create the repos directory
	_ = os.Mkdir(Storage, os.ModePerm)

	var repo Repo
	repo.Dir = Storage + "/" + directory
	//create the repo directory
	_ = os.Mkdir(repo.Dir, os.ModePerm)
	return repo
}

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
