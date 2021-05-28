package main

import (
	appwrite "../../sdks/go"
	// appwrite "../../../examples/go"
)

func main() {
	defaultClient := appwrite.NewDefaultClient()
	defaultClient.AddHeader("Origin", "http://localhost")
	// client.SetSelfSigned()
	println("Test Started")
}
