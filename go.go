// func blogDetail(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "text/html; charset=utf-8")

//     id, _ := strconv.Atoi(mux.Vars(r)["id"])

//     var tmpl, err = template.ParseFiles("views/blog-detail.html")
//     if err != nil {
//         w.WriteHeader(http.StatusInternalServerError)
//         w.Write([]byte("message : " + err.Error()))
//         return
//     }

//     BlogDetail := Blog{}
//     err = connection.Conn.QueryRow(context.Background(), "SELECT id, title, image, content, post_date FROM tb_blog WHERE id=$1", id).Scan(
//         &BlogDetail.Id, &BlogDetail.Title, &BlogDetail.Image, &BlogDetail.Content, &BlogDetail.Post_date)
//     if err != nil {
//         w.WriteHeader(http.StatusInternalServerError)
//         w.Write([]byte("message : " + err.Error()))
//         return
//     }

//     BlogDetail.Author = "Ilham Fathullah"
//     BlogDetail.Format_date = BlogDetail.Post_date.Format("2 January 2006")

//     resp := map[string]interface{}{
//         "Data": Data,
//         "Blog": BlogDetail,
//     }

//     w.WriteHeader(http.StatusOK)
//     tmpl.Execute(w, resp)
// }


// func home(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")

// 	var tmpl, err = template.ParseFiles("views/index.html")
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("message :" + err.Error()))
// 		return
// 	}

// 	var result []dataInput

// 	rows, err := connection.Conn.Query(context.Background(), "SELECT id, name, description, technologies, start_date, end_date FROM tb_project ")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return 
// 	}

// 	for rows.Next() {
// 		var each = dataInput{}

// 		var err = rows.Scan(&each.Id, &each.ProjectName,&each.Description,&each.Technologies,&each.start_date, &each.end_date)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			return
// 		}

// 		each.Duration = selisih(each.start_date, each.end_date)
// 		result = append(result, each)
// 	}




// 	respData := map[string]interface{}{
// 		// "Data":  Data,
// 		"dataInputs": result,
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	tmpl.Execute(w,respData)
// }