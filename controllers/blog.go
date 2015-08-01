package controllers

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"github.com/gorilla/securecookie"
	con "blogger/config"
	models "blogger/models"
	"fmt"
)

type Page struct {
	Title, name,Tag,Active,Url,View,Mode string
	Results     []Entry
}
type Entry struct {
	Name, Mes, Des string
	Ids            int
}

var Title string
var Abstract string
var Description string
var ID int
var db = con.DBConn()
var templates = template.Must(template.ParseFiles("header.html", "footer.html", "main.html", "about.html", "login.html", "home.html", "create.html", "save.html", "deleted.html", "edit.html", "update.html"))

//session
var cookieHandler = securecookie.New(
  securecookie.GenerateRandomKey(64),
  securecookie.GenerateRandomKey(32))

func getUserName(r *http.Request) (userName string) {
  if cookie, err := r.Cookie("session"); err == nil {
    cookieValue := make(map[string]string)
    if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
      userName = cookieValue["name"]
    }
  }
  return userName
}

func setSession(userName string, w http.ResponseWriter) {
  value := map[string]string{
    "name": userName,
  }
  if encoded, err := cookieHandler.Encode("session", value); err == nil {
    cookie := &http.Cookie{
      Name:  "session",
      Value: encoded,
      Path:  "/",
    }
    http.SetCookie(w, cookie)
  }
}

func clearSession(w http.ResponseWriter) {
  cookie := &http.Cookie{
    Name:   "session",
    Value:  "",
    Path:   "/",
    MaxAge: -1,
  }
  http.SetCookie(w, cookie)
}

