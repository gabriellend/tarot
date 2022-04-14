package handler

import (
	"html/template"
	"net/http"
	"time"

	"github.com/GoodUncleFood/gu/pkg/l"
	"github.com/GoodUncleFood/gu/pkg/models/member"
	"github.com/GoodUncleFood/gu/pkg/page"
	"github.com/gorilla/sessions"
)

type Params struct {
	Templates *template.Template
}

type RequestHandler func(p Params, r *http.Request, w http.ResponseWriter) error

type Handler struct {
	Auth      bool
	Handle    RequestHandler
	Store     *sessions.CookieStore
	Templates *template.Template
}

type Error struct {
	StatusCode int
	Message    string
	Debug      string
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	if !h.Auth {
		ip := r.Header.Get("X-Forwarded-For")
		l.Info(
			"request to %s from %s (%s)",
			r.URL.Path, r.RemoteAddr, ip,
		)

		if err := h.Handle(Params{
			Account:   nil,
			Store:     h.Store,
			Templates: h.Templates,
		}, r, w); err != nil {
			l.Error("%v", err)

			page.Page{
				Debug:          err.Error(),
				DisplayMessage: "Sorry, something went wrong.",
				Time:           time.Now().Format(time.RFC3339),
				Title:          "Good Uncle: Error",
				URL:            r.URL.Path,
				View:           "error",
			}.Render(h.Templates, w)

			return
		}

		l.Info("request to %s took %v", r.URL.Path, time.Since(start))
		return
	}

	// Authenticate before handling.
	session, err := h.Store.Get(r, "gu")
	if err != nil {
		http.Redirect(w, r, "/login/", http.StatusSeeOther)
		return
	}

	auth, _ := session.Values["auth"].(bool)
	if !auth {
		http.Redirect(w, r, "/login/", http.StatusSeeOther)
		return
	}
	account, err := member.Get(session.Values["id"].(string))
	if err != nil {
		http.Redirect(w, r, "/login/", http.StatusSeeOther)
		return
	}
	// Don't allow disabled accounts.
	if !account.Enabled {
		http.Redirect(w, r, "/login/", http.StatusSeeOther)
		return
	}

	l.Info(
		"request to %s by %s (%s)",
		r.URL.Path, account.FullName(), account.Id,
	)

	go func() {
		defer l.Panic()

		if err := account.SetLastRequest(); err != nil {
			l.Error(
				"failed to set %s last request time: %v",
				account.Id, err,
			)
		}
	}()

	if err := h.Handle(Params{
		Account:   account,
		Store:     h.Store,
		Templates: h.Templates,
	}, r, w); err != nil {
		l.Error("%v", err)

		page.Page{
			Account:        account,
			Debug:          err.Error(),
			DisplayMessage: "Sorry, something went wrong.",
			Time:           time.Now().Format(time.RFC3339),
			Title:          "Good Uncle: Error",
			URL:            r.URL.Path,
			View:           "error",
		}.Render(h.Templates, w)
	}

	l.Info("request to %s took %v", r.URL.Path, time.Since(start))
}
