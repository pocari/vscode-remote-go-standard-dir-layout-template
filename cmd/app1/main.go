package main

import (
	"fmt"
	"io"
	"my-project/internal/module1"
	"my-project/pkg/calc"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	output := fmt.Sprintf("internal_module1: %s, cal.Add: %d\n", module1.Module1Func("arg1"), calc.Add(2, 3))
	io.WriteString(w, output)
}

func main() {
	portNumber := "9080"
	http.HandleFunc("/", handle)
	fmt.Println("Server listening on port ", portNumber)
	http.ListenAndServe(":"+portNumber, nil)
}
