package main

type Inode interface {
	clone() Inode
	print(string)
}
