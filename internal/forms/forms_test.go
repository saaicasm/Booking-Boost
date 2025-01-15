package forms

import (
	"net/http"
	"net/url"
	"testing"
)

func TestValid(t *testing.T) {
	v := url.Values{}
	v.Set("name", "Ava")
	v.Add("friend", "Jess")
	v.Add("friend", "Sarah")
	v.Add("friend", "Zoe")
	e := errors{}

	f := Form{
		v,
		e,
	}

	bl := f.Valid()

	if len(f.Errors) != 0 && bl == true {
		t.Error("something is wrong")
	}

}

func TestRequired(t *testing.T) {
	postedData := url.Values{}

	postedData.Add("a", "1")
	postedData.Add("b", "2")
	postedData.Add("c", "3")

	r, _ := http.NewRequest("POST", "/test", nil)
	r.PostForm = postedData
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("Required fields not present")
	}

}

func TestHas(t *testing.T) {
	postField := url.Values{}

	postField.Set("key1", "value1")

	r, _ := http.NewRequest("POST", "/testfield", nil)

	r.PostForm = postField
	form := New(r.PostForm)

	res := form.Has("key1", *r)

	if res {
		t.Error("Value missing")
	}

}

// func (f *Form) Has(field string, r http.Request) bool {
// 	x := r.Form.Get(field)
// 	if x == "" {
// 		f.Errors.Add(field, "This feild cannot be blank")
// 		return false
// 	}
// 	return true
// }
