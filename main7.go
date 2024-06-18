package main

import "fmt"

type Player struct {
	HP int
}

func takeDamage(player *Player, amount int) {
	player.HP -= amount
	fmt.Println("PLayer is taking damage, New HP --->", player.HP)

}

func main7() {
	player := Player{
		HP: 200,
	}
	takeDamage(&player, 10)

	fmt.Printf("%+v\n", player)
}
