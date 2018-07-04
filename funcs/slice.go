package funcs

//DeleteInts 删除[]int切片的某个键 并更新键值
func DeleteInts(ints []int, key int) []int {
	tmp := make([]int, len(ints))
	copy(tmp, ints)
	if key == len(tmp)-1 {
		tmp = tmp[:key]
	} else {
		tmp = append(tmp[:key], tmp[key+1:]...)
	}
	return tmp
}
