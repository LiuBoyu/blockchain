package main

import "blockchain/core"

func main() {
    bc := core.NewBlockchain()
    bc.SendData("data:Send 1 BTC to clx")
    bc.SendData("data:Send 1 EOS to clx")
    bc.Print()
}
