package main

import (
	"fmt"
	"net/http"

	"github.com/xgolis/wac-tg-jv-backend/cmd/wac-tg-jv-backend/app"
)

func main() {
	app := app.NewApp()

	fmt.Printf("[Server] Up and running on %s\n", app.Server.Addr)
	http.ListenAndServe(app.Server.Addr, app.Server.Handler)
}
