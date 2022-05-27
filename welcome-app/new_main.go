package main

import (
	"fmt"
	"log"
	"net/http"
)

var page = `<!DOCTYPE html>
<head>
<meta charset="UTF-8">
<link rel="stylesheet" href="/static/stylesheets/welcome-template.css">
<title>Aerothon 4.0 - Team FULL_TOGA</title>
</head>

<html>
<body>
<p><b><center>Aerothon 4.0 - Team FULL_TOGA</center></b></p>

<div class="dropdown_a">
<form action="/upload" method = "get">
   <b>Select Stack</b>
   <select name = "action" size = "1" id ="action-choice">
   <ul class = "active">
	  <option value ="mean">MEAN Stack</option>
	  <option value ="mern">MERN Stack</option>
	  <option value ="mevn">MEVN Stack</option>
	  <option value ="lamp">LAMP Stack</option>
	  <option value ="...">...</option>
   </select> 
   </ul>
</div>

<div class = "button_download"><button>Download</button></div>
</form>

<div class="dropdown_b">
<form>
   <b>Backend</b>
   <select>
   <ul class = "active">
	  <option value ="node">node.js</option>
	  <option value ="...">...</option> 
   </ul>
   </select>
</form>
</div>

<div class="dropdown_c">
<form>
	<b>Frontend</b>
	<select>
	<ul class="active">
	<option value ="angular">angular.js</option>
	<option value ="...">...</option>
	</select>
	</ul>
</form>
</div>

<div class="dropdown_d">
<form>
	<b>Database</b>
	<select>
	<ul class="active">
	<option value ="mongodb">MongoDB</option>
	<option value ="...">...</option>
	</select>
	</ul>
</form>
</div>

</body>
</html>
`

// func handler(w http.ResponseWriter, r *http.Request) {
// 	template.Must(template.ParseFiles("index.html")).Execute(w, nil)
// }

func m(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(page))
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
		w.Write([]byte(r.Form["action"][0]))
		fmt.Println(r.Form["action"])
	}
}
func main() {
	log.Println("started")
	// http.HandleFunc("/", handler)
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", m)
	http.HandleFunc("/upload", upload)
	fmt.Println("Serving on PORT 8080")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
