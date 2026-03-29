package scraping

import (
	"context"
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
	}

	var outage bool

	for _, u := range url {

		allocCtx, cancel := chromedp.NewRemoteAllocator(
			context.Background(),
			"http://localhost:9222",
		)
		defer cancel()

		ctx, cancel := chromedp.NewContext(allocCtx)
		defer cancel()

		err := chromedp.Run(ctx,
			chromedp.Navigate(u),
			chromedp.Sleep(4*time.Second),

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

	}
	return outage
}
