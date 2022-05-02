package services

func TrataErro(err error) {
	if err != nil {
		panic(err.Error())
	}
}
