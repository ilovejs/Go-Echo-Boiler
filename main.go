package main

import (
	"onsite/db"
	"onsite/router"
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
	as := store.NewArticleStore(d)
	// handler
	h := handler.NewHandler(us, as)

	h.Register(v1)
	r.Logger.Fatal(r.Start("127.0.0.1:8585"))
}
