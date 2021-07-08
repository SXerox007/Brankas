package main

import (
	"Brankas/base/utils"
	"Brankas/models/common"
	"net/http"
	"os"
)

// middleware
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if token := r.FormValue("auth"); token == "" {
			utils.RespondWithJSON(w, http.StatusUnauthorized, common.CommonResponse{Success: false, Message: "Unauthorized.", Code: http.StatusUnauthorized, Data: nil})
		} else {
			//
			if token == os.Getenv("AUTH") {
				//Call the next handler, which can be another middleware in the chain, or the final handler.
				next.ServeHTTP(w, r)
			} else {
				utils.RespondWithJSON(w, http.StatusUnauthorized, common.CommonResponse{Success: false, Message: "Invalid Authorization token.", Code: http.StatusUnauthorized, Data: nil})
			}
		}
		return
	})
}
