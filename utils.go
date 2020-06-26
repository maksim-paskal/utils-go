package utils

type Utils struct {
}

func (u *Utils) stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
