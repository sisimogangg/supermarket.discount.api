package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	data "github.com/sisimogangg/supermarket.discount.api/discount/dataaccess"
	"github.com/sisimogangg/supermarket.discount.api/discount/service"
	"github.com/sisimogangg/supermarket.discount.api/handlers"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool("debug") {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func start(router *mux.Router) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}

func main() {
	router := mux.NewRouter()

	repo := data.NewFirebaseRepo()

	timeContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	servicelayer := service.NewDicountService(repo, timeContext)

	handlers.NewDiscountHandler(router, servicelayer)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}

}
