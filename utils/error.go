package utils

func CheckError(err error) {
	if err != nil {
		panic("An error ocurred")
	}
}