func Display(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}
func MainHandler(w http.ResponseWriter, r *http.Request) {
	db.DB()
	db.AutoMigrate(&models.Blog{})
	results := []Entry{}
	tRes := Entry{}
	//  var bloggers models.Blog
	// bloggers = models.Blog{
	// 	Title:         "Android",
	// 	Abstract:      "A Mobile Operating system for smartphones, tablets, PDAs",
	// 	Description:   "Android is a Mobile Operating System based on the Linux kernel and currently developed by Google."}
	// db.Create(&bloggers)
	// bloggers = models.Blog{
	// 	Title:         "Google",
	// 	Abstract:      "Google is an American Multinational Technology Company specializing in Internet-related services and products.",
	// 	Description:   "Google was founded by Larry Page and Sergey Brin while they were Ph.D. students at Stanford University."}
	// db.Create(&bloggers)
	// bloggers = models.Blog{
	// 	Title:         "Facebook",
	// 	Abstract:      "Facebook is an Online Social Networking Service headquartered in Menlo Park, California.",
	// 	Description:   "Its website was launched on February 4, 2004, by Mark Zuckerberg with his Harvard College roommates."}
	// db.Create(&bloggers)
	var golang []models.Blog
	db.Find(&golang)
	for _, k := range golang {
		fmt.Println("Title:", k.Title)
		fmt.Println("Abstract:", k.Abstract)
		fmt.Println("Description:", k.Description)
	}
	rows, _ := db.Model(&models.Blog{}).Select("Title, Abstract,Description,ID").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&Title, &Abstract, &Description, &ID)
		tRes.Name = Title
		tRes.Mes = Abstract
		tRes.Des = Description
		tRes.Ids = ID
		results = append(results, tRes)
	}
		userName := getUserName(r)
			var option string
			var path,visible,modes string
			if userName == "" {
				option = "Login"
				path = "index"
				visible = "hidden"
	 } else {
				option = ""
				path = ""
				modes = "hidden"
	 		}
	Display(w, "home", &Page{Title: "Home",Tag: userName, Active: option, Url: path,View: visible,Mode: modes, Results: results})
}
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	results := []Entry{}
	tRes := Entry{}
	var golang []models.Blog
	db.Find(&golang, id)
	rows, _ := db.Model(&models.Blog{}).Where("id = ?", id).Select("Title, Abstract,Description,ID").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&Title, &Abstract, &Description, &ID)
		tRes.Name = Title
		tRes.Mes = Abstract
		tRes.Des = Description
		tRes.Ids = ID
		results = append(results, tRes)
	}
			userName := getUserName(r)
			var option string
			var path,visible,modes string
			if userName == "" {
				option = "Login"
				path = "index"
				visible = "hidden"
	 } else {
				option = ""
				path = ""
				modes = "hidden"
	 		}

	Display(w, "about", &Page{Title: "About",Tag: userName, Active: option, Url: path,View: visible,Mode: modes,Results: results})
}
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	results := []Entry{}
	tRes := Entry{}
	var golang []models.Blog
	db.Delete(&golang, id)
	rows, _ := db.Model(&models.Blog{}).Select("Title, Abstract,Description,ID").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&Title, &Abstract, &Description, &ID)
		tRes.Name = Title
		tRes.Mes = Abstract
		tRes.Des = Description
		tRes.Ids = ID
		results = append(results, tRes)
		}
				userName := getUserName(r)
				var option string
			var path,visible,modes string
			if userName == "" {
				option = "Login"
				path = "index"
				visible = "hidden"
			http.Redirect(w, r, "/index", 302)

	 } else {
				option = ""
				path = ""
				modes = "hidden"
	 		}

	Display(w, "deleted", &Page{Title: "About", Active: option, Url: path,View: visible,Tag: userName,Mode: modes, Results: results})
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
			userName := getUserName(r)
			var option string
			var path,visible,modes string
			if userName == "" {
				option = "Login"
				path = "index"
				visible = "hidden"
			http.Redirect(w, r, "/index", 302)

	 } else {
				option = ""
				path = ""
				modes = "hidden"
	 		}

	Display(w, "create", &Page{Title: "Create Blog", Active: option, Url: path,View: visible,Mode: modes,Tag: userName})
}
func SaveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("inputTitle")
	abstract := r.FormValue("inputAbstract")
	description := r.FormValue("inputDescription")
	fmt.Println(title)
	fmt.Println(abstract)
	fmt.Println(description)
	db.Exec("insert into blogs(Title,Abstract,Description)values(?,?,?)", title, abstract, description)
	Display(w, "save", &Page{Title: "saved Blog"})
}
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	results := []Entry{}
	tRes := Entry{}
	var golang []models.Blog
	db.Find(&golang, id)
	rows, _ := db.Model(&models.Blog{}).Where("id = ?", id).Select("Title, Abstract,Description,ID").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&Title, &Abstract, &Description, &ID)
		tRes.Name = Title
		tRes.Mes = Abstract
		tRes.Des = Description
		tRes.Ids = ID
		results = append(results, tRes)
	}
				userName := getUserName(r)
				var option string
			var path,visible,modes string
			if userName == "" {
				option = "Login"
				path = "index"
				visible = "hidden"
			http.Redirect(w, r, "/index", 302)

	 } else {
				option = ""
				path = ""
				modes = "hidden"
	 		}

	Display(w, "edit", &Page{Title: "Edit Blog",Tag: userName, Active: option, Url: path,View: visible,Mode: modes, Results: results})
}
func EditHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("inputTitle")
	abstract := r.FormValue("inputAbstract")
	description := r.FormValue("inputDescription")
	id := r.FormValue("inputId")
	fmt.Println(id)
	fmt.Println(title)
	fmt.Println(abstract)
	fmt.Println(description)
	db.Exec("UPDATE blogs SET title=? ,abstract=?, description=? WHERE id IN (?)", title, abstract, description, id)
	Display(w, "update", &Page{Title: "Edit Blog"})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	pass := r.FormValue("password")
	fmt.Println(name)
	fmt.Println(pass)
	redirectTarget := "/index"
	if name != "" && pass != "" {
		setSession(name, w)
		redirectTarget = "/"
	}
	http.Redirect(w, r, redirectTarget, 302)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/", 302)
}

func IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	var visible string
	visible = "hidden"
	Display(w, "main", &Page{Title: "Login",View: visible})

}

func InternalPageHandler(w http.ResponseWriter, r *http.Request) {
	// userName := getUserName(r)
	// if userName != "" {
	//    	Display(w, "home", &Page{Title: "Home",Tag: userName})
	//  } else {
	//    http.Redirect(w, r, "/", 302)
	//  }
}
