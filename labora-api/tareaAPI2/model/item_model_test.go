package model

import (
	"fmt"
	"testing"
	"time"
)

type itemTest struct {
	items    Items
	expected int
}

var itemTests = []itemTest{
	{items: Items{1, "valentina", time.Now(), "naranja", 20, 80, 0}, expected: 1600},
	{items: Items{1, "valentina", time.Now(), "naranja", 10, 9, 0}, expected: 90},
	{items: Items{1, "valentina", time.Now(), "naranja", 76, 20, 0}, expected: 1520},
	{items: Items{1, "valentina", time.Now(), "naranja", 12, 7, 0}, expected: 84},
	{items: Items{1, "valentina", time.Now(), "naranja", 63, 43, 0}, expected: 2709},
	{items: Items{1, "valentina", time.Now(), "naranja", 1, 0, 0}, expected: 0},
}

func TestCalcular(t *testing.T) {
	for _, tt := range itemTests {
		t.Run(fmt.Sprintf("items: %+v, expected: %v", tt.items, tt.expected), func(t *testing.T) {
			resultado := tt.items.CalcularPrecio()
			if resultado != tt.expected {
				t.Errorf("Resultado esperado: %v, pero se obtuvo %v", tt.expected, resultado)
			}
		})
	}
}
