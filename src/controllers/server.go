package main

import (
	"net/http"

	"rpgtemplate/src/controllers/config"
	"rpgtemplate/src/controllers/rpg"
)

var httpRouter Router = NewMuxRouter()

func main() {
	config.LoadEnv()
	config.ConnectToDb()
	// RPG Routes
	httpRouter.POST(rpg.CreateRPGUrl, rpg.CreateRPG)
	httpRouter.PUT(rpg.DetailsRPGUrl, rpg.UpdateRPG)
	httpRouter.DELETE(rpg.DetailsRPGUrl, rpg.DeleteRPG)
	httpRouter.GET(rpg.GetRPGUrl, rpg.GetRPG)
	// Static Files
	muxDispatcher.PathPrefix("/").Handler(http.FileServer(http.Dir("./dist/")))
	// Start Server
	httpRouter.SERVE(":3000")
}
