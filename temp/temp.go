// declarar meu pacote principal
package main

//importar função fmt
import "fmt"

// declaração da variável CONST do ponto de ebulição da água em K
const ebulicaoK = 373

// função principal
func main() {

	tempFK := 273
	tempFC := 0
	tempK := ebulicaoK       // variável do valor da temperatura em graus K
	tempC := ebulicaoK - 273 //Conversão de K para C
	//apareça na tela o resultado
	fmt.Printf("A temperatura de ebulição da água em °K = %v (%T), temperatura de ebulição da água em °C =%v (%T) .", tempK, tempK, tempC, tempC)
	fmt.Printf("A temperatura de fusão da água em °K = %v (%T), e a temperatura de fusão da água em °C é = %v (%T). ", tempFK, tempFK, tempFC, tempFC)
	// A gente espera que apareça k = 273 e C = 100

}
