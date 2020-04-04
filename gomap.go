//Package gomap implements utilities to manipulate generic maps and cast elements into native Go types and native Go structures
//
// The package is really simple to use
package gomap

import "encoding/json"

//GMap overload map with utility functions
type GMap map[string]interface{}

//Element an element of the map
type Element struct {
	Path  []string
	Value interface{}
}

//Get returns map elements along path
func (m GMap) Get(path ...string) *Element {
	if len(path) == 0 {
		return &Element{
			Path:  path,
			Value: m,
		}
	}
	if len(path) == 1 {
		return element(m, path, path[0])
	}
	return element(m, path, path[0]).Get(path[1:]...)
}

//Get returns map elements along path
func (elt *Element) Get(path ...string) *Element {
	next := func(path ...string) *Element {
		if elt.Value == nil {
			return &Element{
				Path: elt.Path,
			}
		}
		m, ok := elt.Value.(map[string]interface{})
		if ok {
			return element(m, elt.Path, path[0])
		}
		gm, ok := elt.Value.(GMap)
		if ok {
			return element(gm, elt.Path, path[0])
		}
		return &Element{
			Path: elt.Path,
		}
	}
	if len(path) == 0 {
		return elt
	}
	if len(path) == 1 {
		return next(path...)
	}
	return next(path...).Get(path[1:]...)
}

func element(m GMap, path []string, key string) *Element {
	if v, ok := m[key]; ok {
		return &Element{
			Value: v,
			Path:  path,
		}
	}
	return &Element{
		Path: path,
	}
}

//ToJSON marshal a gmap into json content
func (m GMap) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

//FromJSON laod json content into a GMap
func (m GMap) FromJSON(content []byte) error {
	return json.Unmarshal(content, &m)
}
