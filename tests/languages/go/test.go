package main

import appwrite "github.com/appwrite/go-sdk"

func main() {
	defaultClient := appwrite.NewDefaultClient()
	defaultClient.AddHeader("Origin", "http://localhost")
	// client.SetSelfSigned()
	println("Test Started")
}
