package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/gplus"
)

/*
sudo kill $(sudo netstat -tlpn | perl -ne 'my @a = split /[ \/]+/; print "$a[6]\n" if m/:8080/gio')
*/

const (
	port              = ":8080"
	gPlusClientID     = "YOUR_KEY"
	gPlusClientSecret = "YOUR_SECRET"

	// sessionName is the key used to access the session store.
	sessionName = "_gothic_session"

	// appKey should be replaced by applications using gothic.
	appKey = "XDZZYmriq8pJ5k8OKqdDuUFym2e7Im5O1MzdyapfotOnrqQ7ZEdTN9AA7K6aPieC"
)

var sessionStore sessions.Store

func init() {
	if sessionStore == nil {
		sessionStore = sessions.NewCookieStore([]byte(appKey))
	}
}

func main() {
	goth.UseProviders(gplus.New(gPlusClientID, gPlusClientSecret, "http://localhost:8080/auth/gplus/callback"))

	mainRouter := http.NewServeMux()
	mainRouter.HandleFunc("/", logInHandler)
	mainRouter.HandleFunc("/auth/gplus", beginAuthHandler)
	mainRouter.HandleFunc("/auth/gplus/callback", callbackHandler)
	mainRouter.HandleFunc("/hello", helloHandler)
	fmt.Println("Serving http://localhost" + port)
	if err := http.ListenAndServe(port, mainRouter); err != nil {
		panic(err)
	}
}

func logInHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		fmt.Fprintf(w, `<p><a href="/auth/gplus">Log in with Google</a></p>`)
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func beginAuthHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		url, err := getAuthURL(w, req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
			return
		}
		http.Redirect(w, req, url, http.StatusTemporaryRedirect)

	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func callbackHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		user, err := completeUserAuth(w, req)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		fmt.Println("user:", user)
		http.Redirect(w, req, "http://localhost:8080/hello", http.StatusSeeOther)

	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		fmt.Fprintf(w, "Hello World!")
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func getAuthURL(w http.ResponseWriter, req *http.Request) (string, error) {
	provider, err := goth.GetProvider("gplus")
	if err != nil {
		return "", err
	}
	sess, err := provider.BeginAuth("state")
	if err != nil {
		return "", err
	}
	url, err := sess.GetAuthURL()
	if err != nil {
		return "", err
	}
	session, _ := sessionStore.Get(req, sessionName)
	session.Values[sessionName] = sess.Marshal()
	if err := session.Save(req, w); err != nil {
		return "", err
	}
	return url, err
}

func completeUserAuth(w http.ResponseWriter, req *http.Request) (goth.User, error) {
	provider, err := goth.GetProvider("gplus")
	if err != nil {
		return goth.User{}, err
	}
	session, _ := sessionStore.Get(req, sessionName)
	if session.Values[sessionName] == nil {
		return goth.User{}, errors.New("could not find a matching session for this request")
	}
	sess, err := provider.UnmarshalSession(session.Values[sessionName].(string))
	if err != nil {
		return goth.User{}, err
	}
	if _, err := sess.Authorize(provider, req.URL.Query()); err != nil {
		return goth.User{}, err
	}
	return provider.FetchUser(sess)
}
