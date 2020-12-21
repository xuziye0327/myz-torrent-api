package common

import (
	"fmt"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	var files Files

	for i := 5; i >= 3; i-- {
		files = append(files, &File{Name: fmt.Sprintf("test%v", i), isDir: false})
	}

	for i := 2; i >= 0; i-- {
		files = append(files, &File{Name: fmt.Sprintf("test%v", i), isDir: true})
	}

	sort.Sort(files)

	for i := 0; i < 3; i++ {
		if !files[i].isDir {
			t.Errorf("%v is not dir", files[i])
		}
		var name = fmt.Sprintf("test%v", i)
		if files[i].Name != name {
			t.Errorf("name of %v is not  %v", files[i].Name, name)
		}
	}

	for i := 3; i < 6; i++ {
		if files[i].isDir {
			t.Errorf("%v is not file", files[i])
		}
		var name = fmt.Sprintf("test%v", i)
		if files[i].Name != name {
			t.Errorf("name of %v is not  %v", files[i].Name, name)
		}
	}
}
