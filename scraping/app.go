package scraping

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

// Notifier
type SlackPayload struct {
	Text string `json:"text"`
}

// Service Status
type ServiceStatus struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

func CheckServiceStatus() *SlackPayload {

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

	allocCtx, cancel := chromedp.NewRemoteAllocator(
		context.Background(),
		"http://localhost:9222",
	)
	defer cancel()

	for name, serviceURL := range url {

		ctx, cancel := chromedp.NewContext(allocCtx)

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
		cancel()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("service: %s | status: %t\n", name, outage)

		services = append(services, ServiceStatus{
			Name:   name,
			Status: outage,
		})
	}

	var message string

	for _, service := range services {
		switch service.Status {
		case true:
			message += fmt.Sprintf("%v : %v\n\n", service.Name, service.Status)
		}
	}

	payload := &SlackPayload{
		Text: message,
	}

	return payload

}

//ERR_INSUFFICIENT_RESOURCES
