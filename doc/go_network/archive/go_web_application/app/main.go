package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/context"

	stdlog "log"

	log "github.com/Sirupsen/logrus"
	auth "github.com/bradrydzewski/go.auth"
	graceful "gopkg.in/tylerb/graceful.v1"
)

func main() {
	keepRunning()
}

func keepRunning() {
	f, err := openToAppend(logPath)
	if err != nil {
		stdlog.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)

	globalStorage.Lock()
	// (X) this will deadlock
	// defer globalStorage.Unlock()
	globalStorage.userIDToData = make(map[string]*data)
	globalStorage.Unlock()

	cx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	xx, err := getConn(cx, isInVPC, "xx")
	if err != nil {
		log.Panic(err)
	}
	xdb = xx

	co, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	oo, err := getConn(co, isInVPC, "oo")
	if err != nil {
		log.Panic(err)
	}
	odb = oo

	cr, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	rr, err := getConn(cr, isInVPC, "rr")
	if err != nil {
		log.Panic(err)
	}
	rdb = rr

	cc, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	cd, err := getConn(cc, isInVPC, "cc")
	if err != nil {
		log.Panic(err)
	}
	cdb = cd

	defer odb.Close()
	defer xdb.Close()
	defer rdb.Close()
	defer cdb.Close()

	defer func() {
		if err := recover(); err != nil {

			log.WithFields(log.Fields{
				"event_type": "panic_recover",
				"error":      err,
			}).Errorln("keepRunning error")

			panicCount++
			if panicCount == recoverLimit {
				log.WithFields(log.Fields{
					"event_type": "panic",
					"error":      err,
				}).Panicln("Too much panic:", panicCount)
			}

			keepRunning()
		}
	}()

	mainRun()
}

func mainRun() {
	lf, err := openToAppend(logPath)
	if err != nil {
		log.Panic(err)
	}
	defer lf.Close()
	log.SetOutput(lf)

	googleRedirect := "http://domain.com/auth/login"
	if !isInVPC {
		googleRedirect = fmt.Sprintf("http://localhost%s/auth/login", port)
	}
	auth.Config.CookieSecret = []byte("YOUR_COOKIE_SECRET")
	auth.Config.LoginSuccessRedirect = "/main"
	auth.Config.CookieSecure = false

	rootContext := context.Background()

	mainRouter := http.NewServeMux()
	// func (mux *ServeMux) Handle(pattern string, handler Handler)
	mainRouter.Handle("/static/", http.StripPrefix(
		"/static/",
		http.FileServer(http.Dir("static")),
	))
	mainRouter.Handle("/", &ContextAdapter{
		ctx:     rootContext,
		handler: ContextHandlerFunc(handlerRoot),
	})
	mainRouter.Handle("/auth/login", auth.Google(
		googleClientID,
		googleClientSecret,
		googleRedirect,
	))
	mainRouter.Handle("/auth/logout", &ContextAdapter{
		ctx:     rootContext,
		handler: ContextHandlerFunc(handlerLogout),
	})
	mainRouter.Handle("/main", &ContextAdapter{
		ctx:     rootContext,
		handler: WithAuthentication(ContextHandlerFunc(handlerMain)),
	})
	mainRouter.Handle("/main/post_form_sentence", &ContextAdapter{
		ctx:     rootContext,
		handler: WithAuthentication(ContextHandlerFunc(handlerMainPostFormSentence)),
	})
	mainRouter.Handle("/main/get_form_sentence", &ContextAdapter{
		ctx:     rootContext,
		handler: WithAuthentication(ContextHandlerFunc(handlerMainGetFormSentence)),
	})
	mainRouter.Handle("/main/reset", &ContextAdapter{
		ctx:     rootContext,
		handler: WithAuthentication(ContextHandlerFunc(handlerMainReset)),
	})
	mainRouter.Handle("/photo", &ContextAdapter{
		ctx:     rootContext,
		handler: WithAuthentication(ContextHandlerFunc(handlerPhoto)),
	})
	mainRouter.Handle("/sendjson", &ContextAdapter{
		ctx:     rootContext,
		handler: ContextHandlerFunc(handlerSendJSON),
	})
	mainRouter.Handle("/sendgob", &ContextAdapter{
		ctx:     rootContext,
		handler: ContextHandlerFunc(handlerSendGOB),
	})
	mainRouter.Handle("/json", &ContextAdapter{
		ctx:     rootContext,
		handler: ContextHandlerFunc(handlerJSON),
	})
	mainRouter.Handle("/gob", &ContextAdapter{
		ctx:     rootContext,
		handler: ContextHandlerFunc(handlerGOB),
	})
	mainRouter.Handle("/db", &ContextAdapter{
		ctx:     rootContext,
		handler: WithDatabase(WithAuthentication(ContextHandlerFunc(handlerDB))),
	})

	stdlog.Println("Server started at http://localhost" + port)
	log.Infoln("Server started at http://localhost" + port)
	graceful.Run(port, 10*time.Second, WithLogrus(mainRouter))
}
