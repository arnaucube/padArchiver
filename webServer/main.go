package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	padArchiver ".."
	"github.com/gorilla/mux"
)

type ItemModel struct {
	Name  string
	Path  string
	IsDir bool
}
type MenuModel struct {
	PageTitle string
	Items     []ItemModel
}
type PageModel struct {
	Title       string
	MenuContent template.HTML
	Content     template.HTML
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", getDir).Methods("GET")
	router.HandleFunc("/dir/{path}", getDir).Methods("GET")
	router.HandleFunc("/page/{path}", getPage).Methods("GET")

	log.Println("padArchiver web server running")
	log.Print("port: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func generateMenuItems(dirpath string) MenuModel {
	var menuPage MenuModel
	menuPage.PageTitle = "padArchiver - Menu"
	_ = filepath.Walk(padArchiver.Storage+dirpath, func(path string, f os.FileInfo, err error) error {
		if path != padArchiver.Storage {
			path = strings.Replace(path, padArchiver.Storage, "", -1)
			var item ItemModel
			item.Name = path
			path = strings.Replace(path, "/", "%", -1)
			if f.IsDir() {
				item.Path = "/dir/" + path
			} else {
				item.Path = "/page/" + path
			}
			item.IsDir = f.IsDir()
			menuPage.Items = append(menuPage.Items, item)
		}
		return nil
	})
	return menuPage
}
func generateMenuHTML(dirpath string) template.HTML {
	menuItems := generateMenuItems(dirpath)
	tmplMenu := template.Must(template.ParseFiles("templates/menuTemplate.html"))
	var tpl bytes.Buffer
	err := tmplMenu.Execute(&tpl, menuItems)
	check(err)
	return template.HTML(tpl.String())
}
func getDir(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var dirpath string
	if _, ok := vars["path"]; ok {
		dirpath = vars["path"]
		dirpath = strings.Replace(dirpath, "%", "/", -1)
	}
	var page PageModel
	page.Title = dirpath
	if dirpath == "" {
		page.Title = "padArchiver"
	}
	page.MenuContent = generateMenuHTML(dirpath)
	content, err := fileToHTML("templates/wellcome.md")
	check(err)
	page.Content = template.HTML(content)

	tmpl := template.Must(template.ParseFiles("templates/pageTemplate.html"))
	tmpl.Execute(w, page)
}

func getPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	path := vars["path"]
	path = strings.Replace(path, "%", "/", -1)
	path = padArchiver.Storage + path

	content, err := fileToHTML(path)
	check(err)

	var page PageModel
	page.Title = path
	page.Content = template.HTML(content)

	page.MenuContent = generateMenuHTML("")

	tmplPage := template.Must(template.ParseFiles("templates/pageTemplate.html"))
	tmplPage.Execute(w, page)
}
