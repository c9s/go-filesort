package filesort

/*

	var allFiles []string
	filepath.Walk("foo", func(path string, info os.FileInfo, err error) error {
		if err == nil {
			allFiles = append(allFiles, path)
		}
		return err
	})
	// sort.Sort(Files(allFiles))
	fmt.Printf("Original: %+v\n", allFiles)

	sort.Sort(MtimeSort{allFiles})
	fmt.Printf("MtimeSort: %+v\n", allFiles)

	sort.Sort(MtimeReverseSort{allFiles})
	fmt.Printf("MtimeReverseSort: %+v\n", allFiles)

	sort.Sort(SizeReverseSort{allFiles})
	fmt.Printf("MtimeSizeSort: %+v\n", allFiles)
*/
import "os"
import "sort"

type Files []string

func (self Files) Len() int           { return len(self) }
func (self Files) Less(i, j int) bool { return self[i] < self[j] }
func (self Files) Swap(i, j int)      { self[i], self[j] = self[j], self[i] }

func (files Files) Sort() { sort.Sort(files) }

type DefaultSort struct{ Files }

func (self DefaultSort) Less(i, j int) bool { return self.Files[i] < self.Files[j] }

/*
MtimeSort sort files From older to newer.
*/
type MtimeSort struct{ Files }

func (self MtimeSort) Less(i, j int) bool {
	info_i, _ := os.Stat(self.Files[i])
	info_j, _ := os.Stat(self.Files[j])
	return info_i.ModTime().UnixNano() < info_j.ModTime().UnixNano()
}

/*
From older files to newer files.
*/
type MtimeReverseSort struct{ Files }

func (self MtimeReverseSort) Less(i, j int) bool {
	info_i, _ := os.Stat(self.Files[i])
	info_j, _ := os.Stat(self.Files[j])
	return info_i.ModTime().UnixNano() > info_j.ModTime().UnixNano()
}

type SizeSort struct{ Files }

func (self SizeSort) Less(i, j int) bool {
	info_i, _ := os.Stat(self.Files[i])
	info_j, _ := os.Stat(self.Files[j])
	return info_i.Size() < info_j.Size()
}

type SizeReverseSort struct{ Files }

func (self SizeReverseSort) Less(i, j int) bool {
	info_i, _ := os.Stat(self.Files[i])
	info_j, _ := os.Stat(self.Files[j])
	return info_i.Size() > info_j.Size()
}
