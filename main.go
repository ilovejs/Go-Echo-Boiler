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

	//todo: turn off in prod
	d := db.New()
	db.AutoMigrate(d)

	//store
	us := store.NewUserStore(d)
	//as := store.NewArticleStore(d)
	ps := store.NewProjectStore(d)

	// handler
	h := handler.NewHandler(ps, us) //ps,us,as
	h.Register(v1)

	r.Logger.Fatal(r.Start("localhost:8585"))
}
