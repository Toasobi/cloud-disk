package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	Id                 int    `json:"id"`
	Identity           string `json:"identity"`
	Name               string `json:"name"`
	jwt.StandardClaims        //必要
}

var myKey = []byte("online-QA-community") //密钥

//用户签发token
func GenerateToken(id int, identity, name string, second int64) (string, error) {
	userClaim := &UserClaims{
		Id:       id,
		Identity: identity,
		Name:     name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(second)).Unix(), //超时
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim) //这个token还不是我们最后想要的字符串token
	tokenString, err := token.SignedString(myKey)                 //注意要传byte数据
	if err != nil {
		return "", err
	}
	//fmt.Println(tokenString) //这样就生成了
	return tokenString, nil
}

//解析token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	//tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiR2V0IiwiaWRlbnRpdHkiOiJ1c2VyXzEifQ.OirWW9EXwML2aj6aqxvZETS3RtO7QVAvs0xI5eek2VI"
	userClaims := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaims, func(t *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err != nil {
		return nil, err
	}
	if !claims.Valid { //正常情况下打印
		//fmt.Println(userClaims)
		return nil, fmt.Errorf("Analyse TOken err:%v", err)
	}
	return userClaims, nil
}
