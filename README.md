# go-modules (Package Management)

## In this repository you will learn how to use the go modules in golang.
## Steps you need to follow.

- create a folder and run below cli commands.
- `mkdir folder_name`
- `cd folder_name`

- Now initialize the module file in this folder

- `go mod init folder_name`
 
Now just start working and write code as you desire and if you want to use a function that is present in an other module then do declare the function name with Capital first letter to make it global and then open the module file in which you want to use it and make that whole module require init. for example in our case we want to use the functions present in cities module we wrote a line in country module `replace cities => ../cities` this is how we make the other module available in our current module and give the path of it where it is present.

## How to run this project
- `cd country`
- `go run country.go`
