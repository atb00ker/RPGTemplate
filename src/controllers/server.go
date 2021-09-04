package main

import (
	"net/http"
)

var httpRouter Router = NewMuxRouter()

func main() {
	// RPG Routes
	// rpgCaller := rpg.rpgCaller{Hasura: &rpg.HasuraClient{}}
	// httpRouter.POST(rpg.InsertRpgPath, rpgCaller.InsertRpgAction)
	// Static Files
	muxDispatcher.PathPrefix("/").Handler(http.FileServer(http.Dir("./dist/")))
	// Start Server
	httpRouter.SERVE(":3000")
}
