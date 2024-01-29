package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/foliveiracamara/bents-api/adapter/driver/lambda/server"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var echoLambda *echoadapter.EchoLambdaV2

func init() {
	isLambda := os.Getenv("LAMBDA")

	if isLambda == "TRUE" {
		e := echo.New()
		server := server.Server{
			Echo: e,
		}
		e = server.InitRoutes(isLambda).(*echo.Echo)
		echoLambda = echoadapter.NewV2(e)
	}
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return echoLambda.Proxy(req)
}

func main() {
	isLambda := os.Getenv("LAMBDA")

	if isLambda == "TRUE" {
		lambda.Start(Handler)
	} else {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
		e := echo.New()
		server := server.Server{
			Echo: e,
		}
		server.InitRoutes(isLambda)
	}
}
