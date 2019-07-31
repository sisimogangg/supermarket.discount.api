package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/micro/go-micro"
	"github.com/spf13/viper"

	"github.com/sisimogangg/supermarket.discount.api/repository"
	"github.com/sisimogangg/supermarket.discount.api/service"

	pb "github.com/sisimogangg/supermarket.discount.api/proto"
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
	srv := micro.NewService(
		micro.Name("supermarket.discount"),
		micro.Version("latest"),
	)

	srv.Init()

	repo := repository.NewFirebaseRepo()

	timeContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	servicelayer := service.NewDicountService(repo, timeContext)

	pb.RegisterDiscountServiceHandler(srv.Server(), servicelayer)

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
