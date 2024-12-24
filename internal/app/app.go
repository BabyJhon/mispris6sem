package app

import (
	"context"
	"os"

	config "github.com/BabyJhon/mispris1-2/configs"
	"github.com/BabyJhon/mispris1-2/internal/handlers"
	"github.com/BabyJhon/mispris1-2/internal/repo"
	"github.com/BabyJhon/mispris1-2/pkg/postgres"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Run() {
	//logrus
	logrus.SetFormatter(new(logrus.JSONFormatter))

	//Configs
	if err := config.InitConfig(); err != nil {
		logrus.Fatalf("error initialization configs: %s", err.Error())
	}

	//.env
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env vars: %s", err.Error())
	}
	//fmt.Println(viper.GetString("db.host") + " zzzzzzzzzzzzz")
	//DB
	pool, err := postgres.NewPG(context.Background(), postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("error while init db:%s", err.Error())
	}

	//pool.Query(context.Background(), "")

	//repo
	repo := repo.NewRepositiry(pool)

	//cli
	h := handlers.NewHandlers(repo)
	cli := handlers.NewCLInterface(h)
	cli.StartCli()
}
