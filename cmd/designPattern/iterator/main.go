package main

import "fmt"

func main() {

	user1 := &User{
		name: "a",
		age:  30,
	}
	user2 := &User{
		name: "b",
		age:  20,
	}

	userCollection := &UserCollection{
		users: []*User{user1, user2},
	}

	iterator := userCollection.createIterator()

	// GolangはWhileの代わりにforを用いる 下の意味は while (u.hasNext())と同じ
	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Println(user.name, user.age)
	}
}
