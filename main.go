package main

import (
	"MOBIL_API_2/app"
	"MOBIL_API_2/controllers"
	"MOBIL_API_2/exception"
	"MOBIL_API_2/helpers"
	"MOBIL_API_2/repositories"
	"MOBIL_API_2/services"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func corsAndAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// CORS HEADERS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-API-Key")

		// PREFLIGHT HARUS DIHENTIKAN DI SINI
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// API KEY CHECK (SETELAH OPTIONS)
		if r.Header.Get("X-API-Key") != "RAHASIA" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	db := app.ConnectDB()
	app.MonitorDB(db)
	validate := validator.New()
	mobilRepository := repositories.NewMobilRepository()
	mobilService := services.NewMobilService(mobilRepository, db, validate)
	mobilController := controllers.NewMobilController(mobilService)

	router := httprouter.New()
	router.POST("/api/mobil/", mobilController.Create)
	router.PUT("/api/mobil/:mobilId", mobilController.Update)
	router.DELETE("/api/mobil/:mobilId", mobilController.Delete)
	router.GET("/api/mobil/:mobilId", mobilController.FindById)
	router.GET("/api/mobil", mobilController.FindAll)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    ":3000",
		Handler: corsAndAuth(router),
	}
	err := server.ListenAndServe()
	helpers.PanicIfError(err)

}
