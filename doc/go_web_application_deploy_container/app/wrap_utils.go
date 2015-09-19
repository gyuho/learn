package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/net/context"

	log "github.com/Sirupsen/logrus"
	auth "github.com/bradrydzewski/go.auth"
	uuid "github.com/satori/go.uuid"
)

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
		// (X)
		// w.WriteHeader(http.StatusBadRequest)
		// http: multiple response.WriteHeader calls
		// recovered from panic: runtime error: invalid memory address
		// or nil pointer dereference

		ip, err := getIP(req)
		if err != nil {
			log.Warnf("getIP error: %v", err)
		}
		log.WithFields(log.Fields{
			"event_type": "error",
			"referrer":   req.Referer(),
			"ua":         req.UserAgent(),
			"method":     req.Method,
			"path":       req.URL.Path,
			"ip":         ip,
			"error":      err,
		}).Errorln("ServeHTTP error")

		// json.NewEncoder(w).Encode(map[string]string{
		// 	"status":  "error",
		// 	"message": err.Error(),
		// })
	}
}

func WithAuthentication(h ContextHandler) ContextHandler {
	return ContextHandlerFunc(func(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
		user, err := auth.GetUserCookie(req)
		// if no active user session then authorize user
		if err != nil || user.Id() == "" {
			http.Redirect(w, req, auth.Config.LoginRedirect, http.StatusSeeOther)
			log.Warnf("unidentified user: %+v", user)
			return nil
		}
		if _, ok := accessibleEmail[user.Email()]; !ok {
			auth.DeleteUserCookie(w, req)
			http.Redirect(w, req, "http://google.com", http.StatusSeeOther)
			log.Warnf("unidentified user: %+v", user)
			return nil
		}

		ctx = context.WithValue(ctx, userKey, &user)

		userID := fmt.Sprintf("%+v", user)

		globalStorage.Lock()
		// (X) this will deadlock
		// defer globalStorage.Unlock()
		if globalStorage.userIDToData[userID] == nil {
			globalStorage.userIDToData[userID] = &data{}
		}
		globalStorage.Unlock()

		return h.ServeHTTPContext(ctx, w, req)
	})
}

func WithDatabase(h ContextHandler) ContextHandler {
	return ContextHandlerFunc(func(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
		ctx = context.WithValue(ctx, odbKey, odb)
		ctx = context.WithValue(ctx, xdbKey, xdb)
		ctx = context.WithValue(ctx, rdbKey, rdb)
		ctx = context.WithValue(ctx, cdbKey, cdb)
		return h.ServeHTTPContext(ctx, w, req)
	})
}

func WithLogrus(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		ip, err := getIP(req)
		if err != nil {
			log.Warnf("getIP error: %v", err)
		}

		defer func() {
			if err := recover(); err != nil {
				log.WithFields(log.Fields{
					"event_type": "panic_recover",
					"referrer":   req.Referer(),
					"ua":         req.UserAgent(),
					"method":     req.Method,
					"path":       req.URL.Path,
					"ip":         ip,
					"error":      err,
				}).Errorln("WithLogrus error")
			}
		}()

		start := nowPacific()
		h.ServeHTTP(w, req)
		took := time.Since(start)

		log.WithFields(log.Fields{
			"event_type": "web_server",
			"referrer":   req.Referer(),
			"ua":         req.UserAgent(),
			"method":     req.Method,
			"path":       req.URL.Path,
			"ip":         ip,
			"uuid":       uuid.NewV4(),
		}).Debugf("took %s", took)
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

func cleanUp(str string) string {
	s := strings.Fields(strings.TrimSpace(str))
	return strings.Join(s, " ")
}

func nowPacific() time.Time {
	tzone, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		return time.Now()
	}
	return time.Now().In(tzone)
}
