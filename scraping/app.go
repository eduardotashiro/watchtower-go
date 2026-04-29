package scraping

import (
	"fmt"
	"log"
	"strings"
	"time"

	cu "github.com/Davincible/chromedp-undetected"
	"github.com/chromedp/chromedp"
)

type SlackPayload struct {
	Text string `json:"text"`
}

type ServiceStatus struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

var Services = map[string]string{
	"Banco do Brasil": "https://downdetector.com.br/fora-do-ar/banco-do-brasil/",
	"Bradesco":        "https://downdetector.com.br/fora-do-ar/bradesco/",
	"Santander":       "https://downdetector.com.br/fora-do-ar/santander/",
	"Pix":             "https://downdetector.com.br/fora-do-ar/pix/",
	"Pic Pay":         "https://downdetector.com.br/fora-do-ar/picpay/",
	"Banco Itaú":      "https://downdetector.com.br/fora-do-ar/banco-itau/",
	"Nubank":          "https://downdetector.com.br/fora-do-ar/nubank/",
	"Mercado Pago":    "https://downdetector.com.br/fora-do-ar/mercadopago/",
	"instagram":       "https://downdetector.com.br/fora-do-ar/instagram/",
}

func CheckServiceStatus() *SlackPayload {
	var services []ServiceStatus

	ctx, cancel, err := cu.New(cu.NewConfig(
		cu.WithTimeout(30 * time.Second),
	))
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()

	for name, serviceURL := range Services {
		var outage bool
		var bodyText string

		err := chromedp.Run(ctx,
			chromedp.Navigate(serviceURL),
			chromedp.Sleep(4*time.Second),
			chromedp.Evaluate(`
(function() {
	if (window.PogoConfig && window.PogoConfig.outage !== undefined) {
		return window.PogoConfig.outage;
	}
	return false;
})()
`, &outage),
			chromedp.Evaluate(`document.body.innerText.toLowerCase()`, &bodyText),
		)

		status := "success"

		if outage {
			status = "danger"
		} else if strings.Contains(bodyText, "não mostram problemas") {
			status = "success"
		} else if strings.Contains(bodyText, "possíveis problemas") {
			status = "warning"
		} else if strings.Contains(bodyText, "mostram problemas") {
			status = "danger"
		}

		if err != nil {
			fmt.Printf("error scrapeando, %s:%v\n", name, err)
		}

		fmt.Printf("service: %s | status: %s\n", name, status)

		services = append(services, ServiceStatus{
			Name:   name,
			Status: status,
		})
	}

	var message string
	for _, service := range services {
		switch service.Status {
		case "success":
			message += fmt.Sprintf("%s : normal\n\n", service.Name)
		case "warning":
			message += fmt.Sprintf("%s : instável\n\n", service.Name)
		case "danger":
			message += fmt.Sprintf("%s : fora do ar\n\n", service.Name)
		}
	}

	return &SlackPayload{Text: message}
}
