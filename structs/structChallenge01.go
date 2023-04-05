package main

import "fmt"

type Astro struct {
    name string
    age int
    mission string
    isNeeded bool
}

type nasaMission struct {
    people []Astro
    number int
    message string
}
func main() {

    a1 := Astro{"Tyler", 26, "make breakfast", true}
    a2 := Astro{"Megan", 26, "crochette", true}
    a3 := Astro{"Ghibli", 7, "stink", false}
    a4 := Astro{"Puddles", 7, "meow for food", true}

    s := make([]Astro, 0)

    s = append(s, a1, a2, a3, a4)

    fmt.Println(s)
    fmt.Println(s[3].mission)

    nm := nasaMission{s, 4, "success"}

    fmt.Println(nm)
}
