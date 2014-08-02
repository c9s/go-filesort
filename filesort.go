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

	sort.Sort(FileMtimeSort{allFiles})
	fmt.Printf("MtimeSort: %+v\n", allFiles)

	sort.Sort(FileMtimeReverseSort{allFiles})
	fmt.Printf("MtimeReverseSort: %+v\n", allFiles)

	sort.Sort(FileSizeReverseSort{allFiles})
	fmt.Printf("MtimeSizeSort: %+v\n", allFiles)
*/

import "os"
import "path/filepath"
import "sort"

type Files []string

func (self Files) Len() int           { return len(self) }
func (self Files) Less(i, j int) bool { return self[i] < self[j] }
func (self Files) Swap(i, j int)      { self[i], self[j] = self[j], self[i] }

type FileSort struct{ Files }

func (self FileSort) Less(i, j int) bool { return self.Files[i] < self.Files[j] }

/*
FileMtimeSort sort files From older to newer.
*/
type FileMtimeSort struct{ Files }

func (self FileMtimeSort) Less(i, j int) bool {
	info_i, _ := os.Stat(self.Files[i])
	info_j, _ := os.Stat(self.Files[j])
	return info_i.ModTime().UnixNano() < info_j.ModTime().UnixNano()
}

/*
From older files to newer files.
*/
type FileMtimeReverseSort struct{ Files }

func (self FileMtimeReverseSort) Less(i, j int) bool {
	info_i, _ := os.Stat(self.Files[i])
	info_j, _ := os.Stat(self.Files[j])
	return info_i.ModTime().UnixNano() > info_j.ModTime().UnixNano()
}

type FileSizeSort struct{ Files }

func (self FileSizeSort) Less(i, j int) bool {
	info_i, _ := os.Stat(self.Files[i])
	info_j, _ := os.Stat(self.Files[j])
	return info_i.Size() < info_j.Size()
}

type FileSizeReverseSort struct{ Files }

func (self FileSizeReverseSort) Less(i, j int) bool {
	info_i, _ := os.Stat(self.Files[i])
	info_j, _ := os.Stat(self.Files[j])
	return info_i.Size() > info_j.Size()
}
