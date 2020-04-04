package gomap

//Bool cast element value into bool
func (elt *Element) Bool(defaultValue ...bool) (bool, error) {
	defValue := func() *bool {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}
	def := defValue()
	if elt.Value == nil {
		def := defValue()
		if def == nil {
			var v bool
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	v, ok := elt.Value.(bool)
	if !ok {
		if def == nil{
			var v bool
			return v, NewWrongTypeError("bool", elt.Value)
		}
		return *def, nil
	}
	return v, nil
}

//Int cast element value into int
func (elt *Element) Int(defaultValue ...int) (int, error) {
	defValue := func() *int {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}
	def := defValue()
	if elt.Value == nil {
		def := defValue()
		if def == nil {
			var v int
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	v, ok := elt.Value.(int)
	if !ok {
		if def == nil{
			var v int
			return v, NewWrongTypeError("int", elt.Value)
		}
		return *def, nil
	}
	return v, nil
}

//Int8 cast element value into int8
func (elt *Element) Int8(defaultValue ...int8) (int8, error) {
	defValue := func() *int8 {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}
	def := defValue()
	if elt.Value == nil {
		def := defValue()
		if def == nil {
			var v int8
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	v, ok := elt.Value.(int8)
	if !ok {
		if def == nil{
			var v int8
			return v, NewWrongTypeError("int8", elt.Value)
		}
		return *def, nil
	}
	return v, nil
}

//Int16 cast element value into int16
func (elt *Element) Int16(defaultValue ...int16) (int16, error) {
	defValue := func() *int16 {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}
	def := defValue()
	if elt.Value == nil {
		def := defValue()
		if def == nil {
			var v int16
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	v, ok := elt.Value.(int16)
	if !ok {
		if def == nil{
			var v int16
			return v, NewWrongTypeError("int16", elt.Value)
		}
		return *def, nil
	}
	return v, nil
}

//Int32 cast element value into int32
func (elt *Element) Int32(defaultValue ...int32) (int32, error) {
	defValue := func() *int32 {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}
	def := defValue()
	if elt.Value == nil {
		def := defValue()
		if def == nil {
			var v int32
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	v, ok := elt.Value.(int32)
	if !ok {
		if def == nil{
			var v int32
			return v, NewWrongTypeError("int32", elt.Value)
		}
		return *def, nil
	}
	return v, nil
}

//Int64 cast element value into int64
func (elt *Element) Int64(defaultValue ...int64) (int64, error) {
	defValue := func() *int64 {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}
	def := defValue()
	if elt.Value == nil {
		def := defValue()
		if def == nil {
			var v int64
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	v, ok := elt.Value.(int64)
	if !ok {
		if def == nil{
			var v int64
			return v, NewWrongTypeError("int64", elt.Value)
		}
		return *def, nil
	}
	return v, nil
}

//Uint8 cast element value into uint8
func (elt *Element) Uint8(defaultValue ...uint8) (uint8, error) {
	defValue := func() *uint8 {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}
	def := defValue()
	if elt.Value == nil {
		def := defValue()
		if def == nil {
			var v uint8
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	v, ok := elt.Value.(uint8)
	if !ok {
		if def == nil{
			var v uint8
			return v, NewWrongTypeError("uint8", elt.Value)
		}
		return *def, nil
	}
	return v, nil
}

//Uint16 cast element value into uint16
func (elt *Element) Uint16(defaultValue ...uint16) (uint16, error) {
	defValue := func() *uint16 {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}
	def := defValue()
	if elt.Value == nil {
		def := defValue()
		if def == nil {
			var v uint16
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	v, ok := elt.Value.(uint16)
	if !ok {
		if def == nil{
			var v uint16
			return v, NewWrongTypeError("uint16", elt.Value)
		}
		return *def, nil
	}
	return v, nil
}

//Uint32 cast element value into uint32
func (elt *Element) Uint32(defaultValue ...uint32) (uint32, error) {
	defValue := func() *uint32 {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}
	def := defValue()
	if elt.Value == nil {
		def := defValue()
		if def == nil {
			var v uint32
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	v, ok := elt.Value.(uint32)
	if !ok {
		if def == nil{
			var v uint32
			return v, NewWrongTypeError("uint32", elt.Value)
		}
		return *def, nil
	}
	return v, nil
}

//Uint64 cast element value into uint64
func (elt *Element) Uint64(defaultValue ...uint64) (uint64, error) {
	defValue := func() *uint64 {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}
	def := defValue()
	if elt.Value == nil {
		def := defValue()
		if def == nil {
			var v uint64
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	v, ok := elt.Value.(uint64)
	if !ok {
		if def == nil{
			var v uint64
			return v, NewWrongTypeError("uint64", elt.Value)
		}
		return *def, nil
	}
	return v, nil
}

//Float32 cast element value into float32
func (elt *Element) Float32(defaultValue ...float32) (float32, error) {
	defValue := func() *float32 {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}
	def := defValue()
	if elt.Value == nil {
		def := defValue()
		if def == nil {
			var v float32
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	v, ok := elt.Value.(float32)
	if !ok {
		if def == nil{
			var v float32
			return v, NewWrongTypeError("float32", elt.Value)
		}
		return *def, nil
	}
	return v, nil
}

//Float64 cast element value into float64
func (elt *Element) Float64(defaultValue ...float64) (float64, error) {
	defValue := func() *float64 {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}
	def := defValue()
	if elt.Value == nil {
		def := defValue()
		if def == nil {
			var v float64
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	v, ok := elt.Value.(float64)
	if !ok {
		if def == nil{
			var v float64
			return v, NewWrongTypeError("float64", elt.Value)
		}
		return *def, nil
	}
	return v, nil
}

//String cast element value into string
func (elt *Element) String(defaultValue ...string) (string, error) {
	defValue := func() *string {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}
	def := defValue()
	if elt.Value == nil {
		def := defValue()
		if def == nil {
			var v string
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	v, ok := elt.Value.(string)
	if !ok {
		if def == nil{
			var v string
			return v, NewWrongTypeError("string", elt.Value)
		}
		return *def, nil
	}
	return v, nil
}

