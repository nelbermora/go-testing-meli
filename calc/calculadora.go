package calculadora

type Logger interface {
	Log(string) error
}

// Funcion que recibe dos enteros y retorna la suma resultante
func Sumar(num1, num2 int, logger Logger) int {
	err := logger.Log("Ingreso a Funcion Sumar")
	if err != nil {
		return -99999
	}
	return num1 + num2
}

type Config interface {
	SumaEnabled(cliente string) bool
}

// Funcion que recibe dos enteros y retorna la suma resultante
func SumarRestricted(num1, num2 int, config Config, cliente string) int {
	if !config.SumaEnabled(cliente) {
		return -99999
	}
	return num1 + num2

}
