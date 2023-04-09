package main

import (
	"context"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/S-Ryouta/notice-latest-program-version"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	if err := funcframework.RegisterEventFunctionContext(ctx, "/", notice_latest_program_version.CheckAndUpdateVersionHandler); err != nil {
		log.Fatalf("funcframework.RegisterEventFunctionContext: %v\n", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
