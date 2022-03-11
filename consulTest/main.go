package main

import (
	"fmt"

	consulAPI "github.com/hashicorp/consul/api"
)

func main() {
	consulClient, _ := consulAPI.NewClient(consulAPI.DefaultConfig())
	kvClient := consulClient.KV()

	//	GET
	pair, _, err := kvClient.Get("KVCli", nil)
	if err != nil {
		fmt.Println("There is an error: ", err)
	}

	if pair == nil {
		fmt.Println("Key does not exist")
	} else {
		fmt.Println(string(pair.Value))
	}

	//	PUT
	// p := &api.KVPair{Key: "KVFirst", Value: []byte("YelloBro")}
	// _, err := kvClient.Put(p, nil)
	// if err != nil {
	// 	panic(err)
	// }
}
