package main

import (
	"database/sql"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"
	"time"

	"golang.org/x/net/context"

	log "github.com/Sirupsen/logrus"
	auth "github.com/bradrydzewski/go.auth"
)

var (
	mainTemplate = template.Must(template.ParseFiles("templates/base.html", "templates/main.html"))
	photoHTML    = `
<html>
<head>
<title>photo!</title>
<style>
* {
	margin: 0;
	padding: 0;
}
	
html { 
	background: url(/static/img/bg.jpg) no-repeat center center fixed; 
	-webkit-background-size: cover;
	-moz-background-size: cover;
	-o-background-size: cover;
	background-size: cover;
}

img.mainpage {
    display: block;
    margin-left: auto;
    margin-right: auto;
}

#page-wrap { width: 400px; margin: 50px auto; padding: 20px; background: white; -moz-box-shadow: 0 0 20px black; -webkit-box-shadow: 0 0 20px black; box-shadow: 0 0 20px black; }
p {
	font: 15px/2 Georgia, Serif;
	margin: 0 0 30px 0;
	text-indent: 40px;
}
</style>
</head>

<body>
	<br>
	<br>
	<br>
	<p>photos!</p>
	<br>
	<a href="/main">Go to main</a>
</body>

</html>		
`
)

func handlerRoot(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case "GET":
		fmt.Fprintf(w, `<a href="/auth/login"><img class="mainpage" src="https://developers.google.com/+/images/branding/sign-in-buttons/Red-signin_Long_base_44dp.png" style="width:250px;height:50px"></a>`)
		return nil

	default:
		http.Error(w, "Method Not Allowed", 405)
		return fmt.Errorf("Method Not Allowed: %+v", req.Method)
	}
}

func handlerLogout(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case "GET":
		auth.DeleteUserCookie(w, req)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return nil

	default:
		http.Error(w, "Method Not Allowed", 405)
		return fmt.Errorf("Method Not Allowed: %+v", req.Method)
	}
}

func handlerMain(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case "GET":
		user := ctx.Value(userKey).(*auth.User)
		type page struct {
			Link     string
			Email    string
			Name     string
			Picture  string
			Provider string
		}
		if err := mainTemplate.ExecuteTemplate(w, "base", &page{
			Link:     (*user).Link(),
			Email:    (*user).Email(),
			Name:     (*user).Name(),
			Picture:  (*user).Picture(), // "/static/img/gopher.png",
			Provider: (*user).Provider(),
		}); err != nil {
			return err
		}
		return nil

		// THIS IS NOT WORKING
		// No data received ERR_EMPTY_RESPONSE in Chrome
		//
		// Not work because ExecuteTemplate returns nil
		// if it succeeds
		//
		// panic(mainTemplate.ExecuteTemplate(w, "base", &page{
		// 	Title:     "Test Title",
		// 	Email:     "gyuhox@gmail.com",
		// 	Name:      "Gyu-Ho Lee",
		// 	Picture:   "/static/img/gopher.png",
		// 	Authority: "Owner",
		// }))

	default:
		http.Error(w, "Method Not Allowed", 405)
		return fmt.Errorf("Method Not Allowed: %+v", req.Method)
	}
}

func handlerMainPostFormSentence(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case "POST":
		user := ctx.Value(userKey).(*auth.User)
		userID := fmt.Sprintf("%+v", *user)

		if err := req.ParseForm(); err != nil {
			return err
		}
		inputSentence := ""
		if _, ok := req.Form["form_name_sentence"]; !ok {
			return fmt.Errorf("wrong Form: %+v", req.Form)
		} else {
			inputSentence = req.Form["form_name_sentence"][0]
		}

		globalStorage.Lock()
		defer globalStorage.Unlock()

		globalStorage.userIDToData[userID].Sentence = inputSentence
		return nil

	default:
		http.Error(w, "Method Not Allowed", 405)
		return fmt.Errorf("Method Not Allowed: %+v", req.Method)
	}
}

func handlerMainGetFormSentence(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case "GET":
		user := ctx.Value(userKey).(*auth.User)
		userID := fmt.Sprintf("%+v", *user)

		globalStorage.Lock()
		defer globalStorage.Unlock()

		if _, ok := globalStorage.userIDToData[userID]; ok {
			globalStorage.userIDToData[userID].Sentence = strings.TrimSpace(strings.ToUpper(globalStorage.userIDToData[userID].Sentence))
		}

		response := struct {
			Sentence string
		}{
			globalStorage.userIDToData[userID].Sentence,
		}
		return json.NewEncoder(w).Encode(response)

	default:
		http.Error(w, "Method Not Allowed", 405)
		return fmt.Errorf("Method Not Allowed: %+v", req.Method)
	}
}

