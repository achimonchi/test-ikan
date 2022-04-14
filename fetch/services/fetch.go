package services

import (
	"encoding/json"
	"fetch/config"
	"fetch/pkg/httpclient"
	"fetch/server/views/web"
	"fmt"
	"strconv"
)

var USD_IDR float32 = 0

type FetchServices struct {
	client    *httpclient.HttpClient
	converter *httpclient.HttpClient
	config    *config.Config
}

func NewFetchServices(client *httpclient.HttpClient, converter *httpclient.HttpClient, config *config.Config) *FetchServices {
	return &FetchServices{
		client:    client,
		converter: converter,
		config:    config,
	}
}

func (f *FetchServices) GetList() (*[]web.GetListEfishery, error) {
	data, err := f.client.GetList("v1/storages/5e1edf521073e315924ceab4/list", map[string]string{})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var list []web.GetListEfishery

	dataByte, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	err = json.Unmarshal(dataByte, &list)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	if USD_IDR == 0 {
		fmt.Println("hit converter")
		val, err := f.converter.Get("api/v7/convert?q=USD_IDR&compact=ultra&apiKey="+f.config.CONVERTER_API, map[string]string{})
		if err != nil {
			USD_IDR = 0
			fmt.Println(err.Error())
		} else {
			usd, ok := val["USD_IDR"]
			if !ok {
				USD_IDR = 0
			} else {
				usd_idr, err := strconv.ParseFloat(fmt.Sprintf("%v", usd), 32)
				if err != nil {
					USD_IDR = 0
					fmt.Println(err.Error())
				} else {
					USD_IDR = float32(usd_idr)
				}
			}
		}
	}

	for i, l := range list {
		if l.PriceIDR != nil {
			priceIDR, _ := strconv.ParseFloat(fmt.Sprintf("%v", l.PriceIDR), 32)
			l.PriceUSD = fmt.Sprintf("%v", priceIDR/float64(USD_IDR))
		}

		list[i] = l
	}

	return &list, nil
}
