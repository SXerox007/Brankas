package router

import (
	"github.com/gorilla/mux"
)

var HeadNodeRouter *mux.Router

/**
*
*  initilize teh router
*
**/
func InitRouter() {
	HeadNodeRouter = mux.NewRouter()
}

/**
*
* create the subRouter
*
**/
func SubRouter(subRouterPath string) *mux.Router {
	return HeadNodeRouter.PathPrefix(subRouterPath).Subrouter()
}
