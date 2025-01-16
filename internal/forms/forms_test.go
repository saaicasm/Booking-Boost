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
	r, _ := http.NewRequest("POST", "/test", nil)
	form := New(r.PostForm)

	isValid := form.Valid()

	if !isValid {
		t.Error("got invalid should be valid")
	}

	postedData := url.Values{}

	postedData.Add("a", "1")
	postedData.Add("b", "2")
	postedData.Add("c", "3")

	r.PostForm = postedData
	form = New(postedData)

	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("Required fields present but got error")
	}

	postedData.Add("empty", "")

	r.PostForm = postedData
	form = New(postedData)
	form.Required("empty")

	if form.Valid() {
		t.Error("should be invalid but got none")
	}

}

func TestHas(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	has := form.Has("whatever")
	if has {
		t.Error("shows a field when it should not")
	}

	postedData = url.Values{}
	postedData.Add("key1", "value1")
	form = New(postedData)

	res := form.Has("key1")

	if !res {
		t.Error("shows no field when it should")
	}

}

func TestMinLength(t *testing.T) {

	postedData := url.Values{}

	postedData.Add("some", "word")
	form := New(postedData)

	minV := form.MinLength("some", 5)

	if minV {
		t.Error("should give an error its short")
	}

	postedData.Add("valid", "wordle")
	form = New(postedData)

	minV = form.MinLength("valid", 5)

	if !minV {
		t.Error("giving an error with valid string")
	}

}

func TestIsMail(t *testing.T) {
	postedDatta := url.Values{}
	form := New(postedDatta)

	email := "iamlego@las.com"

	form.IsEmail("x")
	if form.Valid() {
		t.Error("Your mom is not an email")
	}

	postedDatta = url.Values{}
	postedDatta.Add("email", email)
	form = New(postedDatta)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("Invalid email")
	}

}

func TestGet(t *testing.T) {
	err := errors{}

	err.Add("err1", "first error")
	err.Add("err2", "second error")

	res := err.Get("err1")

	if res == "" {
		t.Error("Got empty when there is field")
	}

	res = err.Get("some")

	if res != "" {
		t.Error("Got value when invalid error field")
	}

}
