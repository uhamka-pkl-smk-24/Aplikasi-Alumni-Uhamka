package controller

import (
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"

	"github.com/jeypc/go-auth/go-auth/config"
	"github.com/jeypc/go-auth/go-auth/entities"
	"github.com/jeypc/go-auth/go-auth/model"
	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Username string
	Password string
	Role     string
}

var userModel = model.NewUserModel()

func Index(w http.ResponseWriter, r *http.Request) {
	session, err := config.Store.Get(r, config.SESSION_ID)
	if err != nil {
		http.Error(w, "Unable to retrieve session", http.StatusInternalServerError)
		return
	}

	if len(session.Values) == 0 || session.Values["loggedIn"] != true {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	lowonganList, err := userModel.GetAllLowongan()
	if err != nil {
		log.Println("error", err.Error())
		http.Error(w, "Failed to fetch job postings", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"username": session.Values["username"],
		"lowongan": lowonganList,
	}

	log.Println("data", data)
	temp, err := template.ParseFiles("go-auth/views/html/index.html")
	if err != nil {
		log.Println("error", err.Error())
		http.Error(w, "Unable to parse template", http.StatusInternalServerError)
		return
	}
	if err := temp.Execute(w, data); err != nil {
		http.Error(w, "Unable to execute template", http.StatusInternalServerError)
		return
	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("go-auth/views/html/login.html")
		if err != nil {
			http.Error(w, "Unable to parse template", http.StatusInternalServerError)
			return
		}
		if err := temp.Execute(w, nil); err != nil {
			http.Error(w, "Unable to execute template", http.StatusInternalServerError)
			return
		}
	} else if r.Method == http.MethodPost {
		r.ParseForm()

		userInput := &UserInput{
			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}

		var user entities.User
		var message error

		// Search for the user by username
		err := userModel.Where(&user, "username", userInput.Username)
		if err != nil || user.Username == "" {
			message = errors.New("username not found")
			log.Println("Error:", message.Error())
		} else {
			// Verify password
			errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))
			if errPassword != nil {
				message = errors.New("incorrect username or password")
				log.Println("Error:", message.Error())
			}
		}

		// If there is an error, display it on the login page
		if message != nil {
			data := map[string]interface{}{
				"error": message.Error(),
			}
			temp, err := template.ParseFiles("go-auth/views/html/login.html")
			if err != nil {
				http.Error(w, "Unable to parse template", http.StatusInternalServerError)
				return
			}
			if err := temp.Execute(w, data); err != nil {
				http.Error(w, "Unable to execute template", http.StatusInternalServerError)
				return
			}
			return
		} else {
			// If login is successful, create a session
			session, err := config.Store.Get(r, config.SESSION_ID)
			if err != nil {
				http.Error(w, "Unable to retrieve session", http.StatusInternalServerError)
				return
			}

			session.Values["loggedIn"] = true
			session.Values["username"] = user.Username
			session.Values["role"] = user.Role
			session.Values["email"] = user.Email
			session.Values["nim"] = user.Nim
			session.Values["no_telp"] = user.No_telp

			if err := session.Save(r, w); err != nil {
				http.Error(w, "Unable to save session", http.StatusInternalServerError)
				return
			}

			// Redirect user based on their role
			switch user.Role {
			case "admin":
				http.Redirect(w, r, "/admin", http.StatusSeeOther)
			case "admin_lowongan":
				http.Redirect(w, r, "/adminlowongan", http.StatusSeeOther)
			case "mahasiswa":
				http.Redirect(w, r, "/mahasiswa", http.StatusSeeOther)
			default:
				http.Redirect(w, r, "/login", http.StatusSeeOther)
			}
		}
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, err := config.Store.Get(r, config.SESSION_ID)
	if err != nil {
		http.Error(w, "Failed to get session", http.StatusInternalServerError)
		log.Println("Error getting session: ", err)
		return
	}

	// Delete session
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		log.Println("Error saving session: ", err)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
	log.Println("User logged out successfully")
}

func SessionInfoHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)

	role := session.Values["role"]
	loggedIn := session.Values["loggedIn"]

	jsonResponse := map[string]interface{}{
		"role":     role,
		"loggedIn": loggedIn,
	}

	json.NewEncoder(w).Encode(jsonResponse)
}
