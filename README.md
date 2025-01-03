# Archer

## 1 Introduction
Archer is a simple just for fun tool to search for packages from the terminal. 

## 2 Dependencies
1. Golang

```shell
$ sudo pacman -S --needed --noconfirm go
```

## 3 Compile

To compile the code, run the following:

```shell
$ gh repo clone f0xb17/Archer
$ cd Archer
$ go run main.go
```

## 3.1 Build

To create an executable file, run the following:

```shell
$ gh repo clone f0xb17/Archer
$ cd Archer
$ go build -o Archer
```

## Usage

Just a simple search:

```shell
$ ./Archer lazygit
```

More advanced search:

```shell
$ ./Archer vim | grep editor
```
