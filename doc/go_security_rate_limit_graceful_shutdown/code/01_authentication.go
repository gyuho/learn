package main

import (
	"fmt"
	"net/http"

	auth "github.com/bradrydzewski/go.auth"
)

const (
	port               = ":8080"
	isProd             = false
	googleClientID     = "YOUR_KEY"
	googleClientSecret = "YOUR_SECRET"
)

func main() {
	googleRedirect := "http://yourrealwebsite.com/auth/login"
	if !isProd {
		googleRedirect = fmt.Sprintf("http://localhost%s/auth/login", port)
	}
	auth.Config.CookieSecret = []byte("YOUR_COOKIE_SECRET")
	auth.Config.LoginSuccessRedirect = "/secret"
	auth.Config.CookieSecure = false

	mainRouter := http.NewServeMux()
	mainRouter.Handle("/auth/login", auth.Google(googleClientID, googleClientSecret, googleRedirect))
	mainRouter.HandleFunc("/", logInHandler)
	mainRouter.HandleFunc("/secret", auth.SecureFunc(secreteHandler))
	mainRouter.HandleFunc("/private", auth.SecureFunc(secreteHandler))
	mainRouter.HandleFunc("/auth/logout", logoutHandler)

	fmt.Println("Serving http://localhost" + port)
	if err := http.ListenAndServe(port, mainRouter); err != nil {
		panic(err)
	}
}

func logInHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		fmt.Fprintf(w, `<a href="/auth/login"><img class="mainpage" src="https://developers.google.com/+/images/branding/sign-in-buttons/Red-signin_Long_base_44dp.png" style="width:250px;height:50px"></a>`)
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func secreteHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		user, err := auth.GetUserCookie(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
			return
		}
		if user.Email() != "gyuho.cs@gmail.com" {
			fmt.Println("only gyuho.cs@gmail.com can access")
			auth.DeleteUserCookie(w, req)
			http.Redirect(w, req, "http://google.com", http.StatusSeeOther)
		}
		fmt.Fprintf(w, `<a href="/auth/logout">logout</a><br>authorized user: %+v`, user)

	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func logoutHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		auth.DeleteUserCookie(w, req)
		http.Redirect(w, req, "/", http.StatusSeeOther)
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}
