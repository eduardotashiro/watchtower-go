package scraping

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

// Notificador
type SlackPayload struct {
	Text string `json:"text"`
}

// Estado del servicio
type ServiceStatus struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

func CheckServiceStatus() *SlackPayload {

	url := map[string]string{
		"Banco do Brasil": "https://downdetector.com.br/fora-do-ar/banco-do-brasil/",
		"Bradesco":        "https://downdetector.com.br/fora-do-ar/bradesco/",
		"Santander":       "https://downdetector.com.br/fora-do-ar/santander/",
		"Pix":             "https://downdetector.com.br/fora-do-ar/pix/",
		"Pic Pay":         "https://downdetector.com.br/fora-do-ar/picpay/",
		"Banco Itaú":      "https://downdetector.com.br/fora-do-ar/banco-itau/",
		"Nubank":          "https://downdetector.com.br/fora-do-ar/nubank/",
		"Mercado Pago":    "https://downdetector.com.br/fora-do-ar/mercadopago/",
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
			fmt.Printf("error scrapeando, %s:%v\n", name, err)
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
		case false:
			message += fmt.Sprintf("%v : %v | bien\n\n", service.Name, service.Status)
		case true:
			message += fmt.Sprintf("%v : %v | malo\n\n", service.Name, service.Status)
		}
	}

	payload := &SlackPayload{
		Text: message,
	}

	return payload

}
