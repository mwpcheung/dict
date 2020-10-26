package dict

func ReadContent(path string) []byte {
	fs := Open(path)
	defer fs.Close()

	return fs.ReadAll()
}

func WriteContent(path string, content []byte) int {
	fs := Create(path)
	defer fs.Close()

	return fs.Write(content)
}