func handlerMainReset(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case "GET":
		user := ctx.Value(userKey).(*auth.User)
		userID := fmt.Sprintf("%+v", *user)

		globalStorage.Lock()
		defer globalStorage.Unlock()

		globalStorage.userIDToData[userID] = &data{}

		response := struct {
			Sentence string
		}{
			"",
		}
		return json.NewEncoder(w).Encode(response)

	default:
		http.Error(w, "Method Not Allowed", 405)
		return fmt.Errorf("Method Not Allowed: %+v", req.Method)
	}
}

func handlerPhoto(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case "GET":
		if _, err := fmt.Fprint(w, photoHTML); err != nil {
			return err
		}
		return nil

	default:
		http.Error(w, "Method Not Allowed", 405)
		return fmt.Errorf("Method Not Allowed: %+v", req.Method)
	}
}

func handlerSendJSON(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case "GET":
		fmt.Fprintf(w, "handlerSendJSON")
		ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
		defer cancel()
		sendRequest(ctx, port, "/json")
		return nil

	default:
		http.Error(w, "Method Not Allowed", 405)
		return fmt.Errorf("Method Not Allowed: %+v", req.Method)
	}
}

func handlerSendGOB(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case "GET":
		fmt.Fprintf(w, "handlerSendGOB")
		ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
		defer cancel()
		sendRequest(ctx, port, "/gob")
		return nil

	default:
		http.Error(w, "Method Not Allowed", 405)
		return fmt.Errorf("Method Not Allowed: %+v", req.Method)
	}
}

func handlerJSON(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case "GET":
		o := Output{}
		o.Name = "Go"
		o.Value = 1000
		o.TS = time.Now().String()[:19]
		if err := json.NewEncoder(w).Encode(o); err != nil {
			return err
		}
		return nil

	default:
		http.Error(w, "Method Not Allowed", 405)
		return fmt.Errorf("Method Not Allowed: %+v", req.Method)
	}
}

func handlerGOB(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case "GET":
		o := Output{}
		o.Name = "Go"
		o.Value = 2000
		o.TS = time.Now().String()[:19]
		if err := gob.NewEncoder(w).Encode(o); err != nil {
			return err
		}
		return nil

	default:
		http.Error(w, "Method Not Allowed", 405)
		return fmt.Errorf("Method Not Allowed: %+v", req.Method)
	}
}

func sendRequest(ctx context.Context, port, endPoint string) {
	done := make(chan struct{})
	go func() {
		client := http.DefaultClient
		req, err := http.NewRequest("GET", "http://localhost"+port+endPoint, nil)
		if err != nil {
			log.Warnln(err)
			return
		}
		resp, err := client.Do(req)
		if err != nil {
			log.Warnln(err)
			return
		}
		defer resp.Body.Close()

		switch endPoint {
		case "/":
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Warnln(err)
			} else {
				log.WithFields(log.Fields{
					"port":     port,
					"endPoint": endPoint,
				}).Infoln("response:", string(b))
			}
			done <- struct{}{}

		case "/json":
			o := Output{}
			for {
				if err := json.NewDecoder(resp.Body).Decode(&o); err == io.EOF {
					log.WithFields(log.Fields{
						"port":     port,
						"endPoint": endPoint,
					}).Infof("response: %+v", o)
					break

				} else if err != nil {
					log.Warnln(err)
					break

				}
			}
			done <- struct{}{}

		case "/gob":
			o := Output{}
			for {
				if err := gob.NewDecoder(resp.Body).Decode(&o); err == io.EOF {
					log.WithFields(log.Fields{
						"port":     port,
						"endPoint": endPoint,
					}).Infof("response: %+v", o)
					break

				} else if err != nil {
					log.Warnln(err)
					break

				}
			}
			done <- struct{}{}

		}
	}()
	select {

	case <-done:
		log.Println("sendRequest done!")
		return

	case <-ctx.Done():
		log.Println("sendRequest timeout:", ctx.Err())
		return

	}
}

func handlerDB(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case "GET":
		user := ctx.Value(userKey).(*auth.User)
		xd := ctx.Value(xdbKey).(*sql.DB)
		userID := fmt.Sprintf("%+v", *user)

		globalStorage.Lock()
		defer globalStorage.Unlock()

		fmt.Println("xd:", *xd)
		fmt.Println("userID:", userID)

		return nil

	default:
		http.Error(w, "Method Not Allowed", 405)
		return fmt.Errorf("Method Not Allowed: %+v", req.Method)
	}
}
