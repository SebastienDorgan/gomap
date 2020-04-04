# gomap
Go library for manipulating generic maps and cast elements into native Go types and native Go structures

__gomap__ is super simple to use.

Start from a generic map, creating by and or using some decoding golang tools

```golang
m := map[string]interface{}{
    "v1": 1,
    "elt": map[string]interface{}{
        "v2": 2,
    },
    "vs": "test",
}

```
Cast it into a GMap

```golang
gm := gomap.GMap(m)
```

And that's it, you can start to play
```golang
vi, err := gm.Get("elt", "v2").Int(42) //42 is the default value
```
When a default value is specified, errors are never raised, if you want to have path and types checks do this instead
```golang
vi, err := gm.Get("elt", "v2").Int() //no default value
```
__gomap__ is also capable to set a custom struct using json annotations
```golang
type Elt struct{
    v2 int `json:"v2,omitempty"`
}
elt := Elt{}
err := gm.Get("elt").Object(&elt)
```
and empty path is ok
```golang
type Elt2 struct {
    V1  int    `json:"v1,omitempty"`
    Elt Elt    `json:"elt,omitempty"`
    VS  string `json:"vs,omitempty"`
}
elt2 := Elt2{}
err := gm.Get().Object(&elt2)
```
You can also play with slices

```golang
gm = GMap{//shortcut to define a GMap
    "v1": []interface{}{1, 2},
    "elt": GMap{
        "v2": []interface{}{int64(2), int64(3)},
    },
    "vs": []interface{}{"test", "test2"},
}
l, err := gm.Get("v2").Int64Slice()

```
See gomap_test.go for more examples