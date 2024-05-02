package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"path"
	"strconv"
	"text/template"

	"30.janschill.de/main/models"
)

type IndexPageData struct {
  Users []models.User
}

// func uniqueLinkHandler(w http.ResponseWriter, r *http.Request) {
//   uniqueLink := getUniqueLinkFromRequest(r)
//   user := getUserFromDatabase(uniqueLink)
//   setSessionCookie(w, user)
//   tmpl := template.Must(template.ParseFiles("templates/unique_link.html"))
//   tmpl.Execute(w, user)
// }

// func formHandler(w http.ResponseWriter, r *http.Request) {
//   user := getUserFromSessionCookie(r)
//   updateUserDataFromForm(user, r.Form)
//   saveUserToDatabase(user)
// }

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  users, err := models.GetAllUsers()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  // for _, v := range users {
  //   println(v.Name)
  // }

  data := IndexPageData{
    Users: users,
  }

  tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/index.html"))
  tmpl.Execute(w, data)
}

func convertToIntSlice(strSlice []string) []int {
  intSlice := make([]int, len(strSlice))
  for i, str := range strSlice {
    intVal, _ := strconv.Atoi(str)
    intSlice[i] = intVal
  }
  return intSlice
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
  url := path.Base(r.URL.Path)
  user, err := models.GetUserByUrl(url)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if r.Method == http.MethodPost {
    r.ParseForm()

    // Update the user's data
    user.Name = r.FormValue("name")
    user.ComesBy = r.FormValue("comesby")
    user.BringsWith = r.Form["bringswith"]
    user.Stays = convertToIntSlice(r.Form["stays"])
    user.Emoji = r.FormValue("emoji")

    user, err = models.UpdateUser(user)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
  }

  cookie, err := r.Cookie("user")
  if err != nil {
    if errors.Is(err, http.ErrNoCookie) {
      cookie = &http.Cookie{
        Name:  "user",
        Value: user.URL,
      }
      http.SetCookie(w, cookie)
    } else {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
  }

  userJson, err := json.Marshal(user)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  data := map[string]interface{}{
    "User":     user,
    "UserJson": userJson, // Use template.JS to avoid escaping in the template
}
  tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/user.html"))
  tmpl.Execute(w, data)
}
