package main

import (
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
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if expire, ok := this.TokenExpireMap[tokenId]; !ok || expire <= currentTime {
		delete(this.TokenExpireMap, tokenId)
		return
	}
	this.TokenExpireMap[tokenId] = currentTime + this.TimeToLive
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	var cnt int
	for tokenId, expire := range this.TokenExpireMap {
		if expire > currentTime {
			cnt++
		} else {
			delete(this.TokenExpireMap, tokenId)
		}
	}
	return cnt
}
