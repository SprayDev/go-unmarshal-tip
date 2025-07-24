package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Response struct {
	Items []Item `json:"items"`
}

type Item struct {
	Status *ItemStatus `json:"status"`
	Name   string      `json:"name"`
}

type ItemStatus struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ItemAlias Item

func (i *Item) UnmarshalJSON(b []byte) error {
	var raw struct {
		ItemAlias
		Status json.RawMessage `json:"status"`
	}

	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	if string(raw.Status) == "[]" || string(raw.Status) == "\"null\"" || string(raw.Status) == "null" {
		i.Status = nil
		return nil
	}

	if err := json.Unmarshal(raw.Status, &i.Status); err != nil {
		return err
	}
	return nil
}

func main() {
	file, err := os.Open("data.json")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	bytes, _ := io.ReadAll(file)

	var response Response

	err = json.Unmarshal(bytes, &response)

	if err != nil {
		panic(err)
	}

	for _, item := range response.Items {
		fmt.Printf("%+v\n", item.Status)
	}
}
