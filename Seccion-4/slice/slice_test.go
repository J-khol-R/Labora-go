package slice

import "testing"

//crear un test y una funcion que corra a la derecha y a la izquierda una cadena tambien elegir cuantas veces correrla

type rotarTest struct {
	original, expected string
}

type rotarVecesTest struct {
	original, expected string
	veces              int
}

var rotarDerechaTests = []rotarTest{
	{"hola", "ahol"},
	{"adios", "sadio"},
	{"valentina", "avalentin"},
	{"juan", "njua"},
}

var rotarIzquierdaTests = []rotarTest{
	{"hola", "olah"},
	{"adios", "diosa"},
	{"valentina", "alentinav"},
	{"juan", "uanj"},
}

var rotarDerechaVecesTests = []rotarVecesTest{
	{"hola", "laho", 2},           //2
	{"adios", "iosad", 3},         //3
	{"valentina", "tinavalen", 4}, //4
	{"juan", "njua", 5},           //5
}

var rotarIzquierdaVecesTests = []rotarVecesTest{
	{"hola", "laho", 2},           //2
	{"adios", "osadi", 3},         //3
	{"valentina", "ntinavale", 4}, //4
	{"juan", "uanj", 5},           //5
}

func TestRotarDerecha(t *testing.T) {

	for _, test := range rotarDerechaTests {
		if output := rotarDerecha(test.original); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
func TestRotarIzquierda(t *testing.T) {

	for _, test := range rotarIzquierdaTests {
		if output := rotarIzquierda(test.original); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
func TestRotarDerechaVeces(t *testing.T) {

	for _, test := range rotarDerechaVecesTests {
		if output := rotarDerechaVeces(test.original, test.veces); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
func TestRotarIzquierdaVeces(t *testing.T) {

	for _, test := range rotarIzquierdaVecesTests {
		if output := rotarIzquierdaVeces(test.original, test.veces); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
