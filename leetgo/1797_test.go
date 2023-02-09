package main

import (
	"fmt"
	"testing"
)

func Test1797(t *testing.T) {

}

type AuthenticationManager struct {
	TimeToLive     int
	TokenExpireMap map[string]int
}

func Constructor1797(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		TimeToLive:     timeToLive,
		TokenExpireMap: map[string]int{},
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.TokenExpireMap[tokenId] = currentTime + this.TimeToLive
	fmt.Println(this.TokenExpireMap)
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if expire, ok := this.TokenExpireMap[tokenId]; !ok || expire < currentTime {
		fmt.Println(this.TokenExpireMap)
		return
	}
	this.TokenExpireMap[tokenId] = currentTime + this.TimeToLive
	fmt.Println(this.TokenExpireMap)
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	var cnt int
	for _, expire := range this.TokenExpireMap {
		if expire > currentTime {
			cnt++
		}
	}
	return cnt
}
