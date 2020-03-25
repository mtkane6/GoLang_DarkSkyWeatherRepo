package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename, err := os.Open("stage1.html")
	
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func stageHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/stage1/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles("stage1.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/stage1/", stageHandler)
	http.ListenAndServe(":8080", nil)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// file, err := os.Open("/SPjson.json")
	// if err != nil {
	// }
	// defer file.Close()
	// f, _ := ioutil.ReadAll(file)
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte(f))
	// })

	// newStPass := domain.StevensPassStruct{}

	// url := os.Args[1:]
	// resp, err := http.Get(strings.Join(url, ""))
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	// }

	// b, err := ioutil.ReadAll(resp.Body)
	// json.Unmarshal([]byte(b), &newStPass)
	// resp.Body.Close()

	// snowfall := domain.SnowfallResponse{}
	// snowfall.NewSnow = newStPass.HeaderSettings.SortOrder
	// fmt.Println()
	// fmt.Println("Snowfall overnight: ", snowfall.NewSnow, "\"")
	// fmt.Println()
}
