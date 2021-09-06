package utils

import "testing"

func TestIfNoFileToCreate(t *testing.T) {
	file := IfNoFileToCreate("/Users/kwinin/go/src/backend_ft_tcs/data/file")
	t.Log(file)

}
func TestIfNoDirToCreate(t *testing.T) {
	dir, err := IfNoDirToCreate("/Users/kwinin/go/src/backend_ft_tcs/data/file/2/1619689581")
	t.Log(dir, err)
}
