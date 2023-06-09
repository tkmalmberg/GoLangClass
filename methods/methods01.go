/* Alta3 Research | RZFeeser
   Methods - defining         */

// Go program to illustrate the
// method with struct type receiver
package main
 
import "fmt"
 
// Author structure
type author struct {
    name      string
    branch    string
    books     int
    salary    int
}
 
// Method with a receiver
// of author type
func (a author) show() {
 
    fmt.Println("Author's Name: ", a.name)
    fmt.Println("Branch Name: ", a.branch)
    fmt.Println("Published articles: ", a.books)
    fmt.Println("Salary: ", a.salary)
}

func (a *author) bookup() {
    a.books++
}

// Main function
func main() {
 
    // Initializing the values
    // of the author structure
    res := author{
        name:      "RZFeeser",
        branch:    "Pennsylvania",
        books:     14,
        salary:    34000,
    }
    
    res.show()
    // Add a book
    res.bookup()
    // Calling the method
    res.show()
}                    
