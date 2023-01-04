package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var Data = map[string]interface{}{
	"Title":   "Personal Web",
	"IsLogin": true,
}

type dataInput struct{
	Id int
	ProjectName string
	Description string
	Technologies []string
	StartDate string
	EndDate string
	Duration string
	Image string
	


}

var dataInputs = []dataInput{
	{
	
		ProjectName:"Pasar Coding di Indonesia Dinilai Masih Menjanjikan",
		Description: "Deskripsi",
		Technologies:[]string {"nodejs","golang","reactjs","python"},
		StartDate: "2022-12-11",
		EndDate: "2022-12-31",
		Duration: "3 Bulan",
		Image:   "/public/assets/female-portrait.jpg",
		
	},
}


func main() {
	route := mux.NewRouter()

	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/",http.FileServer(http.Dir("./public"))))
	
	route.HandleFunc("/home",home).Methods("GET")
	route.HandleFunc("/home",addMyProject).Methods("POST")
	route.HandleFunc("/editProject/{id}",editProject).Methods("GET")
	route.HandleFunc("/updateProject/{id}",updateProject).Methods("POST")
	route.HandleFunc("/projectDetail/{id}",projectDetail).Methods("GET")
	route.HandleFunc("/contactMe",contactMe).Methods("GET")
	route.HandleFunc("/addProject",addProject).Methods("GET")
	route.HandleFunc("/delete-Project/{id}", deleteProject).Methods("GET")

	fmt.Println("Server is running on port 5002")
	http.ListenAndServe("localhost:5002", route)
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
	startDate := r.PostForm.Get("start-date")
	endDate := r.PostForm.Get("end-date")
	descrition := r.PostForm.Get("description")
	checkbox := r.Form["checkbox"]
	image := r.PostForm.Get("image")


	startDateTime,_ := time.Parse("2006-01-02",startDate)

	endDateTime,_ := time.Parse("2006-01-02",endDate)


	distance := endDateTime.Sub(startDateTime)

	var duration string
	year := int(distance.Hours()/(12 * 30 * 24))
	 if year != 0 {
		duration = strconv.Itoa(year) + " tahun"
	}else{
		month := int(distance.Hours()/(30 * 24))
		if month != 0 {
			duration = strconv.Itoa(month) + " bulan"
		}else{
			week := int(distance.Hours()/(7 *24))
			if week != 0 {
				duration = strconv.Itoa(week) +  " minggu"
			} else {
				day := int(distance.Hours()/(24))
				if day != 0 {
					duration = strconv.Itoa(day) + " hari"
				}
			}
		}
	}

	var newProject = dataInput {		
		ProjectName: projectName,
	    Description: descrition,
	    Technologies: checkbox,
		StartDate:startDate,
        EndDate:endDate,
        Duration: duration,
	    Image: image,
	
	}

	dataInputs = append(dataInputs, newProject)
	fmt.Println(dataInputs)
	http.Redirect(w,r,"/home",http.StatusMovedPermanently)
}
func updateProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}  
    projectName :=  r.PostForm.Get("name")
	startDate := r.PostForm.Get("start-date")
	endDate := r.PostForm.Get("end-date")
	description := r.PostForm.Get("description")
	checkbox := r.Form["checkbox"]
	image := r.PostForm.Get("image")
	startDateTime,_ := time.Parse("2006-01-02",startDate)
	endDateTime,_ := time.Parse("2006-01-02",endDate)
	distance := endDateTime.Sub(startDateTime)
	var duration string
	year := int(distance.Hours()/(12 * 30 * 24))
	 if year != 0 {
		duration = strconv.Itoa(year) + " tahun"
	}else{
		month := int(distance.Hours()/(30 * 24))
		if month != 0 {
			duration = strconv.Itoa(month) + " bulan"
		}else{
			week := int(distance.Hours()/(7 *24))
			if week != 0 {
				duration = strconv.Itoa(week) +  " minggu"
			} else {
				day := int(distance.Hours()/(24))
				if day != 0 {
					duration = strconv.Itoa(day) + " hari"
				}
			}
		}
	}

	var newProject = dataInput {		
		ProjectName: projectName,
	    Description: description,
	    Technologies: checkbox,
		StartDate:startDate,
        EndDate:endDate,
        Duration: duration,
	    Image: image,
	
	}


	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	 dataInputs[id] = newProject
	
	http.Redirect(w,r,"/home",http.StatusMovedPermanently)
}


func editProject( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	
	var tmpl,err = template.ParseFiles("views/update.html")
	id,_ := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message :" + err.Error()))
		return
	}
	
	resp := map[string]interface{}{
		"ID" : id,
		"Data":  Data,
		"dataInputs" : dataInputs[id],
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w,resp)

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

	// fmt.Println(id)

	dataInputs = append(dataInputs[:id], dataInputs[id+1:]...)

	http.Redirect(w, r, "/home", http.StatusFound)
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
