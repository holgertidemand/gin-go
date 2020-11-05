package main
import "testing"
func TestGetAllArticles(t *testing.T){
	list := getAllArticles()

	if len(list) != len(articlesList) {
		t.Fail()
	}

	for i, v := range list {
		if v.Content != articlesList[i].Content ||
		 v.ID != articlesList[i].ID ||
		 v.Title != articlesList[i].Title {
		 t.Fail()
		 break
		}
	}
}