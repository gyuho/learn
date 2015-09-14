package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"text/template"
	"time"

	"golang.org/x/net/context"

	stdlog "log"

	log "github.com/Sirupsen/logrus"
)

/*
sudo kill $(sudo netstat -tlpn | perl -ne 'my @a = split /[ \/]+/; print "$a[6]\n" if m/:5000/gio');
go run 07_web_server.go 1>>stdout.log 2>>stderr.log;
*/

var (
	port    = ":5000"
	logPath = "web.log"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(new(log.JSONFormatter))

	// https://godoc.org/github.com/Sirupsen/logrus#Level
	// log.SetLevel(log.PanicLevel)
	// log.SetLevel(log.FatalLevel)
	// log.SetLevel(log.ErrorLevel)
	// log.SetLevel(log.WarnLevel)
	// log.SetLevel(log.InfoLevel)
	log.SetLevel(log.DebugLevel)
}

func wrapRouterWithLogrus(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		start := time.Now()
		h.ServeHTTP(w, req)
		took := time.Since(start)

		ip, err := getIP(req)
		if err != nil {
			log.Warnf("getIP error: %v", err)
		}
		log.WithFields(log.Fields{
			"event_type": "router",
			"referrer":   req.Referer(),
			"ua":         req.UserAgent(),
			"method":     req.Method,
			"path":       req.URL.Path,
			"ip":         ip,
			"uuid":       uuid.NewV4(),
		}).Debugf("took %s", took)
	}
}

var logger = stdlog.New(os.Stdout, "[TEST] ", stdlog.Ldate|stdlog.Ltime)

func wrapHandlerFunc0(fn func(w http.ResponseWriter, req *http.Request)) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		fn(w, req)
		logger.Printf("%s %s   |  Took %s", req.Method, req.URL.Path, time.Since(start))
	}
}

func wrapHandlerFunc1(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, req)
		logger.Printf("%s %s   |  Took %s", req.Method, req.URL.Path, time.Since(start))
	}
}

func wrapHandlerFunc2(fn func(w http.ResponseWriter, req *http.Request)) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		start := time.Now()

		fn(w, req) // execute the handler
		took := time.Since(start)

		ip, err := getIP(req)
		if err != nil {
			log.Warnf("getIP error: %v", err)
		}
		log.WithFields(log.Fields{
			"event_type": runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(),
			"referrer":   req.Referer(),
			"ua":         req.UserAgent(),
			"method":     req.Method,
			"path":       req.URL.Path,
			"ip":         ip,
			"uuid":       uuid.NewV4(),
		}).Debugf("took %s", took)
	}
}

type ContextHandler interface {
	ServeHTTPContext(context.Context, http.ResponseWriter, *http.Request) error
}

// ContextHandlerFunc wraps func(context.Context, ResponseWriter, *Request)
type ContextHandlerFunc func(context.Context, http.ResponseWriter, *http.Request) error

func (f ContextHandlerFunc) ServeHTTPContext(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	return f(ctx, w, req)
}

// ContextAdapter satisfies:
//	1. interface 'ContextHandler'
//	2. interface 'http.Handler'
type ContextAdapter struct {
	ctx     context.Context
	handler ContextHandler // to wrap func(context.Context, ResponseWriter, *Request) error
}

func (ca *ContextAdapter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if err := ca.handler.ServeHTTPContext(ca.ctx, w, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		ip, err := getIP(req)
		if err != nil {
			log.Warnf("getIP error: %v", err)
		}
		log.WithFields(log.Fields{
			"referrer": req.Referer(),
			"ua":       req.UserAgent(),
			"method":   req.Method,
			"path":     req.URL.Path,
			"ip":       ip,
			"error":    err.Error(),
		}).Errorln("ServeHTTP error")

		// json.NewEncoder(w).Encode(map[string]string{
		// 	"status":  "error",
		// 	"message": err.Error(),
		// })
	}
}

