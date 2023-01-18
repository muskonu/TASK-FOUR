package pkg

import (
	"bufio"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func Judge(s string) bool {
	var class = "工具书,期刊杂志,生活休闲,保健养生,电脑网络,教育学习,自然科学,儿童文学,哲学宗教,战争军事,历史地理,励志成功,经营管理,金融投资,外国文学,人物传记,人文社科,推理惊悚,科幻玄幻,职场小说,言情小说,网络小说,武侠小说,古典文学,诗歌散文"
	return strings.Contains(class, s)
}

func ReadNumber() int {
	f, _ := os.Open("number")
	reader := bufio.NewReader(f)
	number, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	n, _ := strconv.Atoi(string(number))
	return n
}

func WriteNumber(new int) {
	f, _ := os.OpenFile("number", os.O_RDWR, 0666)
	_, err := f.WriteAt([]byte(strconv.Itoa(new)), 0)
	if err != nil {
		fmt.Println(err)
	}
}

func Token(member Member) (string, error) { //生成token
	claims := jwt.StandardClaims{
		Audience:  member.Account,                          // 受众
		ExpiresAt: time.Now().Unix() + int64(time.Hour*24), // 失效时间
		Id:        strconv.Itoa(member.MemberId),           // 编号
		IssuedAt:  time.Now().Unix(),                       // 签发时间
		Issuer:    "muksonu",                               // 签发人
		NotBefore: time.Now().Unix(),                       // 生效时间
		Subject:   "login",                                 // 主题
	}
	var jwtSecret = []byte("abaaba")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(jwtSecret)
	ss = "Bearer " + ss
	return ss, err
}

func CheckTiker() { //若超时未还，等级归零
	ticker := time.NewTicker(6 * time.Hour)
	for {
		db := CreateDb()
		db.Exec("call ticker()")
		<-ticker.C
	}
}
