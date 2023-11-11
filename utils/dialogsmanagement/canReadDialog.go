package dialogsmanagement

func CanReadDialog(id int, listDialogs []int) bool {
	for _, v := range listDialogs {
		if v == id {
			return false
		}
	}
	return true
}
