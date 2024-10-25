package jwt

import (
	"github.com/golang-jwt/jwt/v4"
)

/**
$payload = [
"iss" => "", //签发者
"aud" => "", //面向的用户
"iat" => time(), //签发时间
"nbf" => time(), //在什么时候jwt开始生效
"exp" => time() + 86400, //token 过期时间 秒
'xx' => 'xxx'  //自定义
];
*/

func GenerateToken(secretKey string, expire int64, now int64, userId string, roleId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["iat"] = now
	claims["exp"] = now + expire
	claims["uid"] = userId
	claims["roleId"] = roleId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
