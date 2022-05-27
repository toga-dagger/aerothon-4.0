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
	font-size: 15px;
 }
 
 .dropdown_a
 {
  position: absolute;
  top: 38%;
  left: 20%;
  transform: translate(-50%, -50%);
  width: auto;
 }
 .dropdown_b
 {
  position: absolute;
  top: 38%;
  left: 40%;
  transform: translate(-50%, -50%);
  width: 100px;
 }
 .dropdown_c
 {
  position: absolute;
  top: 38%;
  left: 60%;
  transform: translate(-50%, -50%);
  width: 100px;
 }
 .dropdown_d
 {
  position: absolute;
  top: 38%;
  left: 80%;
  transform: translate(-50%, -50%);
  width: 100px;
 }
 .button_download
   {
      position: absolute;
      top: 50%;
      left: 45%;
	  text-align: center;
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

 .note {
	border: 5px solid #FFFF00;
	text-align: center;
	top: 38%;
	font-size: 15px;
	padding: 10px;
}

</style>
<title>Aerothon 4.0 - Team FULL_TOGA</title>
</head>
<html>
<body>
<p><b><center>Aerothon 4.0 - Team FULL_TOGA</center></b></p>
<form action="/upload" method = "get">
	<div class="dropdown_a">
	<b>Select Stack</b>
	<select name = "action" size = "1" id ="action-choice">
		<option value ="sel">Create your own stack</option>
		<option value ="mean">MEAN Stack</option>
		<option value ="mern">MERN Stack</option>
		<option value ="mevn">MEVN Stack</option>
		<option value ="lamp">LAMP Stack</option>
		<option value ="...">...</option>
	</select> 
	</div>
	<div class = "button_download">
		<center>
		<button>Template Builder</button>
		</center>
	</div>
	<div class="dropdown_b">
	<b>Backend</b>
	<select name = "action_b" size = "1" id ="action-choice">
		<option value ="node">node.js</option>
		<option value ="python">Python</option>
		<option value ="ruby">Ruby</option>
		<option value ="django">Django</option>
		<option value ="firebase">Firebase</option>
		<option value ="...">...</option> 
	</select>
	</div>
	<div class="dropdown_c">
		<b>Frontend</b>
		<select name = "action_c" size = "1" id ="action-choice">
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
	</div>
	<div class="dropdown_d">
	<b>Database</b>
	<select name = "action_d" size = "1" id ="action-choice">
		<option value ="mongodb">MongoDB</option>
		<option value ="mysql">MySQL</option>
		<option value ="postgre">PostgreSQL</option>
		<option value ="sqlite">SQLite</option>
		<option value ="...">...</option>
	</select>
	</div>
</form>
<div class = "note">
<center>
<p><b>Note: </b>On choosing an existing stack, their corresponding technologies are predefined and any modifications made to the front-end, back-end and database drop-downs will be disregarded. </p>
</center>
</div>
</body>
</html>`

// func handler(w http.ResponseWriter, r *http.Request) {
// 	template.Must(template.ParseFiles("index.html")).Execute(w, nil)
// }

func m(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(page))
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Frontend:", r.FormValue("action_c"), "Backend:", r.FormValue("action_b"), "Database:", r.FormValue("action_d"))
	if r.Method == "GET" {
		r.ParseForm()
		// w.Write([]byte(r.Form["action"][0]))
		if r.Form["action"][0] == "sel" {
			if r.Form["action_c"][0] == "angular" && r.Form["action_b"][0] == "node" && r.Form["action_d"][0] == "mysql" {
				// intrinsic MEAN with MySQL instead of MongoDB
				http.Redirect(w, r, fmt.Sprintf("https://github.com/toga-dagger/aerothon-4.0/tree/main/templates/mean_SQL"), 302)
			} else if r.Form["action_c"][0] == "angular" && r.Form["action_b"][0] == "node" && r.Form["action_d"][0] == "mongodb" {
				// intrinsic MEAN (sel + other parameters)
				http.Redirect(w, r, fmt.Sprintf("https://github.com/toga-dagger/aerothon-4.0/tree/main/templates/mean"), 302)
			} else if r.Form["action_c"][0] == "bootstrap" && r.Form["action_b"][0] == "ruby" && r.Form["action_d"][0] == "sqlite" {
				// bootstrap + ruby + SQLite
				http.Redirect(w, r, fmt.Sprintf("https://github.com/toga-dagger/aerothon-4.0/tree/main/templates/bootstrap_ruby"), 302)
				// } else if r.Form["action_c"][0] == "django" && r.Form["action_b"][0] == "reactn" && r.Form["action_d"][0] == "mongoDB" {
				// 	// Cross - Platform
				// 	http.Redirect(w, r, fmt.Sprintf(""), 302)
			} else {
				// The stack is not supported
				w.Write([]byte("Sorry, this stack is in development"))
			}
		}
		// fmt.Println(r.Form["action"][0])
		if r.Form["action"][0] == "mean" {
			// Print the github link for the mean stack template
			http.Redirect(w, r, fmt.Sprintf("https://github.com/toga-dagger/aerothon-4.0/tree/main/templates/mean"), 302)
		} else if r.Form["action"][0] == "mern" {
			// Print the github link for the mern stack template
			http.Redirect(w, r, fmt.Sprintf("https://github.com/toga-dagger/aerothon-4.0/tree/main/templates/mern"), 302)
			// http.Redirect(w, r, fmt.Sprintf("https://github.com/amazingandyyy/mern"), 302)
		} else if r.Form["action"][0] == "mevn" {
			// Print the github link for the mevn stack template
			http.Redirect(w, r, fmt.Sprintf("https://github.com/toga-dagger/aerothon-4.0/tree/main/templates/mevn"), 302)
		} else if r.Form["action"][0] == "lamp" {
			// Print the github link for the lamp stack template
			http.Redirect(w, r, fmt.Sprintf("https://github.com/toga-dagger/aerothon-4.0/tree/main/templates/lamp"), 302)
		} else {
			w.Write([]byte(""))
		}
		// In case we have to create our own stack using frontend as Angular, backend as node.js and database as mysql
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
