package lib

func CheckFuncNameExist(name string, list []string) int {
	for i, n := range list {
		if n == name {
			return i
		}
	}
	return -1
}