func main() {
	lf, err := openToAppend(logPath)
	if err != nil {
		log.Panic(err)
	}
	defer lf.Close()
	log.SetOutput(lf)

	rootContext := context.Background()
	// rootContext = context.WithValue(rootContext, dbKey, db)

	mainRouter := http.NewServeMux()
	mainRouter.Handle("/", &ContextAdapter{
		ctx:     rootContext,
		handler: ContextHandlerFunc(handlerRoot),
	})
	mainRouter.Handle("/sendjson", &ContextAdapter{
		ctx:     rootContext,
		handler: ContextHandlerFunc(handlerSendJSON),
	})
	mainRouter.Handle("/sendgob", &ContextAdapter{
		ctx:     rootContext,
		handler: ContextHandlerFunc(handlerSendGOB),
	})
	mainRouter.Handle("/hello", &ContextAdapter{
		ctx:     rootContext,
		handler: ContextHandlerFunc(handlerHello),
	})
	mainRouter.Handle("/photo", &ContextAdapter{
		ctx:     rootContext,
		handler: ContextHandlerFunc(handlerPhoto),
	})
	mainRouter.Handle("/json", &ContextAdapter{
		ctx:     rootContext,
		handler: ContextHandlerFunc(handlerJSON),
	})
	mainRouter.Handle("/gob", &ContextAdapter{
		ctx:     rootContext,
		handler: ContextHandlerFunc(handlerGOB),
	})

	// func (mux *ServeMux) Handle(pattern string, handler Handler)
	mainRouter.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Serving http://localhost" + port)
	log.Infoln("Server started at", port)
	if err := http.ListenAndServe(port, wrapRouterWithLogrus(mainRouter)); err != nil {
		log.Panic(err)
	}
}

func handlerRoot(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case "GET":
		fmt.Fprintf(w, "Hello World!")
		return nil

	default:
		http.Error(w, "Method Not Allowed", 405)
		return fmt.Errorf("Method Not Allowed")
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
		return fmt.Errorf("Method Not Allowed")
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
		return fmt.Errorf("Method Not Allowed")
	}
}

func sendRequest(ctx context.Context, port, endPoint string) {
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
				return
			}
			log.WithFields(log.Fields{
				"port":     port,
				"endPoint": endPoint,
			}).Infoln("response:", string(b))

		case "/json":
			data := Data{}
			for {
				if err := json.NewDecoder(resp.Body).Decode(&data); err == io.EOF {
					break
				} else if err != nil {
					log.Warnln(err)
					return
				}
			}
			log.WithFields(log.Fields{
				"port":     port,
				"endPoint": endPoint,
			}).Infof("response: %+v", data)

		case "/gob":
			data := Data{}
			for {
				if err := gob.NewDecoder(resp.Body).Decode(&data); err == io.EOF {
					break
				} else if err != nil {
					log.Warnln(err)
					return
				}
			}
			log.WithFields(log.Fields{
				"port":     port,
				"endPoint": endPoint,
			}).Infof("response: %+v", data)
		}
	}()
	select {
	case <-ctx.Done():
		// Done channel is closed when the deadline expires(times out)
		log.Println("sendRequest timed out!")
		return
	}
}

func handlerHello(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case "GET":
		type page struct {
			Title     string
			Email     string
			Name      string
			Picture   string
			Authority string
		}
		if err := helloTemplate.ExecuteTemplate(w, "base", &page{
			Title:     "Test Title",
			Email:     "gyuhox@gmail.com",
			Name:      "Gyu-Ho Lee",
			Picture:   "/static/img/gopher.png",
			Authority: "Owner",
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
		// panic(helloTemplate.ExecuteTemplate(w, "base", &page{
		// 	Title:     "Test Title",
		// 	Email:     "gyuhox@gmail.com",
		// 	Name:      "Gyu-Ho Lee",
		// 	Picture:   "/static/img/gopher.png",
		// 	Authority: "Owner",
		// }))

	default:
		http.Error(w, "Method Not Allowed", 405)
		return fmt.Errorf("Method Not Allowed")
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
		return fmt.Errorf("Method Not Allowed")
	}
}

type Data struct {
	Name  string
	Value float64
	TS    string
}

func handlerJSON(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case "GET":
		data := Data{}
		data.Name = "Go"
		data.Value = 1000
		data.TS = time.Now().String()[:19]
		if err := json.NewEncoder(w).Encode(data); err != nil {
			return err
		}
		return nil

	default:
		http.Error(w, "Method Not Allowed", 405)
		return fmt.Errorf("Method Not Allowed")
	}
}

func handlerGOB(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case "GET":
		data := Data{}
		data.Name = "Go"
		data.Value = 2000
		data.TS = time.Now().String()[:19]
		if err := gob.NewEncoder(w).Encode(data); err != nil {
			return err
		}
		return nil

	default:
		http.Error(w, "Method Not Allowed", 405)
		return fmt.Errorf("Method Not Allowed")
	}
}

func openToAppend(fpath string) (*os.File, error) {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		f, err = os.Create(fpath)
		if err != nil {
			return f, err
		}
	}
	return f, nil
}

// getIP extracts the user IP address from req, if present.
// https://blog.golang.org/context/userip/userip.go
func getIP(req *http.Request) (net.IP, error) {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
	}
	userIP := net.ParseIP(ip)
	if userIP == nil {
		return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
	}
	return userIP, nil
}

var (
	helloTemplate = template.Must(template.ParseFiles("templates/base.html", "templates/hello.html"))
	photoHTML     = `
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
</body>

</html>		
`
)
