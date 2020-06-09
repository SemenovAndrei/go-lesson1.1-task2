package main

func main() {
	balance := 15_000_000_00  // 15 миллионов в копейках
	transfer := 10_000_000_00 // 10 миллионов в копейках
	total := balance + transfer // int32 + int32 будет int32
	println(total)
}
