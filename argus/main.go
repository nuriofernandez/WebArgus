package main

import (
	_ "github.com/joho/godotenv/autoload" // Important to keep this as the first import!
	"github.com/nuriofernandez/WebArgus/cron"
	"github.com/nuriofernandez/WebArgus/http"
	"os"

	"fmt"
)

func main() {
	fmt.Println("[Argus] Starting...")

	fmt.Println("[TESTING MODE] Testing phone is: " + os.Getenv("TEST_ORDER_PHONE"))

	fmt.Println("[Argus] Starting cron worker...")
	go cron.Start()

	fmt.Println("[Argus] Starting http server...")
	http.Start()
}
