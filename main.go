package main

import (
	"onsite/db"
	"onsite/handler"
	"onsite/router"
	"onsite/store"
)

func main() {
	r := router.New()
	// not to mix with page router
	v1 := r.Group("/api")

	//init db and logger
	d := db.New()

	//store DI
	ps := store.NewProjectStore(d)
	us := store.NewUserStore(d)
	bts := store.NewTradeCategoryStore(d)
	cs := store.NewClaimStore(d)

	// NewHandler type and DI constructor is defined in handler/url.go
	h := handler.NewHandler(ps, us, bts, cs)
	h.Register(v1)

	r.Logger.Fatal(r.Start("localhost:8585"))
}
