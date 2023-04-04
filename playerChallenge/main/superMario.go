package main

import (
    "github.com/tkmalmberg/supermario/model"
)

func main() {
    mario := model.Player{3, 1, []string{"mushroom"}}

    mario.Show()

    mario.GreenMushroom()
    mario.Pickup("Star")
    mario.CanWhistle()

    mario.Show()
}
