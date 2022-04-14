package router

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/GoodUncleFood/gu/internal/api/control/handler"
	"github.com/GoodUncleFood/gu/internal/api/control/handler/fs"
	"github.com/GoodUncleFood/gu/internal/api/control/handler/page/account"
	"github.com/GoodUncleFood/gu/internal/api/control/handler/page/circuit"
	"github.com/GoodUncleFood/gu/internal/api/control/handler/page/config"
	"github.com/GoodUncleFood/gu/internal/api/control/handler/page/data"
	"github.com/GoodUncleFood/gu/internal/api/control/handler/page/home"
	"github.com/GoodUncleFood/gu/internal/api/control/handler/page/login"
	"github.com/GoodUncleFood/gu/internal/api/control/handler/page/manifest"
	"github.com/GoodUncleFood/gu/internal/api/control/handler/page/markets"
	"github.com/GoodUncleFood/gu/internal/api/control/handler/page/member"
	"github.com/GoodUncleFood/gu/internal/api/control/handler/page/menu"
	"github.com/GoodUncleFood/gu/internal/api/control/handler/page/message"
	"github.com/GoodUncleFood/gu/internal/api/control/handler/page/operate"
	"github.com/GoodUncleFood/gu/internal/api/control/handler/page/order"
	"github.com/GoodUncleFood/gu/internal/api/control/handler/page/pars"
	"github.com/GoodUncleFood/gu/internal/api/control/handler/page/points"
	"github.com/GoodUncleFood/gu/internal/api/control/handler/page/product"
	"github.com/GoodUncleFood/gu/internal/api/control/handler/page/promocodes"
	"github.com/GoodUncleFood/gu/internal/api/control/handler/page/run"
	"github.com/GoodUncleFood/gu/internal/api/control/handler/page/tag"
	"github.com/GoodUncleFood/gu/internal/api/control/handler/page/user"
	"github.com/GoodUncleFood/gu/internal/api/control/handler/status"

	"github.com/gorilla/mux"
)

const (
	pathMarkets = "/markets/"
	pathStatic  = "/static/"
)

type Params struct {
	Templates *template.Template
	StaticDir string
}

func (p *Params) Validate() error {
	if p.Templates == nil {
		return fmt.Errorf("missing templates")
	}
	if p.StaticDir == "" {
		return fmt.Errorf("missing static directory")
	}

	return nil
}

