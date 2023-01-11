package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Judge(s string) bool {
	var class = "工具书,期刊杂志,生活休闲,保健养生,电脑网络,教育学习,自然科学,儿童文学,哲学宗教,战争军事,历史地理,励志成功,经营管理,金融投资,外国文学,人物传记,人文社科,推理惊悚,科幻玄幻,职场小说,言情小说,网络小说,武侠小说,古典文学"
	return !strings.Contains(class, s)
}

func ReadNumber() int {
	f, _ := os.Open("number")
	reader := bufio.NewReader(f)
	num, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(num)
	fmt.Println(n)
	return n
}

func WriteNumber(new int) {
	f, _ := os.OpenFile("number", os.O_RDWR, 0666)
	n, err := f.WriteAt([]byte(strconv.Itoa(new)), 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
