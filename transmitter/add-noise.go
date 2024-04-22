package transmitter

import (
	"fmt"
	"math/rand/v2"
)

func AddNoise(s []byte) []byte {
	n := len(s)

	for i:=0; i<n; i++{
		random_no := rand.IntN(100)
		if random_no>=95{
			s[i] = 9
			fmt.Println(s) 
		}
	}
	
	return s
}
