package rod

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/segmentio/ksuid"
)

func randFl(min float64, max float64) float64 {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Float64()*(max-min)
}
func sumSlice(a []float64) float64 {
	var result float64
	for _, v := range a {
		result += v
	}
	return result
}

func sumMap(m map[string]float64) float64 {
	var result float64
	for _, v := range m {
		result += v
	}
	return result
}

const COUNT_CHAPTERS = 3

var ARRAY_NAME_PRODUCT = [COUNT_CHAPTERS]string{"piple", "tovar", "servise"}

type Famile struct {
	id           string
	name_prodact string
	money        float64   // сколько денег у рода
	psiho        float64   // психологическое состояние рода
	old_money    []float64 // история накоплений
	//sena float64  // цена продажи умений рода

	buyers_emo   map[string]map[string]float64 // емоции, которуе род испытывает к покупателям, зависят от цены
	buyers_price map[string]map[string]float64 // цена проданного товара
	buyers_prod  map[string]map[string]float64 // количество проданных умений
	prob_prod    map[string]map[string]float64 // вероятность передачи товара i-тому покупателю товара string
	money_sold   map[string]map[string]float64 // количество денег, вырученные за товар string

	sellers_emo     map[string]map[string]float64 // емоции, которуе род испытывает к продавцам, зависят от цены
	sellers_price   map[string]map[string]float64 // цена купленного товара
	sellers_prod    map[string]map[string]float64 // количество закупленных товаров
	money_spent     map[string]map[string]float64 // количество денег, потраченных на каждый товар
	prob_prod_money map[string]map[string]float64 // вероятность передачи денег i-тому производителю товара string

	make_prodact map[string]float64 // количество товаров для производства, закупленных у разных продавцов

	Prodact map[string]float64 // сколько умений у рода на продажу

}

func PrintProdact(f *Famile) {
	fmt.Println("Print Prodact: ", f.id)
	for k, v := range f.Prodact {
		fmt.Println(k, ": ", v)
	}
	fmt.Println("money: ", f.money)
}
func (f *Famile) Init(nm_prod string) {
	f.id = ksuid.New().String()
	f.name_prodact = nm_prod
	f.money = randFl(0.0, 10000.0)
	f.psiho = 0.01

	f.sellers_emo = make(map[string]map[string]float64)
	f.sellers_prod = make(map[string]map[string]float64)
	f.money_spent = make(map[string]map[string]float64)
	f.buyers_emo = make(map[string]map[string]float64)
	f.buyers_prod = make(map[string]map[string]float64)
	f.money_sold = make(map[string]map[string]float64)
	f.prob_prod = make(map[string]map[string]float64)
	f.prob_prod_money = make(map[string]map[string]float64)
	f.buyers_price = make(map[string]map[string]float64)
	f.sellers_price = make(map[string]map[string]float64)
	f.Prodact = make(map[string]float64)
	f.make_prodact = make(map[string]float64)

	for _, nm := range ARRAY_NAME_PRODUCT {
		f.sellers_emo[nm] = make(map[string]float64)
		f.sellers_prod[nm] = make(map[string]float64)
		f.make_prodact[nm] = 0.0
		f.Prodact[nm] = 0.0
		f.money_spent[nm] = make(map[string]float64)
		f.buyers_emo[nm] = make(map[string]float64)
		f.buyers_prod[nm] = make(map[string]float64)
		f.money_sold[nm] = make(map[string]float64)
		f.prob_prod[nm] = make(map[string]float64)
		f.prob_prod_money[nm] = make(map[string]float64)
		f.buyers_price[nm] = make(map[string]float64)
		f.sellers_price[nm] = make(map[string]float64)
		f.Prodact[nm] = randFl(0.0, 1000.0)

	}
}
func (f *Famile) prob_normalization() {
	var sm float64
	for _, nm := range ARRAY_NAME_PRODUCT {
		sm = sumMap(f.prob_prod_money[nm])
		if sm > 0 {
			for k, _ := range f.prob_prod_money {
				f.prob_prod_money[nm][k] = f.prob_prod_money[nm][k] / sm
			}
		} else {
			for k, _ := range f.prob_prod_money {
				f.prob_prod_money[nm][k] = 0.0
			}
		}
		sm = sumMap(f.prob_prod[nm])
		if sm > 0 {
			for k, _ := range f.prob_prod {
				f.prob_prod[nm][k] = f.prob_prod[nm][k] / sm
			}
		} else {
			for k, _ := range f.prob_prod {
				f.prob_prod[nm][k] = 0.0
			}
		}
	}
}
func (f *Famile) String() string {
	return f.id
}

func (f *Famile) Make_pr(name_prodact string) {
	/*производим товар на продажу, формула Коба-Дугласа) */
	var result, sm float64
	var stp = make(map[string]float64)
	for nm_prod, slc_prod := range f.make_prodact {
		sm = slc_prod
		stp[nm_prod] = sm
		result += sm
	}
	for _, nm := range ARRAY_NAME_PRODUCT {
		stp[nm] = stp[nm] / result
	}
	result = math.E
	for nm_prod, slc_prod := range f.make_prodact {
		sm = slc_prod
		result *= math.Pow(sm, stp[nm_prod])
	}
	f.Prodact[name_prodact] += result
}
func (r *Famile) ChangeProbabilites() {

	for _, nm_tov := range ARRAY_NAME_PRODUCT {
		//buyers
		sm := sumMap(r.buyers_price[nm_tov])
		for i, _ := range r.buyers_emo[nm_tov] {
			r.buyers_emo[nm_tov][i] = (r.buyers_price[nm_tov][i] / sm) - 0.5
		}
		for i := range r.prob_prod[nm_tov] {
			r.prob_prod[nm_tov][i] += r.buyers_emo[nm_tov][i]
			if r.prob_prod[nm_tov][i] < 0.0 {
				r.prob_prod[nm_tov][i] = 0.0
			}
		}
		//sellers
		sm = sumMap(r.sellers_price[nm_tov])
		for i, _ := range r.sellers_emo[nm_tov] {
			r.sellers_emo[nm_tov][i] = (r.sellers_price[nm_tov][i] / sm) - 0.5
		}
		for i, _ := range r.prob_prod_money[nm_tov] {
			r.prob_prod_money[nm_tov][i] -= r.sellers_emo[nm_tov][i]
			if r.prob_prod_money[nm_tov][i] < 0.0 {
				r.prob_prod_money[nm_tov][i] = 0.0
			}
		}
		r.prob_normalization()

	}

}

/*
func main() {
	r1 := new(Rod)
	m1 := map[string]int{"piple": 10, "tovar": 10, "servise": 10}
	m2 := map[string]int{"piple": 20, "tovar": 20, "servise": 20}
	r1.Init("tovar", m1, m2)
	fmt.Printf("%s", "Comon Rod!")
}
*/
