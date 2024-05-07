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

func IsInArray(val int, array []int) bool {
  for _, v := range array {
    if v == val {
      return true
    }
  }
  return false
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
    stays := make([]bool, 5)
    for _, v := range r.Form["stays"] {
        if i, err := strconv.Atoi(v); err == nil && i >= 0 && i < 5 {
            stays[i] = true
        }
    }
    user.Stays = stays
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

  defaultBringsWith := []string{"alcohol", "food", "partner", "friend", "dog", "dye"}
  filteredBringsWith := []string{}
  for _, defaultItem := range defaultBringsWith {
      found := false
      for _, userItem := range user.BringsWith {
          if defaultItem == userItem {
              found = true
              break
          }
      }
      if !found {
          filteredBringsWith = append(filteredBringsWith, defaultItem)
      }
  }

  users, err := models.GetAllUsers()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  filteredUsers := []models.User{}
  for _, u := range users {
      hasTrue := false
      for _, stay := range u.Stays {
          if stay {
              hasTrue = true
              break
          }
      }
      if hasTrue {
          filteredUsers = append(filteredUsers, u)
      }
  }

  data := map[string]interface{}{
    "User":              user,
    "UserJson":          userJson,
    "DefaultBringsWith": filteredBringsWith,
    "Users":             filteredUsers,
  }
  tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/user.html"))
  tmpl.Execute(w, data)
}
