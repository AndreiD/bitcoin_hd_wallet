package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

// ExportETHToCSV ..... does just that
func ExportETHToCSV() {

	mnemonic := "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about"

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("eth.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ';'

	defer writer.Flush()

	for i := 0; i < 1000; i++ {

		path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/" + strconv.Itoa(i))
		account, err := wallet.Derive(path, false)
		if err != nil {
			log.Fatal(err)
		}

		err = writer.Write([]string{strconv.Itoa(i), "ETH", account.Address.Hex()})
		if err != nil {
			panic(err)
		}

		if i%100 == 0 {
			fmt.Printf("Done %d records\n", i)
		}

	}

}