func New(p Params) (*mux.Router, error) {
	if err := p.Validate(); err != nil {
		return nil, fmt.Errorf("invalid params: %v", err)
	}

	// Create the HTTP multiplexer, then attach the handlers.
	router := mux.NewRouter()
	router.StrictSlash(true)

	// Add the fs handler.
	fsHandler := fs.New(fs.Params{
		Dir:    p.StaticDir,
		Prefix: pathStatic,
	})
	router.PathPrefix(pathStatic).Handler(&handler.Handler{
		Auth:      false,
		Handle:    fsHandler,
		Templates: p.Templates,
	}).Methods(http.MethodGet)

	// Attach the login handlers.
	if err := login.Route(login.Params{
		Router:    router,
		Templates: p.Templates,
	}); err != nil {
		return nil, fmt.Errorf(
			"failed to set login routes: %v", err,
		)
	}

	// Add the home handler.
	router.Path("/").Handler(&handler.Handler{
		Auth:      true,
		Handle:    home.Home,
		Templates: p.Templates,
	}).Methods(http.MethodGet)

	// Add the account handlers.
	if err := account.Route(account.Params{
		Router:    router,
		Store:     p.Store,
		Templates: p.Templates,
	}); err != nil {
		return nil, fmt.Errorf(
			"failed to set account routes: %v", err,
		)
	}

	// Attach the circuit handlers.
	if err := circuit.Route(circuit.Params{
		Router:    router,
		Store:     p.Store,
		Templates: p.Templates,
	}); err != nil {
		return nil, fmt.Errorf(
			"failed to set circuit routes: %v", err,
		)
	}

	// Attach the config handlers.
	if err := config.Route(config.Params{
		Router:    router,
		Store:     p.Store,
		Templates: p.Templates,
	}); err != nil {
		return nil, fmt.Errorf(
			"failed to set config routes: %v", err,
		)
	}

	// Attach the data handlers.
	if err := data.Route(data.Params{
		Router:    router,
		Store:     p.Store,
		Templates: p.Templates,
	}); err != nil {
		return nil, fmt.Errorf(
			"failed to set member routes: %v", err,
		)
	}

	// Attach the manifest handlers.
	if err := manifest.Route(manifest.Params{
		Router:    router,
		Store:     p.Store,
		Templates: p.Templates,
	}); err != nil {
		return nil, fmt.Errorf(
			"failed to set manifest routes: %v", err,
		)
	}

	// Attach the market handlers.
	if err := markets.Route(markets.Params{
		Router:    router,
		Store:     p.Store,
		Templates: p.Templates,
	}); err != nil {
		return nil, fmt.Errorf(
			"failed to set market routes: %v", err,
		)
	}

	// Attach the member handlers.
	if err := member.Route(member.Params{
		Router:    router,
		Store:     p.Store,
		Templates: p.Templates,
	}); err != nil {
		return nil, fmt.Errorf(
			"failed to set member routes: %v", err,
		)
	}

	// Attach the menu handlers.
	if err := menu.Route(menu.Params{
		Router:    router,
		Store:     p.Store,
		Templates: p.Templates,
	}); err != nil {
		return nil, fmt.Errorf(
			"failed to set menu routes: %v", err,
		)
	}

	// Attach the message handlers.
	if err := message.Route(message.Params{
		Router:    router,
		Store:     p.Store,
		Templates: p.Templates,
	}); err != nil {
		return nil, fmt.Errorf(
			"failed to set message routes: %v", err,
		)
	}

	// Attach the operate handlers.
	if err := operate.Route(operate.Params{
		Router:    router,
		Store:     p.Store,
		Templates: p.Templates,
	}); err != nil {
		return nil, fmt.Errorf(
			"failed to set operate routes: %v", err,
		)
	}

	// Attach the order handlers.
	if err := order.Route(order.Params{
		Router:    router,
		Store:     p.Store,
		Templates: p.Templates,
	}); err != nil {
		return nil, fmt.Errorf(
			"failed to set order routes: %v", err,
		)
	}

	// Attach the promocode handlers.
	if err := promocodes.Route(promocodes.Params{
		Router:    router,
		Store:     p.Store,
		Templates: p.Templates,
	}); err != nil {
		return nil, fmt.Errorf(
			"failed to set promocode routes: %v", err,
		)
	}

	// Attach the points handlers.
	if err := points.Route(points.Params{
		Router:    router,
		Store:     p.Store,
		Templates: p.Templates,
	}); err != nil {
		return nil, fmt.Errorf(
			"failed to set point routes: %v", err,
		)
	}

	// Attach the par value handlers.
	if err := pars.Route(pars.Params{
		Router:    router,
		Store:     p.Store,
		Templates: p.Templates,
	}); err != nil {
		return nil, fmt.Errorf(
			"failed to set par value routes: %v", err,
		)
	}

	// Attach the products handlers.
	if err := product.Route(product.Params{
		Router:    router,
		Store:     p.Store,
		Templates: p.Templates,
	}); err != nil {
		return nil, fmt.Errorf(
			"failed to set products routes: %v", err,
		)
	}

	// Attach the run handlers.
	if err := run.Route(run.Params{
		Router:    router,
		Store:     p.Store,
		Templates: p.Templates,
	}); err != nil {
		return nil, fmt.Errorf(
			"failed to set run routes: %v", err,
		)
	}

	// Add the status handler.
	router.Path("/status/").Handler(&handler.Handler{
		Auth:      false,
		Handle:    status.Handle,
		Store:     p.Store,
		Templates: p.Templates,
	}).Methods(http.MethodGet)

	// Attach the tags handlers.
	if err := tag.Route(tag.Params{
		Router:    router,
		Store:     p.Store,
		Templates: p.Templates,
	}); err != nil {
		return nil, fmt.Errorf(
			"failed to set tag routes: %v", err,
		)
	}

	// Attach the user handlers.
	if err := user.Route(user.Params{
		Router:    router,
		Store:     p.Store,
		Templates: p.Templates,
	}); err != nil {
		return nil, fmt.Errorf(
			"failed to set user routes: %v", err,
		)
	}

	// Add the catch-all handler.
	router.PathPrefix("/").Handler(&handler.Handler{
		Auth:      true,
		Handle:    home.Missing,
		Store:     p.Store,
		Templates: p.Templates,
	}).Methods(http.MethodGet)

	return router, nil
}
