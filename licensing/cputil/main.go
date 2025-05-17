package main

import (
	"fmt"
	"github.com/lgirma/gofx/licensing"
	"log"
)

func main() {

	fmt.Println("GoFX Copy Protection Utility")
	productName, _ := ReadString("Product Name: ")
	edition, _ := ReadString("Product Edition: ")
	reqCode, _ := ReadString("Request Code: ")
	authCode, _ := ReadString("Auth Code: ")

	service := licensing.NewCopyProtectionService(nil, licensing.Options{
		Product:  licensing.Product{Name: productName, Edition: edition},
		AuthCode: authCode,
	})
	activationCode, err := service.GetActivationCode(reqCode)
	if err != nil {
		log.Fatalf("Error getting activation code: %v", err)
	} else {
		fmt.Printf("Activation Code: %s\n", activationCode)
	}

}
