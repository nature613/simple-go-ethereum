package main


import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	store "../../contracts" // for demo
)

func main() {
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/0b650af********************397ac")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x71733bc1d86F5de7bCBf0587502bB1BeE176E788")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version) // "1.0"
}
