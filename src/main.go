package main

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func main() {
	log.Println("Running...")
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	var result string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`http://cnpj.info/Achei-Quimioterapia-e-Oncologia-Ltda`),
		chromedp.WaitVisible(`#content > table > tbody > tr:nth-child(1) > td:nth-child(2) > a`),
		chromedp.Text(`#content > table > tbody > tr:nth-child(1) > td:nth-child(2) > a`,
			&result),
		chromedp.Stop(),
	)

	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("CNPJ %q", result)

}
