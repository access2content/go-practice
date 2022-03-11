package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

type Input struct {
	SecretKey string                 `json:"secretKey"`
	Method    string                 `json:"method"`
	URL       string                 `json:"URL"`
	Body      map[string]interface{} `json:"body"`
}

//	Signing a message with Private Key
func Sign(privateKey string, message string) string {
	priv, _ := hex.DecodeString(privateKey)
	return hex.EncodeToString(ed25519.Sign(ed25519.NewKeyFromSeed(priv), []byte(message)))
}

func main() {

	//	READ the file path
	var path string
	flag.StringVar(&path, "path", "inputs.json", "File path containing the inputs")
	flag.Parse()

	fmt.Println("Reading file:", path)

	//	READ the file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Could not find the specified file")
		return
	}
	fmt.Println("----------	Input	----------")
	fmt.Printf("%+v", string(data))

	//	CONVERT the input to required format
	userInput := &Input{}
	err = json.Unmarshal(data, &userInput)
	if err != nil {
		fmt.Println("Error in input data format")
		return
	}
	inputBody, err := json.Marshal(userInput.Body)
	if err != nil {
		fmt.Println("Incorrect body format")
		return
	}

	//	CREATE the message to sign
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	message := timestamp + userInput.Method + userInput.URL + string(inputBody)

	fmt.Println("----------	Message	----------")
	fmt.Println(message)

	//	SIGN the message
	signature := Sign(userInput.SecretKey, message)
	fmt.Println("----------	Signature	----------")
	fmt.Println(signature)
	fmt.Println("----------	Timestamp	----------")
	fmt.Println(timestamp)
}

/*
   "publicKey": "09c81b0e973eb70209a9aa50251bbae4f3f2013993e6bea39fb517da09985a84",
   "secretKey": "74c750eb01a1a77dcac38f0ebb9b80d4f240374d13d6ca4397b99463a03e9636"
*/
