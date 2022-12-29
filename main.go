package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var Data = map[string]interface{}{
	"Title":   "Personal Web",
	"IsLogin": true,
}

type dataInput struct{
	// Id int
	ProjectName string
	Description string
	Technologies []string
	// StartDate string
	// EndDate string
	Duration string
	Image string
	


}

var dataInputs = []dataInput{
	{
		// Id: 1
		ProjectName:"Pasar Coding di Indonesia Dinilai Masih Menjanjikan",
		Description: "Deskripsi",
		Technologies:[]string {"nodejs","golang","reactjs","python"},
		// StartDate: "2022-12-11",
		// EndDate: "2022-12-31",
		Duration: "3 Months",
		Image:   "/public/assets/female-portrait.jpg",
		
	},
}


func main() {
	route := mux.NewRouter()

	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/",http.FileServer(http.Dir("./public"))))
	
	route.HandleFunc("/home",home).Methods("GET")
	route.HandleFunc("/home",addMyProject).Methods("POST")
	route.HandleFunc("/projectDetail/{id}",projectDetail).Methods("GET")
	route.HandleFunc("/contactMe",contactMe).Methods("GET")
	route.HandleFunc("/addProject",addProject).Methods("GET")
	route.HandleFunc("/detailProject",detailProject).Methods("GET")
	route.HandleFunc("/delete-Project/{id}", deleteProject).Methods("GET")

	fmt.Println("Server is running on port 5000")
	http.ListenAndServe("localhost:5000", route)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message :" + err.Error()))
		return
	}
	respData := map[string]interface{}{
		// "Data":  Data,
		"dataInputs": dataInputs,
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w,respData)
}

func addMyProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}


    projectName :=  r.PostForm.Get("name")
	// startDate := r.PostForm.Get("start-date")
	// endDate := r.PostForm.Get("end-date")
	// duration := endDate - startDate
	descrition := r.PostForm.Get("description")
	checkbox := r.Form["checkbox"]
	image := r.PostForm.Get("image")
	var newProject = dataInput {
		ProjectName: projectName,
	    Description: descrition,
	    Technologies: checkbox,
        Duration: "3 Months",
	    Image: image,
	
	}

	dataInputs = append(dataInputs, newProject)
	fmt.Println(dataInputs)
	http.Redirect(w,r,"/home",http.StatusMovedPermanently)
}

func projectDetail( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html;charset=utf-8")

	id,_ := strconv.Atoi(mux.Vars(r)["id"])

	var tmpl,err = template.ParseFiles("views/detail-page.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message :" + err.Error()))
		return
	}
	resp := map[string]interface{}{
		"dataInputs" : dataInputs[id],
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w,resp)

}

func deleteProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	fmt.Println(id)

	dataInputs = append(dataInputs[:id], dataInputs[id+1:]...)

	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
}

func contactMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/get-in-touch.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message :" + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w,Data)
}
func addProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/add-project.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message :" + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w,Data)
}
func detailProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/detail-page.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message :" + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w,Data)
}