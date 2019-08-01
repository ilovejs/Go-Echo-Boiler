package main

import (
	"onsite/db"
	"onsite/handler"
	"onsite/router"
	"onsite/store"
)

func main() {
	// see router/settings.go type definition
	r := router.New()

	// not to mix with page router
	v1 := r.Group("/api")

	// init db and logger
	d := db.New()

	// store DI
	ps := store.NewProjectStore(d)
	us := store.NewUserStore(d)
	bts := store.NewTradeCategoryStore(d)
	ts := store.NewTradeStore(d)
	cs := store.NewClaimStore(d)

	// NewHandler type and DI constructor is defined in handler/url.go
	h := handler.NewHandler(ps, us, bts, ts, cs)
	h.Register(v1)

	// don't use 8080 since vue app would occupy that common port number.
	r.Logger.Fatal(r.Start("localhost:8585"))
}
