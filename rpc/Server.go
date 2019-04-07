package main

import (
	"blockchain/core"
	"encoding/json"
	"io"
	"net/http"
)

var blockchain *core.Blockchain

func run() {
	//获取链上的数据
	http.HandleFunc("/blockchain/get", blockchainGetHandler)
	//写数据到链上
	http.HandleFunc("/blockchain/write", blockchainWriteHandler)
	//启动对端口的监听
	http.ListenAndServe("localhost:8888", nil)
}

//获取链上的数据
func blockchainGetHandler(w http.ResponseWriter, r *http.Request) {
	//转化为json数据格式
	bytes, error := json.Marshal(blockchain)
	//如果erroe不为空
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

//往区块链上写数据
func blockchainWriteHandler(w http.ResponseWriter, r *http.Request) {
	blockData := r.URL.Query().Get("data")
	//发送数据
	blockchain.SendData(blockData)
	//把最新的区块链数据会显给调用者
	blockchainGetHandler(w, r)
}

//启动之后访问： http://localhost:8888/blockchain/get
//添加数据到区块链： http://localhost:8888/blockchain/write?data=hello world
func main() {
	//创建一个blockchain
	blockchain = core.NewBlockchain()
	run()
}
