package csv

// --------- EXAMPLE HOW TO EXPORT 'EM AS CSV

import (
	"encoding/csv"
	"fmt"
	"gokata/hdwallet"
	"log"
	"os"
	"strconv"

	"github.com/btcsuite/btcd/chaincfg"
)

// ExportBTCToCSV ..... does just that
func ExportBTCToCSV() {

	file, err := os.Create("btc_testnet.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	mnemonic := "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about"

	master, err := hdwallet.NewKey(
		hdwallet.Mnemonic(mnemonic),
		hdwallet.Password("test"),
	)
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)
	writer.Comma = ';'

	defer writer.Flush()

	for i := 0; i <= 1000; i++ {
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

		err = writer.Write([]string{strconv.Itoa(i), "BTCTEST", addressP2WPKH, addressP2WPKHInP2SH, legacy})
		if err != nil {
			panic(err)
		}

		if i%100 == 0 {
			fmt.Printf("Done %d records\n", i)
		}
	}

}
