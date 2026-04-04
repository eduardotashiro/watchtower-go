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

// status
type ServiceStatus struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

// todos serviços
type AllServices struct {
	Services []ServiceStatus `json:"services"`
	Time     string          `json:"timestamp"`
}

func CheckServiceStatus() *bytes.Buffer {

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

	var services []ServiceStatus

	for name, serviceURL := range url {

		allocCtx, cancel := chromedp.NewRemoteAllocator(
			context.Background(),
			"http://localhost:9222",
		)
		defer cancel()

		ctx, cancel := chromedp.NewContext(allocCtx)
		defer cancel()

		var outage bool

		err := chromedp.Run(ctx,
			chromedp.Navigate(serviceURL),
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

		fmt.Printf("service: %s | status: %t\n", name, outage)

		services = append(services, ServiceStatus{
			Name:   name,
			Status: outage,
		})
	}

	allServices := AllServices{
		Services: services,
		Time:     time.Now().Format(time.RFC3339),
	}

	JsonData, err := json.Marshal(allServices)
	if err != nil {
		log.Fatal(err)
	}

	return bytes.NewBuffer(JsonData)

}
