# Software architecture lab2
## Evenloop

### Task

In the cmd / example / main.go file, you need to implement getting an expression from the parameters
command line, call its function from the option and display the results (or
output error information in the input data).

Reading the input expression has
support or pass the expression itself on the command line,
or specify a file with an expression. 

An optional definition of the file to write to must also be supported
the result of the transformation.

### Flags:
Flag for direct reading from the console:
` -e “+ 5 5” `

Flag for reading from a file:
`-f input.txt`

Flag for writing to a file:
`-o output.txt`

### Instructions
  
Install:
```
git clone https://github.com/Mixa8503/arch_lab_2.git
```
Run:
```
go run ./cmd/example/main.go -e “- 42 1” -o result.txt
```
```
go run ./cmd/example/main.go -f input.txt
```
Test:
```
go test
```

Completed by:
- Hakavyi Oleksandr [@sanix_to](https://t.me/sanix_to)
- Lukianenko Mikhail [@lukianenko78](https://t.me/lukianenko78)
