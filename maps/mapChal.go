package main

import "fmt"

func main() {
	var fileExtensions = map[string]string{
		"Python":  ".py",
		"C++":     ".cpp",
		"Java":    ".java",
		"Golang":  ".go",
		"Ansible": ".yml",
	}

	fmt.Println(fileExtensions)

	delete(fileExtensions, "C++")

	fileExtensions["Julia"] = ".jl"

	fmt.Println(fileExtensions)
}
