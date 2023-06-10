package main

func main() {
	file1 := &File{
		name: "File1",
	}
	file2 := &File{
		name: "File2",
	}
	file3 := &File{
		name: "File3",
	}

	folder1 := &Folder{
		children: []Inode{
			file1,
		},
		name: "Folder1",
	}
	folder2 := &Folder{
		children: []Inode{
			folder1, file2, file3,
		},
		name: "Folder2",
	}

	folder2.print(" ")

	clone := folder2.clone()
	clone.print(" ")
}
