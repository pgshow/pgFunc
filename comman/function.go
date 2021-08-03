package comman

import (
	"io/ioutil"
	"regexp"
	"strings"
)

// GetListByTxt 从文件获取为切片
// file:文件名
// splitStr: 分割符
func GetListByTxt(file string, splitStr string) []string {
	data, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), splitStr)
}

// StrContainAny 判断目标字符串是否包含在列表里面
func StrContainAny(str string, list []string) bool {
	for _, n := range list {
		if strings.Contains(str, n) {
			return true
		}
	}
	return false
}

// StrInAny 判断目标字符串是否在列表里面，完全相等
func StrInAny(str string, list []string) bool {
	for _, n := range list {
		if n == str {
			return true
		}
	}
	return false
}

// DeleteExtraSpace 删除字符串中的多余空格，有多个空格时，仅保留一个空格
func DeleteExtraSpace(s string) string {
	s1 := strings.Replace(s, "\n", " ", -1)      //替换tab为空格
	reg, _ := regexp.Compile(`\s{2,}`)             //两个及两个以上空格的正则表达式
	s2 := make([]byte, len(s1))                  //定义字符数组切片
	copy(s2, s1)                                 //将字符串复制到切片
	spc_index := reg.FindStringIndex(string(s2)) //在字符串中搜索
	for len(spc_index) > 0 {                     //找到适配项
		s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...) //删除多余空格
		spc_index = reg.FindStringIndex(string(s2))            //继续在字符串中搜索
	}
	return string(s2)
}
