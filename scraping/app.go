package scraping

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func CheckServiceStatus() bool {

	url := map[string]string{
		"BANCO_DO_BRASIL": "https://downdetector.com.br/fora-do-ar/banco-do-brasil/",
		"BRADESCO":        "https://downdetector.com.br/fora-do-ar/bradesco/",
		"SANTANDER":       "https://downdetector.com.br/fora-do-ar/santander/",
		"PIX":             "https://downdetector.com.br/fora-do-ar/pix/",
		"PICPAY":          "https://downdetector.com.br/fora-do-ar/picpay/",
		"ITAU":            "https://downdetector.com.br/fora-do-ar/banco-itau/",
		"NUBANK":          "https://downdetector.com.br/fora-do-ar/nubank/",
		"MERCADO_PAGO":    "https://downdetector.com.br/fora-do-ar/mercadopago/",
		"SPARKLIGHT":      "https://downdetector.com/es/problemas/sparklight/",
	}

	var outage bool

	for n, u := range url {

		allocCtx, cancel := chromedp.NewRemoteAllocator(
			context.Background(),
			"http://localhost:9222",
		)
		defer cancel()

		ctx, cancel := chromedp.NewContext(allocCtx)
		defer cancel()

		err := chromedp.Run(ctx,
			chromedp.Navigate(u),
			chromedp.Sleep(2*time.Second),

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

		fmt.Printf("service: %s | status: %t\n", n, outage)
	}
	return outage
}
