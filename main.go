package main

func main() {

	s := &Song{}

	files := []string{
		"top50.csv",
	}

	s.readList(files)
}
