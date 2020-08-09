package main

import (
	"fmt"

	"../rod"
)

//var mem []primer

type primer struct {
	a   map[string]int
	b   []int
	mem []primer
}

func (p *primer) writeMemory() {
	var temp primer
	temp.a = make(map[string]int)
	temp.b = make([]int, 0, 1)
	for k, v := range p.a {
		temp.a[k] = v
	}
	for _, v := range p.b {
		temp.b = append(temp.b, v)
	}

	fmt.Println(temp)
	p.mem = append(p.mem, temp) // TEMP: )
}

func main() {
	fmt.Println("Start process...")
	//var a rod.Famile
	//a.Init("piple")
	//fmt.Println(a)
	var p primer
	p.a = make(map[string]int)
	p.b = make([]int, 0, 1)
	p.mem = make([]primer, 0, 1)

	p.a["1"] = 1
	p.b = append(p.b, 1)
	p.writeMemory()
	fmt.Println(p)

	p.a["1"] = 2
	p.b = append(p.b, 2)
	p.writeMemory()
	fmt.Println(p)
	////////////////////////////////////////
	settingsMarket := map[string]int{
		"piple":   10,
		"tovar":   20,
		"servise": 30,
	}

	var market rod.Market
	market.Init(settingsMarket)
	fmt.Println(market)
	fmt.Println(*market.Fam["tovar"].ListFamile[0])
	fam := market.Fam["tovar"].ListFamile[0]
	rod.PrintProdact(fam)
	market.Torg()
	rod.PrintProdact(fam)
	market.Life(10)
}
