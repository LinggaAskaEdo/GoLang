package main

import (
	"fmt"
	"strconv"

	"github.com/fatih/structs"
	"github.com/go-redis/redis"
)

func main() {
	client := newClient()

	errPing := ping(client)

	if errPing != nil {
		fmt.Println(errPing)
	}

	errSet := set(client)

	if errSet != nil {
		fmt.Println(errSet)
	}

	errGet := get(client)

	if errGet != nil {
		fmt.Println(errGet)
	}

	errHmSet := hmSet(client)

	if errHmSet != nil {
		fmt.Println(errHmSet)
	}

	errHmGetAll := hmGetAll(client)

	if errHmGetAll != nil {
		fmt.Println(errHmGetAll)
	}
}

// create connection
func newClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  //use default DB
	})

	return client
}

// ping test connectivity for redis (PONG should be returned)
func ping(client *redis.Client) error {
	pong, err := client.Ping().Result()

	if err != nil {
		return err
	}

	fmt.Println(pong, err) // output: PONG <nil>

	return nil
}

// set executes the redis Set command
func set(client *redis.Client) error {
	err := client.Set("key", "value", 0).Err()

	if err != nil {
		return err
	}

	return nil
}

// get value from redis after set
func get(client *redis.Client) error {
	val, err := client.Get("key").Result()

	if err != nil {
		return err
	}

	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()

	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist

	return nil
}

// User is a simple user struct for this example
type User struct {
	Username  string `json:"username"`
	MobileID  int    `json:"mobile_id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func hmSet(client *redis.Client) error {
	usr := User{
		Username:  "otto",
		MobileID:  1234567890,
		Email:     "ottoM@repoman.com",
		FirstName: "Otto",
		LastName:  "Maddox",
	}

	usrM := structs.Map(usr)

	err := client.HMSet("user:"+usr.Username, usrM).Err()

	if err != nil {
		return err
	}

	return nil
}

func hmGetAll(client *redis.Client) error {
	m, err := client.HGetAll("user:otto").Result()

	if err != nil {
		return err
	}

	usr := User{}

	for key, value := range m {
		switch key {
		case "Username":
			usr.Username = value
		case "MobileID":
			usr.MobileID, err = strconv.Atoi(value)

			if err != nil {
				return err
			}
		case "Email":
			usr.Email = value
		case "FirstName":
			usr.FirstName = value
		case "LastName":
			usr.LastName = value
		}
	}

	fmt.Printf("%+v\n", usr)

	return nil
}
