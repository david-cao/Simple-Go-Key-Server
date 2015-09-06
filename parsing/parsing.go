package parsing

import (
	"errors"
	"fmt"

	"github.com/dcao96/guess4/database"

	"github.com/garyburd/redigo/redis"
)

func AddUser(fbID string, publicKey string) error {

	conn := database.Get()
	defer conn.Close()

	//validation checks here
	exists, _ := redis.Bool(conn.Do("EXISTS", fbID))
	if exists {
		return errors.New("Facebook ID is already registered.")
	}

	fmt.Println("SET", fbID, publicKey)
	conn.Do("SET", fbID, publicKey)

	return nil
}

func GetPublicKey(fbID string) (publicKey string, err error) {

	conn := database.Get()
	defer conn.Close()

	key, err := redis.String(conn.Do("GET", fbID))
	if err != nil {
		return "", err
	} else {
		return key, nil
	}

}
