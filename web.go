package main

import (
	"fmt"
	"net/http"
)

const tpl = `
<div>
	{{range .Thoughts}}
	<div>
		<div>{{.Date}}</div>
		<a href={{.Location}}> {{.Title}}</a>
	</div>
    {{end}}
</div>
`

type Thought struct {
	Title    string
	Date     string
	Location string
}

type IndexPageData struct {
	PageTitle string
	Thoughts  []Thought
}

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	files, err := ioutil.ReadDir(renderDir())
// 	if err != nil {
// 		return
// 	}
// 	var thoughs []Thought
// 	for _, f := range files {
// 		thoughs = append(thoughs, Thought{
// 			Title:    f.Name(),
// 			Date:     f.ModTime().Format("2006/01/02"),
// 			Location: "/" + f.Name(),
// 		})
// 	}
// 	data := IndexPageData{
// 		PageTitle: "My Thoughts",
// 		Thoughts:  thoughs,
// 	}

// 	t, err := template.New("webpage").Parse(tpl)
// 	t.Execute(w, data)

// }

// ServeStatic - server static html file
func serveStatic() {
	http.Handle("/", http.FileServer(http.Dir(renderDir())))
	fmt.Println("All your thoughs is here: http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
