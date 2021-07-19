package calculadora

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert" // Se importa testify
)

// se crea un un struct dummyLogger
type dummyLogger struct{}

//  Se escribe las funciones necesarioas para que dummyLogger cumpla con la interfaz que va a reemplazar (Logger)
func (d *dummyLogger) Log(string) error {
	return nil
}

func TestSumar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	resultadoEsperado := 8
	// Se genera el objeto dummy a usar para satisfacer la necesidad de la funcion Sumar
	myDummy := &dummyLogger{}
	// Se ejecuta el test
	resultado := Sumar(num1, num2, myDummy)

	// Se validan los resultados aprovechando testify
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")

}

// se crea un un struct stubLogger
type stubLogger struct{}

//  Se escribe las funciones necesarias para que stubLogger retorne exactamente lo que necesitamos
func (s *stubLogger) Log(string) error {
	return errors.New("error desde stub")
}

func TestSumarError(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	resultadoEsperado := -99999
	// Se genera el objeto stub a usar para satisfacer la necesidad de la funcion Sumar
	myStub := &stubLogger{}
	// Se ejecuta el test
	resultado := Sumar(num1, num2, myStub)

	// Se validan los resultados aprovechando testify
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")

}

// se crea un un struct spy compuesto por un booleano que nos informará cuando ocurre el llamado a Log
type spyLogger struct {
	spyCalled bool
}

//  Se escribe las funciones necesarias para que spy cumpla con la interfaz y ademas espiar cuando es invocado. Para esto seteamos en true spyCalled si entra al metodo
func (s *spyLogger) Log(string) error {
	s.spyCalled = true
	return nil
}

func TestSumarConSpy(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	resultadoEsperado := 8
	// Se genera el objeto spy a usar
	mySpy := &spyLogger{}
	// Se ejecuta el test
	resultado := Sumar(num1, num2, mySpy)
	// Se validan los resultados aprovechando testify
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
	// Se valida que la variable spyCalled sea true para dar el test por valido
	assert.True(t, mySpy.spyCalled)

}

// se crea un un struct mockConfig
type mockConfig struct {
	clienteUsado string
}

// El mock debe implementar el metodo necesatio y comprobar que SumaEnabled sea llamado y que se haga exactamente con el mismo cliente que recibió SumarRestricted
func (m *mockConfig) SumaEnabled(cliente string) bool {
	m.clienteUsado = cliente
	return true
}

func TestSumarRestricted(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	cliente := "John Doe"
	resultadoEsperado := 8
	// Se genera el objeto dummy a usar para satisfacer la necesidad de la funcion Sumar
	myMock := &mockConfig{}
	// Se ejecuta el test
	resultado := SumarRestricted(num1, num2, myMock, cliente)
	// Se validan los resultados aprovechando testify
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
	// Se valida que la informacion obtenida en el mock sea la esperada
	assert.Equal(t, cliente, myMock.clienteUsado)
}

// se crea un un struct fakeConfig
type fakeConfig struct{}

//  Se escribe las funciones necesarioas para que fakeConfig implemente una logica en la que solo habilita la suma al cliente "John Doe"
func (f *fakeConfig) SumaEnabled(cliente string) bool {
	return cliente == "John Doe"
}

func TestSumarRestrictedFake(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	cliente := "John Doe"
	cliente_dos := "Mister Pmosh"
	resultadoEsperado := 8
	resultadoEsperadoError := -99999
	// Se genera el objeto fake a usar
	myFake := &fakeConfig{}
	// Se ejecuta el test
	resultado := SumarRestricted(num1, num2, myFake, cliente)
	// Se validan que para el cliente autorizado devuelva el resultado correcto de la suma
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")

	resultado2 := SumarRestricted(num1, num2, myFake, cliente_dos)
	// Se validan que para el cliente no autorizado devuelva el numero -99999
	assert.Equal(t, resultadoEsperadoError, resultado2, "deben ser iguales")
}
