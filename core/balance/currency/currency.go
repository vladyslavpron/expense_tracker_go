package currency

type Currency string

const (
	USD Currency = "USD"
	EUR Currency = "EUR"
	UAH Currency = "UAH"
)

var Currencies = []Currency{USD, EUR, UAH}

func (Currency) Values() (values []string) {
	for _, v := range Currencies {
		values = append(values, string(v))
	}
	return
}
