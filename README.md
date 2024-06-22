# JT's Inventory!

This project was initially created to be used for my employer. Fortunately, the project is a
generic platform for a base layer of an inventory system. The project is made in Golang
using the [Charm](https://charm.sh/) suite of tools. These tools work well in creating a terminal or command line 
application. Data is stored in a SQLite database that is created on the initial start of the project. 
By default each inventory item is saved in a table called "inventory". Each product in the
 inventory has an id, name, quantity, creation time, and the time for the most recent update.


# Quick Start

If you are on Linux or have Make installed then this will be your easiest option.
```
make run
```

If you do not have Make then you will need to enter these simple commands.
```
go mod tidy
go build ./cmd
./cmd.exe
``` 
# Commands
Each command will be entered using the command flag. The flags name and quantity are optional
depending on the command.
#### new
Create new product to your inventory
```
./cmd.exe -command=new -name=yourProduct -quantity=yourCurrentQuantity
```
#### add
Add to existing quantity
```
./cmd.exe -command=add -name=yourProduct -quantity=quantityToAdd
```
#### sub
Remove from existing quantity
```
./cmd.exe -command=sub -name=yourProduct -quantity=yourCurrentQuantity
```
#### del
Delete product from table by name
```
./cmd.exe -command=del -name=yourProduct
```
#### view
View all products if no name is entered. Otherwise, view single products data.
```
./cmd.exe -command=view -name=yourProduct
```
