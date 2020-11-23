package router

import (
	"go-sqlite/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	// create a new router
	router := mux.NewRouter()

	// ------------------------------ Mac Adress Record routes ---------------------------------------------

	router.HandleFunc("/createMacAddressRecord", middleware.CreateMacAddressRecord).Methods("POST")

	router.HandleFunc("/getAllMacAddressRecords", middleware.GetAllMacAddressRecords).Methods("GET")

	router.HandleFunc("/getMacAddressRecord/{macID}", middleware.GetMacAddressRecord).Methods("GET")

	router.HandleFunc("/updateMacAddressOfRecord/{macID}/{macAddress}", middleware.UpdateMacAddressOfRecord).Methods("PUT")

	router.HandleFunc("/updateMacGroupOfRecord/{macID}/{macGroup}", middleware.UpdateMacGroupOfRecord).Methods("PUT")

	router.HandleFunc("/deleteMacAddressRecord/{macID}", middleware.DeleteMacAddressRecord).Methods("DELETE")

	router.HandleFunc("/deleteMacAddressByGroup/{macGroup}", middleware.DeleteMacAddressByGroup).Methods("DELETE")

	router.HandleFunc("/getMacAddressesByGroup/{macGroup}", middleware.GetMacAddressesByGroup).Methods("GET")

	return router
}
