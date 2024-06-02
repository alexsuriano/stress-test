package tester

import "flag"

type Configs struct {
	Url         string
	Requests    int64
	Concurrency int64
}

func GetParams() *Configs {
	var url string
	var requests int64
	var concurrency int64

	flag.StringVar(&url, "url", "", "URL do serviço a ser testado")
	flag.Int64Var(&requests, "requests", 1, "Número total de requests")
	flag.Int64Var(&concurrency, "concurrency", 1, "Número de chamadas simultâneas")
	flag.Parse()

	return &Configs{
		Url:         url,
		Requests:    requests,
		Concurrency: concurrency,
	}
}
