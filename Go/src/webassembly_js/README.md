### Go 1.1 : Web assembly

This file allows to test web assembly integration with go.
This example allows to display one element in console, to add contain in a html div and show js alert.


##### Documentation

Web assembly go documentation : https://github.com/golang/go/wiki/WebAssembly

##### Tools used : 
    - goexec : to execute go code (https://github.com/shurcooL/goexec#goexec)

##### To execute this file :
    - install goexec
    - copy wasm exec js file on your repository : cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
    - compile go program : GOOS=js GOARCH=wasm go build -o main.wasm
    - launch http server : goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
    
    - Navigate on http://localhost:8080/ and look the development console.
    - Navigate on http://localhost:8080/wasm_exec.html and click on the development console.
    
    - Execute go program on commande line : GOOS=js GOARCH=wasm go run -exec="$(go env GOROOT)/misc/wasm/go_js_wasm_exec" .
            