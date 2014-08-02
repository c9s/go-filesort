
```go
import filesort "github.com/c9s/go-filesort"

var allFiles []string
filepath.Walk("foo", func(path string, info os.FileInfo, err error) error {
    if err == nil {
        allFiles = append(allFiles, path)
    }
    return err
})

fmt.Printf("Original: %+v\n", allFiles)

sort.Sort(filesort.FileSort{allFiles})
fmt.Printf("Default Sort: %+v\n", allFiles)

sort.Sort(filesort.FileMtimeSort{allFiles})

fmt.Printf("MtimeSort: %+v\n", allFiles)

sort.Sort(filesort.FileMtimeReverseSort{allFiles})

fmt.Printf("MtimeReverseSort: %+v\n", allFiles)

sort.Sort(filesort.FileSizeReverseSort{allFiles})

fmt.Printf("MtimeSizeSort: %+v\n", allFiles)
```

Or you can do:

```go
var files Files = Files{ "file1", "file2", "file3" }
files.Sort()
```
