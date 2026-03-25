package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {

	allocCtx, cancel := chromedp.NewRemoteAllocator(
		context.Background(),
		"http://localhost:9222",
	)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	var outage bool

	err := chromedp.Run(ctx,
		chromedp.Navigate("https://downdetector.com/es/problemas/battle-net/"),

		chromedp.Sleep(5*time.Second),

		chromedp.Evaluate(`
                        (function() {
                                if (window.PogoConfig && window.PogoConfig.outage !== undefined) {
                                        return window.PogoConfig.outage;
                                }
                                return false;
                        })()
                `, &outage),
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Outage......>", outage)

}
