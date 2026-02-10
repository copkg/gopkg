package util

import "strconv"

// GetGenderFromIDCard 根据身份证号码判断性别
// 输入: idCard string - 身份证号码字符串
// 输出: gender string - 1-"男" 或 2-"女"; 0 - 号码格式是否有效
func GetGenderFromIDCard(idCard string) int {
	length := len(idCard)
	// 基本格式校验：只处理15位或18位身份证号
	if length != 15 && length != 18 {
		return 0
	}
	var genderDigitStr string
	// 根据身份证长度，确定性别位的位置
	if length == 18 {
		genderDigitStr = idCard[16:17] // 18位身份证，取第17位（索引16）
	} else {
		genderDigitStr = idCard[14:15] // 15位身份证，取第15位（索引14）
	}
	// 将字符串形式的数字转换为整数
	genderDigit, err := strconv.Atoi(genderDigitStr)
	if err != nil {
		return 0
	}
	// 根据奇偶性判断性别
	if genderDigit%2 == 0 {
		return 2
	} else {
		return 1
	}
}
