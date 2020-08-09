package rod

import "fmt"

type Market struct {
	Prodacts    []string
	Fam         map[string]*Rod
	CountMakers map[string]int
}

func (m *Market) Init(countMakers map[string]int) {
	/*инициализация рынка, вход - количество торговцев каждым товаром*/
	m.CountMakers = make(map[string]int)
	m.Prodacts = make([]string, 0, 1)
	m.Fam = make(map[string]*Rod)
	for prodact, N := range countMakers {
		m.CountMakers[prodact] = N
		m.Prodacts = append(m.Prodacts, prodact)
		rod := new(Rod)
		rod.Init(N, prodact)
		m.Fam[prodact] = rod
	}
}

func Sdelka(buyers, sallers *Famile) {
	nameProdact := sallers.name_prodact
	var deltaMoney, deltaProdact float64
	///////////////////////////////////////////////////////////////////////////
	p, ok := sallers.prob_prod_money[nameProdact][buyers.id]
	if !ok {
		p = randFl(0.0, 0.05)
		sallers.prob_prod_money[nameProdact][buyers.id] = p
	}
	deltaMoney = p * buyers.money
	sallers.money += deltaMoney
	buyers.money -= deltaMoney
	buyers.money_spent[nameProdact][sallers.id] = deltaMoney
	sallers.money_sold[nameProdact][buyers.id] = deltaMoney
	/////////////////////////////////////////////////////////////////////////
	t, ok2 := sallers.prob_prod[nameProdact][buyers.id]
	if !ok2 {
		t = randFl(0.0, 0.05)
		sallers.prob_prod[nameProdact][buyers.id] = t
	}
	deltaProdact = t * sallers.Prodact[nameProdact]
	sallers.Prodact[nameProdact] -= deltaProdact
	buyers.make_prodact[nameProdact] += deltaProdact
	buyers.sellers_prod[nameProdact][sallers.id] = deltaProdact
	sallers.buyers_prod[nameProdact][buyers.id] = deltaProdact
	///////////////////////////////////////////////////////////////////
	price := deltaMoney / deltaProdact
	buyers.sellers_price[nameProdact][sallers.id] = price
	sallers.buyers_price[nameProdact][buyers.id] = price
}

func (m *Market) Torg() {
	//fmt.Println("Start torg")
	for pr, listBuyers := range m.Fam {
		for _, nameProdact := range m.Prodacts {
			//fmt.Println(nameProdact)
			if pr == nameProdact {
				continue
			}
			listSellers := m.Fam[nameProdact]
			for _, buyer := range listBuyers.ListFamile {
				for _, seller := range listSellers.ListFamile {
					Sdelka(buyer, seller)
				}
			}
		}
	}
	fmt.Println("Finish torg")
}

func (m *Market) Life(countIterations int) {
	var allMoney float64
	var tovar float64
	for i := 0; i < countIterations; i++ {
		m.Torg()
		for prodact, objRod := range m.Fam {
			for _, famile := range objRod.ListFamile {
				famile.ChangeProbabilites()
				famile.Make_pr(prodact)
				allMoney += famile.money
				tovar += famile.Prodact[prodact]
				/*
					if k == 0 {
						fmt.Println(prodact)
						PrintProdact(famile)
					}
				*/
			}
			fmt.Printf("ALL %s: %v\n", prodact, tovar)
			tovar = 0.0
		}
		fmt.Println("ALL MONEY: ", allMoney)
		allMoney = 0.0
	}
}
