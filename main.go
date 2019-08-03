package main

import (
	"context"
	"fmt"
	"log"
	"time"

	firebase "firebase.google.com/go"
	"github.com/micro/go-micro"
	"github.com/spf13/viper"
	"google.golang.org/api/option"

	"github.com/sisimogangg/supermarket.discount.api/repository"
	"github.com/sisimogangg/supermarket.discount.api/service"

	"github.com/sisimogangg/supermarket.discount.api/utils"

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

func seeding(app *firebase.App) {
	ctx := context.Background()
	client, err := app.Database(ctx)
	if err != nil {
		log.Fatal(err)
	}

	var rawDiscounts map[string]pb.ProductDiscount
	err = client.NewRef("discounts").Get(ctx, &rawDiscounts)
	if err != nil {
		log.Fatal(err)
	}

	if len(rawDiscounts) == 0 {
		for _, d := range utils.ProductDiscounts {
			if err := client.NewRef(fmt.Sprintf("discounts/%s", d.DiscountID)).Set(ctx, d); err != nil {
				log.Fatal(err)
			}
		}
	}

}

func initializeFirebase() *firebase.App {
	opt := option.WithCredentialsFile("firebaseServiceAccount.json")

	ctx := context.Background()
	config := &firebase.Config{
		DatabaseURL: "https://supermarket-8aee3.firebaseio.com",
	}

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatal(err)
	}
	return app
}

func main() {
	srv := micro.NewService(
		micro.Name("supermarket.discount"),
		micro.Version("latest"),
	)

	srv.Init()

	app := initializeFirebase()

	if viper.GetBool("debug") {
		seeding(app)
	}

	repo := repository.NewFirebaseRepo(app)

	timeContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	servicelayer := service.NewDicountService(repo, timeContext)

	pb.RegisterDiscountServiceHandler(srv.Server(), servicelayer)

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
