package main
import (
	"fmt"
	"log"
	"context"
	"crypto/ecdsa"
	"math/big"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	Todo "github.com/idirall22/go-solidity-learn/gen"
)
var (
	URL = "wss://rinkeby.infura.io/ws/v3/14a9f178d7564fecac3afa7e62d095ff"
)

func main() {
	//------------------------------------------------------------
	PrivateKey, err := crypto.HexToECDSA("381c1e5cbe5a6dc7e34e58f65dc55bf957859b7331296c9d6d679b58666aae37")
    if err != nil {
        log.Fatal(err)
    }
	//------------------------------------------------------------
    publicKey := PrivateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("error casting public key to ECDSA")
    }
	//------------------------------------------------------------
    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//------------------------------------------------------------
	client, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatal(err)
	}
	//------------------------------------------------------------
	nonce, err := client.PendingNonceAt(context.Background(),fromAddress) 
	if err != nil {
		log.Fatal(err)
	}
	//------------------------------------------------------------
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//------------------------------------------------------------
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//------------------------------------------------------------
	auth, err := bind.NewKeyedTransactorWithChainID(PrivateKey,chainID)
	if err != nil {
		log.Fatal(err)
	}
	//------------------------------------------------------------
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice
	//------------------------------------------------------------
	a, tx, _, err := Todo.DeployTodo(auth,client)
	if err != nil {
		log.Fatal(err)
	}	
	//------------------------------------------------------------
	fmt.Println("Contract Address: ",a.Hex())
	fmt.Println("Transaction Address: ",tx.Hash().Hex())
}