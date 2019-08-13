package main

import (
	"fmt"
	"gokata/hdwallet"

	"github.com/btcsuite/btcd/chaincfg"
)

var (
	mnemonic = "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about"
)

func main() {
	master, err := hdwallet.NewKey(
		hdwallet.Mnemonic(mnemonic),
		hdwallet.Password("secret password $"), // <---- your secret password here
		hdwallet.Params(&chaincfg.MainNetParams),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n~~~~~~~~~~~ BITCOIN MAINNET ~~~~~~~~~~~")
	for i := 0; i < 3; i++ {
		wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.BTC), hdwallet.AddressIndex(uint32(i)), hdwallet.Params(&chaincfg.MainNetParams))
		legacy, _ := wallet.GetAddress()
		addressP2WPKH, _ := wallet.GetKey().AddressP2WPKH()
		addressP2WPKHInP2SH, _ := wallet.GetKey().AddressP2WPKHInP2SH()

		wifCompressed, _ := wallet.GetKey().PrivateWIF(true)
		fmt.Printf("WIF (compressed): %s\n", wifCompressed)
		fmt.Printf("Address %s\n", addressP2WPKH)
		fmt.Printf("Segwit %s\n", addressP2WPKHInP2SH)
		fmt.Printf("Legacy %s\n", legacy)
		fmt.Println("------------")
	}

	master, err = hdwallet.NewKey(
		hdwallet.Mnemonic(mnemonic),
		hdwallet.Password("secret password $"),
		hdwallet.Params(&chaincfg.TestNet3Params),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n\n~~~~~~~~~~~ BITCOIN TESTNET ~~~~~~~~~~~\n\n")
	for i := 0; i < 3; i++ {
		wallet, _ := master.GetWallet(hdwallet.AddressIndex(uint32(i)), hdwallet.Params(&chaincfg.TestNet3Params))
		legacy, _ := wallet.GetAddress()
		addressP2WPKH, _ := wallet.GetKey().AddressP2WPKH()
		addressP2WPKHInP2SH, _ := wallet.GetKey().AddressP2WPKHInP2SH()

		wifCompressed, _ := wallet.GetKey().PrivateWIF(true)
		fmt.Printf("WIF (compressed): %s\n", wifCompressed)
		fmt.Printf("Address %s\n", addressP2WPKH)
		fmt.Printf("Segwit %s\n", addressP2WPKHInP2SH)
		fmt.Printf("Legacy %s\n", legacy)
		fmt.Println("------------")
	}

}
