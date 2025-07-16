package node

import "sync"

var Reputation = make(map[string]int)
var mutex sync.Mutex

func UpdateReputation(ip string, score int) {
	mutex.Lock()
	defer mutex.Unlock()
	Reputation[ip] += score
}

func GetReputation(ip string) int {
	mutex.Lock()
	defer mutex.Unlock()
	return Reputation[ip]
}
