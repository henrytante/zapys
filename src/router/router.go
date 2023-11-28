package router

import (
	"fmt"
	"log"
	"net/http"
	"zapys/admin"
	"zapys/consultas/cep"
	cpfv1 "zapys/consultas/cpfs/CPFV1"
	cpfv2 "zapys/consultas/cpfs/CPFV2"
	"zapys/consultas/nome/nomev1"
	"zapys/src/config"
	"zapys/src/middleware"

	"github.com/gorilla/mux"
)

func INITSERVER() {
	// mux
	r := mux.NewRouter()

	// rotas
	r.HandleFunc("/api/v1/cpf", middleware.TokenMiddleware(cpfv1.CPFV1)).Methods(http.MethodGet)
	r.HandleFunc("/api/v2/cpf", middleware.TokenMiddleware(cpfv2.CPFV2)).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/nome", middleware.TokenMiddleware(nomev1.NOMEV1)).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/cep", middleware.TokenMiddleware(cep.CEP)).Methods(http.MethodGet)
	r.HandleFunc("/admin/criar", middleware.ADMMIDDLEWARE(admin.CriarUser)).Methods(http.MethodGet)
	r.HandleFunc("/admin/users", middleware.ADMMIDDLEWARE(admin.Users)).Methods(http.MethodGet)
	r.HandleFunc("/admin/delete", middleware.ADMMIDDLEWARE(admin.DeleteUser)).Methods(http.MethodGet)
	

	// configurar roteamento
	PORT := config.GetVarEnv("PORT")
	fmt.Printf("Servidor rodando http://localhost:%s", PORT)
	if err := http.ListenAndServe(":"+PORT, r); err != nil {
		log.Fatal(err)
	}
}
