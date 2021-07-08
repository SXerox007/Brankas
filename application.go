package main

import (
	db "Brankas/base/db/postgres"
	"Brankas/base/environment"
	"Brankas/base/router"
	"Brankas/base/router/server"
	"log"
)

// init
func Init() {
	env := environment.GetEnv()
	port := environment.GetPort()
	//db.DBConnecting()
	router.InitRouter()
	setupRouter(env, port)
}

func main() {
	pg := db.InitPG()
	defer pg.Close()
	Init()
}

func setupRouter(env, port string) {
	templateMux := router.SubRouter("/brankas")
	templateMux.HandleFunc("/test", TestHandler()).Methods("GET")
	templateMux.HandleFunc("/upload", GetFileUploadPage()).Methods("GET")

	//
	brankasMux := router.SubRouter("/{api}/{version}/brankas")
	brankasMux.Use(AuthMiddleware)
	brankasMux.HandleFunc("/upload", UploadFile()).Methods("POST")

	log.Println("Server serve at", env+":"+port)
	server.StartServer(port)
}
