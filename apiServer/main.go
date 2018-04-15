package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	padArchiver ".."
	"github.com/gorilla/mux"
)

//PadModel is the data structure for each pad
type PadModel struct {
	Link     string `json:"link"`
	Dir      string `json:"dir"`
	Title    string `json:"title"`
	IpfsHash string `json:"ipfsHash"`
}

//Repo contains all the pads --currently not used--
type Repo struct {
	Pads []string `json:"pads"`
}

func main() {
	readConfig("config.json")

	router := mux.NewRouter()
	router.HandleFunc("/repos", GetReposList).Methods("GET")
	router.HandleFunc("/repos/{repoid}", GetRepoIDList).Methods("GET")
	router.HandleFunc("/repos/{repoid}/pad", PostStorePad).Methods("POST")

	log.Println("padArchiver API server running")
	log.Print("port: ")
	log.Println(config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, router))
}

//GetReposList is the endpoint to get the list of current repos
func GetReposList(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open(padArchiver.Storage)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	list, _ := file.Readdirnames(0) // 0 to read all files and folders

	var repos []string
	for _, name := range list {
		repos = append(repos, strings.Replace(name, padArchiver.Storage+"/", "", -1))
	}

	jResp, err := json.Marshal(repos)
	check(err)
	fmt.Fprintln(w, string(jResp))
}

//GetRepoIDList is the endpoint to get one repo by id
func GetRepoIDList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	repoid := vars["repoid"]

	fileList := []string{}
	err := filepath.Walk(padArchiver.Storage+"/"+repoid, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})

	var files []string
	for _, file := range fileList {
		files = append(files, strings.Replace(file, padArchiver.Storage+"/", "", -1))
	}
	jResp, err := json.Marshal(files)
	check(err)
	fmt.Fprintln(w, string(jResp))
}

//PostStorePad is the endpoint to post the signal to store one pad
func PostStorePad(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	repoid := vars["repoid"]

	//open the repo
	repo := padArchiver.OpenRepo(repoid)
	//get the pad json
	var pad PadModel
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&pad)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	ipfsHash, err := repo.StorePad(pad.Link, pad.Dir, pad.Title, true)
	if err != nil {
		http.Error(w, "error storing pad", http.StatusConflict)
	}
	pad.IpfsHash = ipfsHash
	jResp, err := json.Marshal(pad)
	check(err)
	fmt.Fprintln(w, string(jResp))
}
