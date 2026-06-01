package util

func ArrayInGroupsOf(arr []string, num int64) [][]string {
	maxNum := int64(len(arr))
	//判断数组大小是否小于等于指定分割大小的值，是则把原数组放入二维数组返回
	if maxNum <= num {
		return [][]string{arr}
	}
	//获取应该数组分割为多少份
	var quantity int64
	if maxNum%num == 0 {
		quantity = maxNum / num
	} else {
		quantity = (maxNum / num) + 1
	}
	//声明分割好的二维数组
	var segments = make([][]string, 0)
	//声明分割数组的截止下标
	var start, end, i int64
	for i = 1; i <= quantity; i++ {
		end = i * num
		if i != quantity {
			segments = append(segments, arr[start:end])
		} else {
			segments = append(segments, arr[start:])
		}
		start = i * num
	}
	return segments
}
