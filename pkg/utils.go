package pkg

import "strings"

func Judge(s string) bool {
	if strings.Compare(s, "自然科学") != 0 && strings.Compare(s, "电脑网络") != 0 && strings.Compare(s, "外国文学") != 0 && strings.Compare(s, "诗歌散文") != 0 {
		return true
	} else {
		return false
	}
}
