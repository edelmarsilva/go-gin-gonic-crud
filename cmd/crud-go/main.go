package main

import (
	"crud-go/internal/route"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	//carregar arquivos de configuracao do .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Estudo de aplicacao escrito por :")
	fmt.Println(os.Getenv("AUTOR"))

	//gin.Default cria um logger e middleware de recovey
	//gin.New tera que criar tudo do zero
	router := gin.Default()
	//apontando diretamente para o objeto original
	route.InitRouter(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
