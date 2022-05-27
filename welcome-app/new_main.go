package main

import (
	"fmt"
	"log"
	"net/http"
)

var page = `<!DOCTYPE html>
<head>
<meta charset="UTF-8">
<style>
 body{
    margin: 0;
    padding: 0;
    font-family: sans-serif;
 }
 
 .dropdown_a
 {
  position: absolute;
  top: 25%;
  left: 20%;
  transform: translate(-50%, -50%);
  width: auto;
 }

 .dropdown_b
 {
  position: absolute;
  top: 25%;
  left: 40%;
  transform: translate(-50%, -50%);
  width: 100px;
 }

 .dropdown_c
 {
  position: absolute;
  top: 25%;
  left: 60%;
  transform: translate(-50%, -50%);
  width: 100px;
 }

 .dropdown_d
 {
  position: absolute;
  top: 25%;
  left: 80%;
  transform: translate(-50%, -50%);
  width: 100px;
 }

 .button_download
   {
      position: absolute;
      top: 50%;
      left: 50%;
   }
 
 ul
 {
    position: absolute;
    margin: 0;
    padding: 0;
    width: 100%;
    background: #ccc;
    transform-origin: top;
    transform: perspective(1000px)rotateX(-90deg);
    transition: 0.5s;
 }
 ul.active{
   transform: perspective(1000px)rotateX(0deg);
 }
 ul li{
    list-style: none;
 }
</style>
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
	  <option value ="python">Python</option>
	  <option value ="ruby">Ruby</option>
	  <option value ="django">Django</option>
	  <option value ="firebase">Firebase</option>
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
	<option value ="react">react.js</option>
	<option value ="reactn">react native</option>
	<option value ="vue">vue.js</option>
	<option value ="html">HTML</option>
	<option value ="bootstrap">Bootstrap</option>
	<option value ="foundation">Foundation</option>
	<option value ="bulma">Bulma</option>
	<option value ="materialize">Materialize</option>
	<option value ="vanilla">Vanilla JS</option>
	<option value ="Svelte">Svelte</option>
	<option value ="flutter">Flutter</option>
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
	<option value ="mysql">MySQL</option>
	<option value ="postgre">PostgreSQL</option>
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
		// w.Write([]byte(r.Form["action"][0]))
		// fmt.Println(r.Form["action"][0])
		if r.Form["action"][0] == "mean" {
			// Print the github link for the mean stack template
			http.Redirect(w, r, fmt.Sprintf("https://github.com/agonxgashi/MEAN-template"), 302)
		} else if r.Form["action"][0] == "mern" {
			// Print the github link for the mern stack template
			http.Redirect(w, r, fmt.Sprintf("https://github.com/bradwindy/mern-stack-template"), 302)
			http.Redirect(w, r, fmt.Sprintf("https://github.com/amazingandyyy/mern"), 302)
		} else if r.Form["action"][0] == "mevn" {
			// Print the github link for the mevn stack template
			http.Redirect(w, r, fmt.Sprintf("https://github.com/aturingmachine/mevn-stack"), 302)
		} else if r.Form["action"][0] == "lamp" {
			// Print the github link for the lamp stack template
			http.Redirect(w, r, fmt.Sprintf("https://github.com/teddysun/lamp"), 302)
		} else {
			http.Redirect(w, r, fmt.Sprintf("Please select a stack"), 302)
		}
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
