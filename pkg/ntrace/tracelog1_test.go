package ntrace

func LogFunc1() {
	LogFunc2()
}

func LogFunc2() {
	LogFunc3()
}
