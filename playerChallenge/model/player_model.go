package model

import (
    "fmt"
)

type Player struct {
    Lives int
    Stage int
    Inventory []string
}

func (p *Player) GreenMushroom() {
    p.Lives++
}

func (p *Player) Pickup(i string) {
    p.Inventory = append(p.Inventory, i)
}

func (p Player) CanWhistle() bool {
    if p.Stage >= 5 {
        return true
    }

    return false
}

func (p Player) Show() {
    fmt.Println("Lives: ", p.Lives)
    fmt.Println("Stage: ", p.Stage)
    fmt.Println("Inventory: ", p.Inventory)
}
