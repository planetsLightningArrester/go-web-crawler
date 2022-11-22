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

	var cnpj string
	var nome_da_empresa string
	var inicio_da_atividade string
	var natureza_juridica string
	var situacao_cadastral string
	var motivo_situacao_cadastral string
	var qualificacao_do_responsavel string
	var capital_social string
	var porte_da_empresa string
	var opcao_pelo_simples string
	var opcao_pelo_mei string

	err := chromedp.Run(ctx,
		chromedp.Navigate(`http://cnpj.info/Achei-Quimioterapia-e-Oncologia-Ltda`),
		chromedp.WaitVisible(`#content > table > tbody > tr:nth-child(1) > td:nth-child(2) > a`),
		chromedp.TextContent(`#content > table > tbody > tr:nth-child(1) > td:nth-child(2) > a`, &cnpj),
		chromedp.TextContent(`#content > table > tbody > tr:nth-child(2) > td:nth-child(2)`, &nome_da_empresa),
		chromedp.TextContent(`#content > table > tbody > tr:nth-child(3) > td:nth-child(2)`, &inicio_da_atividade),
		chromedp.TextContent(`#content > table > tbody > tr:nth-child(4) > td:nth-child(2)`, &natureza_juridica),
		chromedp.TextContent(`#content > table > tbody > tr:nth-child(5) > td:nth-child(2)`, &situacao_cadastral),
		chromedp.TextContent(`#content > table > tbody > tr:nth-child(6) > td:nth-child(2)`, &motivo_situacao_cadastral),
		chromedp.TextContent(`#content > table > tbody > tr:nth-child(7) > td:nth-child(2)`, &qualificacao_do_responsavel),
		chromedp.TextContent(`#content > table > tbody > tr:nth-child(8) > td:nth-child(2)`, &capital_social),
		chromedp.TextContent(`#content > table > tbody > tr:nth-child(9) > td:nth-child(2)`, &porte_da_empresa),
		chromedp.TextContent(`#content > table > tbody > tr:nth-child(10) > td:nth-child(2)`, &opcao_pelo_simples),
		chromedp.TextContent(`#content > table > tbody > tr:nth-child(11) > td:nth-child(2)`, &opcao_pelo_mei),
		chromedp.Stop(),
	)

	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("CNPJ %q", cnpj)
	log.Printf("nome_da_empresa %q", nome_da_empresa)
	log.Printf("inicio_da_atividade %q", inicio_da_atividade)
	log.Printf("natureza_juridica %q", natureza_juridica)
	log.Printf("situacao_cadastral %q", situacao_cadastral)
	log.Printf("motivo_situacao_cadastral %q", motivo_situacao_cadastral)
	log.Printf("qualificacao_do_responsavel %q", qualificacao_do_responsavel)
	log.Printf("capital_social %q", capital_social)
	log.Printf("porte_da_empresa %q", porte_da_empresa)
	log.Printf("opcao_pelo_simples %q", opcao_pelo_simples)
	log.Printf("opcao_pelo_mei %q", opcao_pelo_mei)

}
