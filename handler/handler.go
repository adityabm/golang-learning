package handler

import (
	"firstweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views","index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Woops! Something Wrong With This Site", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Woops! Something Wrong With This Site", http.StatusInternalServerError)
		return
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Beda routing ini post"))
}

func SiswaHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNum, err := strconv.Atoi(id)

	if err != nil || idNum < 1 {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views","siswa.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Woops! Something Wrong With This Site", http.StatusInternalServerError)
		return
	}

	data := []entity.Siswa{
		{ID: 1, Name: "Asep", Class: "XII IPA A"},
		{ID: 2, Name: "Adit", Class: "XII IPA B"},
		{ID: 3, Name: "Anton", Class: "XII IPS A"},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Woops! Something Wrong With This Site", http.StatusInternalServerError)
		return
	}
}

func PostGet(w http.ResponseWriter, r *http.Request)  {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("Ini adalah GET"))
	case "POST":
		w.Write([]byte("Ini adalah POST"))
	default:
		http.Error(w, "Woops! Something Wrong With This Site", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views","form.html"), path.Join("views", "layout.html"))

		if err != nil {
			log.Println(err)
			http.Error(w, "Woops! Something Wrong With This Site", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "Woops! Something Wrong With This Site", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "Woops! Something Wrong With This Site", http.StatusBadRequest)
}

func Submit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Woops! Something Wrong With This Site", http.StatusInternalServerError)
			return
		}

		data := map[string]interface{}{
			"Name" : r.Form.Get("name"),
			"Class" : r.Form.Get("class"),
		}

		tmpl, err := template.ParseFiles(path.Join("views","result.html"), path.Join("views", "layout.html"))

		if err != nil {
			log.Println(err)
			http.Error(w, "Woops! Something Wrong With This Site", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Woops! Something Wrong With This Site", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "Woops! Something Wrong With This Site", http.StatusBadRequest)
}