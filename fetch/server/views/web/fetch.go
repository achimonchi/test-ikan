package web

type GetListEfishery struct {
	UUID         interface{} `json:"uuid"`
	Komoditas    interface{} `json:"komoditas"`
	AreaProvinsi interface{} `json:"area_provinsi"`
	AreaKota     interface{} `json:"area_kota"`
	Size         interface{} `json:"size"`
	PriceUSD     interface{} `json:"price_usd"`
	PriceIDR     interface{} `json:"price"`
	TglParsed    interface{} `json:"tgl_parsed"`
	Timestamp    interface{} `json:"timestamp"`
}
