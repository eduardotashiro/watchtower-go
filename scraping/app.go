package scraping

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func CheckServiceStatus() []string{

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

	var (
		outage       bool
		responseBody *bytes.Buffer
		acumulate    []string
	)

	for name, url := range url {

		allocCtx, cancel := chromedp.NewRemoteAllocator(
			context.Background(),
			"http://localhost:9222",
		)
		defer cancel()

		ctx, cancel := chromedp.NewContext(allocCtx)
		defer cancel()

		err := chromedp.Run(ctx,
			chromedp.Navigate(url),
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

		// acumulate = //acumular com append

		data := map[string]map[string]interface{}{
			"text": {
				"service": name,
				"status":  outage,
			},
		}

		// for c,v := range data{
		// 	fmt.Println(c, v)
		// 	test := append(acumulate,)

		// }

		postBody, err := json.Marshal(data)

		if err != nil {
			log.Fatalf("erro ao converter p JSON: %v", err)
		}

		responseBody = bytes.NewBuffer(postBody)

		fmt.Println(responseBody)
	}
	return acumulate
}

// expectativa
//{"text":"Hello, World!"}

// realidade
//{"text":{"service":"NUBANK","status":false}}
