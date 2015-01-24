package MultiwaySwitch

import (
	"github.com/coopernurse/gorp"
	"math/rand"
	"time"
)

type Drone struct {
	Id        int64 `db:"drone_id"`
	Created   int64
	Token     string
	SecretKey string
}

func randStr(strlen int) string {
	rand.Seed(time.Now().Unix())
	data := make([]byte, strlen)
	var num int
	for i := 0; i < strlen; i++ {
		num = rand.Intn(57) + 65
		for {
			if num > 90 && num < 97 {
				num = rand.Intn(57) + 65
			} else {
				break
			}
		}
		data[i] = byte(num)
	}
	return string(data)
}
func CheckSecretKey(id int64, secretKey string) bool {
	if obj, err := dbmap.Get(Drone{}, id); err == nil {
		drone := obj.(*Drone)
		if drone.SecretKey == secretKey {
			return true
		}
	}
	return false
}

func FlushToken(id int64) (string, error) {
	obj, err := dbmap.Get(Drone{}, id)
	if err != nil {
		return "", err
	}

	drone := obj.(*Drone)
	drone.Token = randStr(16)
	if _, err := dbmap.Update(drone); err != nil {
		return "", err
	}

	return drone.Token, nil
}
