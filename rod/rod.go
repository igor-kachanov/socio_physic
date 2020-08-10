package rod

//Rod - список фамилий и память (список взаимодействий) этих фамилий
type Rod struct {
	MemFamile  []Famile
	ListFamile []*Famile
}

//Init - инициализируем Rod
func (r *Rod) Init(countFamile int, nameProdact string) {
	/*
	 countFamile - количество фамилий в Роду
	 nameProdact - товар, который производит Род (все его фамилии)
	*/
	r.MemFamile = make([]Famile, 0, 1)
	r.ListFamile = make([]*Famile, 0, 1)
	//var temp Famile
	for i := 0; i < countFamile; i++ {
		temp := new(Famile)
		temp.Init(nameProdact)
		r.ListFamile = append(r.ListFamile, temp)
	}
}

//WriteMemory - Род помнит все свои состояния
func (r *Rod) WriteMemory(f Famile) {
	var temp Famile
	temp.Init(f.name_prodact)
	r.MemFamile = append(r.MemFamile, f.GetDeepCopy())

}

//GetDeepCopy Получаем глубокую копию
func (f Famile) GetDeepCopy() Famile {
	/*  */
	var temp Famile
	temp.Init(f.name_prodact)
	temp.id = f.id
	temp.name_prodact = f.name_prodact
	temp.money = f.money
	temp.psiho = f.psiho

	temp.buyers_emo = copyMapMap(f.buyers_emo)     //map[string]map[string]float64  емоции, которуе род испытывает к покупателям, зависят от цены
	temp.buyers_price = copyMapMap(f.buyers_price) //map[string]map[string]float64 // цена проданного товара
	temp.buyers_prod = copyMapMap(f.buyers_prod)   //map[string]map[string]float64 // количество проданных умений
	temp.prob_prod = copyMapMap(f.prob_prod)       //map[string]map[string]float64 // вероятность передачи товара i-тому покупателю товара string
	temp.money_sold = copyMapMap(f.money_sold)     //map[string]map[string]float64 // количество денег, вырученные за товар string

	temp.sellers_emo = copyMapMap(f.sellers_emo)         //map[string]map[string]float64 // емоции, которуе род испытывает к продавцам, зависят от цены
	temp.sellers_price = copyMapMap(f.sellers_price)     //map[string]map[string]float64 // цена пкупленного товара
	temp.sellers_prod = copyMapMap(f.sellers_prod)       //map[string]map[string]float64 // количество закупленных товаров
	temp.money_spent = copyMapMap(f.money_spent)         //map[string]map[string]float64 // количество денег, потраченных на каждый товар
	temp.prob_prod_money = copyMapMap(f.prob_prod_money) //map[string]map[string]float64 // вероятность передачи денег i-тому производителю товара string

	temp.make_prodact = copyMap(f.make_prodact) //map[string]map[string]float64 // количество товаров для производства, закупленных у разных продавцов

	temp.Prodact = copyMap(f.Prodact) //map[string]float64 // сколько умений у рода на продажу

	return temp
}

func copyMap(in map[string]float64) map[string]float64 {
	temp := make(map[string]float64)
	for k, v := range in {
		temp[k] = v
	}
	return temp
}

func copyMapMap(in map[string]map[string]float64) map[string]map[string]float64 {
	temp := make(map[string]map[string]float64)
	for k, v := range in {
		temp[k] = copyMap(v)
	}
	return temp

}
