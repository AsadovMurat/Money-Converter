package application

import "testing"

func TestConvert(t *testing.T) {

	// Первый тестик
	t.Run("first", func(t *testing.T) {
		var (
			from_currency   string  = "RUB"
			to_currency     string  = "USD"
			value           float64 = 358
			expected_result float64 = 3.88
		)

		real_result, err := Convert(from_currency, to_currency, value)
		if err != nil {
			t.Errorf("В ходе выполнения теста произошла такая ошибка %v", err)
		}

		if real_result != expected_result {
			t.Error(real_result, " != ", expected_result)
		}
	})
	//Второй тестик
	t.Run("second", func(t *testing.T) {
		var (
			from_currency   string  = "EUR"
			to_currency     string  = "USD"
			value           float64 = 10.43
			expected_result float64 = 11.47
		)

		real_result, err := Convert(from_currency, to_currency, value)
		if err != nil {
			t.Errorf("В ходе выполнения теста произошла такая ошибка %v", err)
		}

		if real_result != expected_result {
			t.Error(real_result, " != ", expected_result)
		}
	})

}
