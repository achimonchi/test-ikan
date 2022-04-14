package services

import (
	"fetch/pkg/httpclient"
	"fmt"
)

type FetchServices struct {
	client *httpclient.HttpClient
}

func NewFetchServices(client *httpclient.HttpClient) *FetchServices {
	return &FetchServices{
		client: client,
	}
}

func (f *FetchServices) GetList() {
	data, err := f.client.Get("v1/storages/5e1edf521073e315924ceab4/list", map[string]string{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%+v\n", data)
}
