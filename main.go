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

	//store
	us := store.NewUserLoginStore(d)
	ps := store.NewProjectStore(d)

	// handler
	h := handler.NewHandler(ps, us) //ps,us,as
	h.Register(v1)

	r.Logger.Fatal(r.Start("localhost:8585"))
}
