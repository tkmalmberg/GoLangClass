package main

import (
    "fmt"
)

type Virtmach struct {
    ip string
    hostname string
    diskgb int
    ram int
}

func (v *Virtmach) downloadRam() {
    v.ram++
}

func (v Virtmach) show() {
    fmt.Println("IP Address: ", v.ip)
    fmt.Println("Hostname: ", v.hostname)
    fmt.Println("Disk Size (GB): ", v.diskgb)
    fmt.Println("Memory: ", v.ram)
}

func main() {
    v := Virtmach{"192.127.0.1", "123456", 256, 8}

    v.show()

    v.downloadRam()

    v.show()
}
