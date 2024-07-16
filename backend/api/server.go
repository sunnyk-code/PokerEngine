package api

import (
	"net/http"


	p "backend/pokerLogic"
)



func StartServer() {
	http.HandleFunc("/winning-percentage", ApiKeyMiddleware(WinningPercentageHandler))
	http.ListenAndServe(":8080", nil)
}
