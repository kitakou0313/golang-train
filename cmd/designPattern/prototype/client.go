package main

/*
ProtoType pattern
あるobjectのコピーをそのクラスに直接依存せず実装可能にするpattern

Prob:
- あるObjectのコピーを作成したいとき...
  - そのObjectが属するクラスのfiealdが全て参照可能なわけではない(private)
  - クラスごとに依存してしまう
  - interfaceでしか参照できない場合，具体的なclassがわからないため↑もできない
*/
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
