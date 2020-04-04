package gomap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoMapTypes(t *testing.T) {
	m := GMap{
		"v1": 1,
		"elt": GMap{
			"v2": 2,
		},
		"vs": "test",
	}
	v, err := m.Get("v1").Int(42)
	assert.NoError(t, err)
	assert.Equal(t, 1, v)

	v, err = m.Get("elt", "v2").Int(42)
	assert.NoError(t, err)
	assert.Equal(t, 2, v)

	v, err = m.Get("a", "b", "v").Int(42)
	assert.NoError(t, err)
	assert.Equal(t, v, 42)

	v, err = m.Get("a", "b", "v").Int()
	assert.Error(t, err)
	assert.Equal(t, "Wrong path a/b/v", err.Error())
	assert.Equal(t, v, 0)

	vs, err := m.Get("vs").String("42")
	assert.NoError(t, err)
	assert.Equal(t, "test", vs)

	vs, err = m.Get("vs2").String("42")
	assert.NoError(t, err)
	assert.Equal(t, "42", vs)

	vs, err = m.Get("vs2").String()
	assert.Error(t, err)
	assert.Equal(t, "Wrong path vs2", err.Error())
	assert.Equal(t, "", vs)
}

func TestGoSlices(t *testing.T) {
	m := GMap{
		"v1": []interface{}{1, 2},
		"elt": GMap{
			"v2": []interface{}{int64(2), int64(3)},
		},
		"vs": []interface{}{"test", "test2"},
	}
	v, err := m.Get("v1").IntSlice([]int{42})
	assert.NoError(t, err)
	assert.EqualValues(t, []int{1, 2}, v)

	v64, err := m.Get("elt", "v2").Int64Slice([]int64{42})
	assert.NoError(t, err)
	assert.EqualValues(t, []int64{2, 3}, v64)

	v, err = m.Get("elt", "v2").IntSlice([]int{42})
	assert.NoError(t, err)
	assert.EqualValues(t, []int{42}, v)

	v, err = m.Get("elt", "v2").IntSlice()
	assert.Error(t, err)
	assert.Equal(t, "Wrong type: expected int actual int64", err.Error())

	vs, err := m.Get("vs").StringSlice()
	assert.NoError(t, err)
	assert.EqualValues(t, []string{"test", "test2"}, vs)
}

func TestGoMapStruct(t *testing.T) {
	type Elt struct {
		V2 int `json:"v2,omitempty"`
	}

	type Elt2 struct {
		V1  int    `json:"v1,omitempty"`
		Elt Elt    `json:"elt,omitempty"`
		VS  string `json:"vs,omitempty"`
	}
	m := GMap{
		"v1": 1,
		"elt": GMap{
			"v2": 2,
		},
		"vs": "test",
	}
	elt := Elt{}
	err := m.Get("elt").Object(&elt)
	assert.NoError(t, err)
	assert.Equal(t, 2, elt.V2)

	elt2 := Elt2{}
	err = m.Get().Object(&elt2)
	assert.NoError(t, err)
	assert.EqualValues(t,
		Elt2{
			V1: 1,
			Elt: Elt{
				V2: 2,
			},
			VS: "test",
		},
		elt2)

}
