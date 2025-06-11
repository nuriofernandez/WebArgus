package main

import (
	_ "github.com/joho/godotenv/autoload" // Important to keep this as the first import!
	"os"

	"fmt"
	"github.com/nuriofernandez/WebArgus/checker"
	"github.com/nuriofernandez/WebArgus/notify"
	"github.com/nuriofernandez/WebArgus/orders"
)

func main() {
	fmt.Println("[Argus] Starting...")

	fmt.Println("[TESTING MODE] Testing phone is: " + os.Getenv("TEST_ORDER_PHONE"))

	fmt.Println("[Argus] Preparing orders checker...")
	order, err := orders.NextOrder()
	if err != nil {
		fmt.Println("[Argus] There is no orders available")
		return
	}

	isOnline := checker.IsOnline(order.Url)
	if isOnline {
		fmt.Println("[Argus] Order passed check!")

		fmt.Println("[Argus] Preparing notification...")
		notify.Notify(order)
	}

	fmt.Println("[Argus] Operation completed.")
}
