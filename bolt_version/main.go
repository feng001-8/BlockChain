package main

//import "fmt"

func main() {
	bc := NewBlockChain("1EfMZBiKMiE4XnLARzTjnWxHFqx2xZQLFq", 10000000)
	//newTx := NewCoinbaseTX("ls", "无", 20)
	//bc.AddBlock([]*Transaction{newTx})
	cli := CLI{bc}
	cli.Run()
}

//bc.AddBlock("111111111111111")
//bc.AddBlock("222222222222222")
//
