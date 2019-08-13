package main

import (
	"fmt"
	"gokata/hdwallet"

	"github.com/btcsuite/btcd/chaincfg"
)

const mnemonic = "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about"

func main() {
	master, err := hdwallet.NewKey(
		hdwallet.Mnemonic(mnemonic),
		hdwallet.Password("secret password $"), // <---- your secret password here
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n~~~~~~~~~~~ BITCOIN MAINNET ~~~~~~~~~~~")
	for i := 0; i < 3; i++ {
		wallet, err := master.GetWallet(hdwallet.CoinType(hdwallet.BTC), hdwallet.AddressIndex(uint32(i)), hdwallet.Params(&chaincfg.MainNetParams))
		if err != nil {
			panic(err)
		}
		legacy, err := wallet.GetAddress()
		if err != nil {
			panic(err)
		}
		addressP2WPKH, err := wallet.GetKey().AddressP2WPKH()
		if err != nil {
			panic(err)
		}
		addressP2WPKHInP2SH, err := wallet.GetKey().AddressP2WPKHInP2SH()
		if err != nil {
			panic(err)
		}

		wifCompressed, err := wallet.GetKey().PrivateWIF(true)
		if err != nil {
			panic(err)
		}
		fmt.Printf("WIF (compressed): %s\n", wifCompressed)
		fmt.Printf("Address %s\n", addressP2WPKH)
		fmt.Printf("Segwit %s\n", addressP2WPKHInP2SH)
		fmt.Printf("Legacy %s\n", legacy)
		fmt.Println("------------")
	}

	fmt.Printf("\n\n~~~~~~~~~~~ BITCOIN TESTNET ~~~~~~~~~~~\n\n")
	for i := 0; i < 3; i++ {
		wallet, err := master.GetWallet(hdwallet.AddressIndex(uint32(i)), hdwallet.Params(&chaincfg.TestNet3Params))
		if err != nil {
			panic(err)
		}
		legacy, err := wallet.GetAddress()
		if err != nil {
			panic(err)
		}
		addressP2WPKH, err := wallet.GetKey().AddressP2WPKH()
		if err != nil {
			panic(err)
		}
		addressP2WPKHInP2SH, err := wallet.GetKey().AddressP2WPKHInP2SH()
		if err != nil {
			panic(err)
		}

		wifCompressed, err := wallet.GetKey().PrivateWIF(true)
		if err != nil {
			panic(err)
		}
		fmt.Printf("WIF (compressed): %s\n", wifCompressed)
		fmt.Printf("Address %s\n", addressP2WPKH)
		fmt.Printf("Segwit %s\n", addressP2WPKHInP2SH)
		fmt.Printf("Legacy %s\n", legacy)
		fmt.Println("------------")
	}

}
