package slice

func rotarDerecha(original string) string {
	cambiado := string(original[len(original)-1])
	cambiado = cambiado + string(original[:len(original)-1])
	return cambiado
}

func rotarDerechaVeces(original string, veces int) string {
	rotado := original
	for i := 0; i < veces; i++ {
		rotado = rotarDerecha(rotado)
	}
	return rotado
}

func rotarIzquierda(original string) string {
	cambiado := string(original[1:])
	cambiado = cambiado + string(original[0])
	return cambiado
}

func rotarIzquierdaVeces(original string, veces int) string {
	rotado := original
	for i := 0; i < veces; i++ {
		rotado = rotarIzquierda(rotado)
	}
	return rotado
}

func RotateRightVitu(text string) string {
	size := len(text)
	rotatedText := ""
	for fromIndex := range text {
		toIndex := (size - 1 + fromIndex) % size
		char := string(text[toIndex])
		rotatedText = rotatedText + char
	}
	return rotatedText
}
