package gomap

//BoolSlice cast element value into bool
func (elt *Element) BoolSlice(defaultValue ...[]bool) ([]bool, error) {
	defValue := func() *[]bool {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}
	def := defValue()
	if elt.Value == nil {
		if def == nil {
			var v []bool
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	switch v := elt.Value.(type) {
	case []interface{}:
		return boolCastInterfaceType(v, def)
	case GSlice:
		return boolCastGSliceType(v, def)
	default:
		return nil, NewWrongTypeError("[]bool", elt.Value)
	}
}

func boolCastInterfaceType(list []interface{}, def *[]bool) ([]bool, error){
	var res []bool
	for _, e := range list{
		v, ok := e.(bool)
		if !ok {
			if def == nil{
				return nil, NewWrongTypeError("bool", e)
			}
			return *def, nil
		}
		res = append(res, v)
	}
	return res, nil
}

func boolCastGSliceType(list GSlice, def *[]bool) ([]bool, error){
	var res []bool
	for _, e := range list{
		v, ok := e.(bool)
		if !ok {
			if def == nil{
				return nil, NewWrongTypeError("bool", e)
			}
			return *def, nil
		}
		res = append(res, v)
	}
	return res, nil
}

//IntSlice cast element value into int
func (elt *Element) IntSlice(defaultValue ...[]int) ([]int, error) {
	defValue := func() *[]int {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}
	def := defValue()
	if elt.Value == nil {
		if def == nil {
			var v []int
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	switch v := elt.Value.(type) {
	case []interface{}:
		return intCastInterfaceType(v, def)
	case GSlice:
		return intCastGSliceType(v, def)
	default:
		return nil, NewWrongTypeError("[]int", elt.Value)
	}
}

func intCastInterfaceType(list []interface{}, def *[]int) ([]int, error){
	var res []int
	for _, e := range list{
		v, ok := e.(int)
		if !ok {
			if def == nil{
				return nil, NewWrongTypeError("int", e)
			}
			return *def, nil
		}
		res = append(res, v)
	}
	return res, nil
}

func intCastGSliceType(list GSlice, def *[]int) ([]int, error){
	var res []int
	for _, e := range list{
		v, ok := e.(int)
		if !ok {
			if def == nil{
				return nil, NewWrongTypeError("int", e)
			}
			return *def, nil
		}
		res = append(res, v)
	}
	return res, nil
}

//Int8Slice cast element value into int8
func (elt *Element) Int8Slice(defaultValue ...[]int8) ([]int8, error) {
	defValue := func() *[]int8 {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}
	def := defValue()
	if elt.Value == nil {
		if def == nil {
			var v []int8
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	switch v := elt.Value.(type) {
	case []interface{}:
		return int8CastInterfaceType(v, def)
	case GSlice:
		return int8CastGSliceType(v, def)
	default:
		return nil, NewWrongTypeError("[]int8", elt.Value)
	}
}

func int8CastInterfaceType(list []interface{}, def *[]int8) ([]int8, error){
	var res []int8
	for _, e := range list{
		v, ok := e.(int8)
		if !ok {
			if def == nil{
				return nil, NewWrongTypeError("int8", e)
			}
			return *def, nil
		}
		res = append(res, v)
	}
	return res, nil
}

func int8CastGSliceType(list GSlice, def *[]int8) ([]int8, error){
	var res []int8
	for _, e := range list{
		v, ok := e.(int8)
		if !ok {
			if def == nil{
				return nil, NewWrongTypeError("int8", e)
			}
			return *def, nil
		}
		res = append(res, v)
	}
	return res, nil
}

//Int16Slice cast element value into int16
func (elt *Element) Int16Slice(defaultValue ...[]int16) ([]int16, error) {
	defValue := func() *[]int16 {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}
	def := defValue()
	if elt.Value == nil {
		if def == nil {
			var v []int16
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	switch v := elt.Value.(type) {
	case []interface{}:
		return int16CastInterfaceType(v, def)
	case GSlice:
		return int16CastGSliceType(v, def)
	default:
		return nil, NewWrongTypeError("[]int16", elt.Value)
	}
}

func int16CastInterfaceType(list []interface{}, def *[]int16) ([]int16, error){
	var res []int16
	for _, e := range list{
		v, ok := e.(int16)
		if !ok {
			if def == nil{
				return nil, NewWrongTypeError("int16", e)
			}
			return *def, nil
		}
		res = append(res, v)
	}
	return res, nil
}

func int16CastGSliceType(list GSlice, def *[]int16) ([]int16, error){
	var res []int16
	for _, e := range list{
		v, ok := e.(int16)
		if !ok {
			if def == nil{
				return nil, NewWrongTypeError("int16", e)
			}
			return *def, nil
		}
		res = append(res, v)
	}
	return res, nil
}

//Int32Slice cast element value into int32
func (elt *Element) Int32Slice(defaultValue ...[]int32) ([]int32, error) {
	defValue := func() *[]int32 {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}
	def := defValue()
	if elt.Value == nil {
		if def == nil {
			var v []int32
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	switch v := elt.Value.(type) {
	case []interface{}:
		return int32CastInterfaceType(v, def)
	case GSlice:
		return int32CastGSliceType(v, def)
	default:
		return nil, NewWrongTypeError("[]int32", elt.Value)
	}
}

func int32CastInterfaceType(list []interface{}, def *[]int32) ([]int32, error){
	var res []int32
	for _, e := range list{
		v, ok := e.(int32)
		if !ok {
			if def == nil{
				return nil, NewWrongTypeError("int32", e)
			}
			return *def, nil
		}
		res = append(res, v)
	}
	return res, nil
}

func int32CastGSliceType(list GSlice, def *[]int32) ([]int32, error){
	var res []int32
	for _, e := range list{
		v, ok := e.(int32)
		if !ok {
			if def == nil{
				return nil, NewWrongTypeError("int32", e)
			}
			return *def, nil
		}
		res = append(res, v)
	}
	return res, nil
}

//Int64Slice cast element value into int64
func (elt *Element) Int64Slice(defaultValue ...[]int64) ([]int64, error) {
	defValue := func() *[]int64 {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}
	def := defValue()
	if elt.Value == nil {
		if def == nil {
			var v []int64
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	switch v := elt.Value.(type) {
	case []interface{}:
		return int64CastInterfaceType(v, def)
	case GSlice:
		return int64CastGSliceType(v, def)
	default:
		return nil, NewWrongTypeError("[]int64", elt.Value)
	}
}

func int64CastInterfaceType(list []interface{}, def *[]int64) ([]int64, error){
	var res []int64
	for _, e := range list{
		v, ok := e.(int64)
		if !ok {
			if def == nil{
				return nil, NewWrongTypeError("int64", e)
			}
			return *def, nil
		}
		res = append(res, v)
	}
	return res, nil
}

func int64CastGSliceType(list GSlice, def *[]int64) ([]int64, error){
	var res []int64
	for _, e := range list{
		v, ok := e.(int64)
		if !ok {
			if def == nil{
				return nil, NewWrongTypeError("int64", e)
			}
			return *def, nil
		}
		res = append(res, v)
	}
	return res, nil
}

//Uint8Slice cast element value into uint8
func (elt *Element) Uint8Slice(defaultValue ...[]uint8) ([]uint8, error) {
	defValue := func() *[]uint8 {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}
	def := defValue()
	if elt.Value == nil {
		if def == nil {
			var v []uint8
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	switch v := elt.Value.(type) {
	case []interface{}:
		return uint8CastInterfaceType(v, def)
	case GSlice:
		return uint8CastGSliceType(v, def)
	default:
		return nil, NewWrongTypeError("[]uint8", elt.Value)
	}
}

func uint8CastInterfaceType(list []interface{}, def *[]uint8) ([]uint8, error){
	var res []uint8
	for _, e := range list{
		v, ok := e.(uint8)
		if !ok {
			if def == nil{
				return nil, NewWrongTypeError("uint8", e)
			}
			return *def, nil
		}
		res = append(res, v)
	}
	return res, nil
}

func uint8CastGSliceType(list GSlice, def *[]uint8) ([]uint8, error){
	var res []uint8
	for _, e := range list{
		v, ok := e.(uint8)
		if !ok {
			if def == nil{
				return nil, NewWrongTypeError("uint8", e)
			}
			return *def, nil
		}
		res = append(res, v)
	}
	return res, nil
}

//Uint16Slice cast element value into uint16
func (elt *Element) Uint16Slice(defaultValue ...[]uint16) ([]uint16, error) {
	defValue := func() *[]uint16 {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}
	def := defValue()
	if elt.Value == nil {
		if def == nil {
			var v []uint16
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	switch v := elt.Value.(type) {
	case []interface{}:
		return uint16CastInterfaceType(v, def)
	case GSlice:
		return uint16CastGSliceType(v, def)
	default:
		return nil, NewWrongTypeError("[]uint16", elt.Value)
	}
}

func uint16CastInterfaceType(list []interface{}, def *[]uint16) ([]uint16, error){
	var res []uint16
	for _, e := range list{
		v, ok := e.(uint16)
		if !ok {
			if def == nil{
				return nil, NewWrongTypeError("uint16", e)
			}
			return *def, nil
		}
		res = append(res, v)
	}
	return res, nil
}

func uint16CastGSliceType(list GSlice, def *[]uint16) ([]uint16, error){
	var res []uint16
	for _, e := range list{
		v, ok := e.(uint16)
		if !ok {
			if def == nil{
				return nil, NewWrongTypeError("uint16", e)
			}
			return *def, nil
		}
		res = append(res, v)
	}
	return res, nil
}

//Uint32Slice cast element value into uint32
func (elt *Element) Uint32Slice(defaultValue ...[]uint32) ([]uint32, error) {
	defValue := func() *[]uint32 {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}
	def := defValue()
	if elt.Value == nil {
		if def == nil {
			var v []uint32
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	switch v := elt.Value.(type) {
	case []interface{}:
		return uint32CastInterfaceType(v, def)
	case GSlice:
		return uint32CastGSliceType(v, def)
	default:
		return nil, NewWrongTypeError("[]uint32", elt.Value)
	}
}

func uint32CastInterfaceType(list []interface{}, def *[]uint32) ([]uint32, error){
	var res []uint32
	for _, e := range list{
		v, ok := e.(uint32)
		if !ok {
			if def == nil{
				return nil, NewWrongTypeError("uint32", e)
			}
			return *def, nil
		}
		res = append(res, v)
	}
	return res, nil
}

func uint32CastGSliceType(list GSlice, def *[]uint32) ([]uint32, error){
	var res []uint32
	for _, e := range list{
		v, ok := e.(uint32)
		if !ok {
			if def == nil{
				return nil, NewWrongTypeError("uint32", e)
			}
			return *def, nil
		}
		res = append(res, v)
	}
	return res, nil
}

//Uint64Slice cast element value into uint64
func (elt *Element) Uint64Slice(defaultValue ...[]uint64) ([]uint64, error) {
	defValue := func() *[]uint64 {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}
	def := defValue()
	if elt.Value == nil {
		if def == nil {
			var v []uint64
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	switch v := elt.Value.(type) {
	case []interface{}:
		return uint64CastInterfaceType(v, def)
	case GSlice:
		return uint64CastGSliceType(v, def)
	default:
		return nil, NewWrongTypeError("[]uint64", elt.Value)
	}
}

func uint64CastInterfaceType(list []interface{}, def *[]uint64) ([]uint64, error){
	var res []uint64
	for _, e := range list{
		v, ok := e.(uint64)
		if !ok {
			if def == nil{
				return nil, NewWrongTypeError("uint64", e)
			}
			return *def, nil
		}
		res = append(res, v)
	}
	return res, nil
}

func uint64CastGSliceType(list GSlice, def *[]uint64) ([]uint64, error){
	var res []uint64
	for _, e := range list{
		v, ok := e.(uint64)
		if !ok {
			if def == nil{
				return nil, NewWrongTypeError("uint64", e)
			}
			return *def, nil
		}
		res = append(res, v)
	}
	return res, nil
}

//Float32Slice cast element value into float32
func (elt *Element) Float32Slice(defaultValue ...[]float32) ([]float32, error) {
	defValue := func() *[]float32 {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}
	def := defValue()
	if elt.Value == nil {
		if def == nil {
			var v []float32
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	switch v := elt.Value.(type) {
	case []interface{}:
		return float32CastInterfaceType(v, def)
	case GSlice:
		return float32CastGSliceType(v, def)
	default:
		return nil, NewWrongTypeError("[]float32", elt.Value)
	}
}

func float32CastInterfaceType(list []interface{}, def *[]float32) ([]float32, error){
	var res []float32
	for _, e := range list{
		v, ok := e.(float32)
		if !ok {
			if def == nil{
				return nil, NewWrongTypeError("float32", e)
			}
			return *def, nil
		}
		res = append(res, v)
	}
	return res, nil
}

func float32CastGSliceType(list GSlice, def *[]float32) ([]float32, error){
	var res []float32
	for _, e := range list{
		v, ok := e.(float32)
		if !ok {
			if def == nil{
				return nil, NewWrongTypeError("float32", e)
			}
			return *def, nil
		}
		res = append(res, v)
	}
	return res, nil
}

//Float64Slice cast element value into float64
func (elt *Element) Float64Slice(defaultValue ...[]float64) ([]float64, error) {
	defValue := func() *[]float64 {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}
	def := defValue()
	if elt.Value == nil {
		if def == nil {
			var v []float64
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	switch v := elt.Value.(type) {
	case []interface{}:
		return float64CastInterfaceType(v, def)
	case GSlice:
		return float64CastGSliceType(v, def)
	default:
		return nil, NewWrongTypeError("[]float64", elt.Value)
	}
}

func float64CastInterfaceType(list []interface{}, def *[]float64) ([]float64, error){
	var res []float64
	for _, e := range list{
		v, ok := e.(float64)
		if !ok {
			if def == nil{
				return nil, NewWrongTypeError("float64", e)
			}
			return *def, nil
		}
		res = append(res, v)
	}
	return res, nil
}

func float64CastGSliceType(list GSlice, def *[]float64) ([]float64, error){
	var res []float64
	for _, e := range list{
		v, ok := e.(float64)
		if !ok {
			if def == nil{
				return nil, NewWrongTypeError("float64", e)
			}
			return *def, nil
		}
		res = append(res, v)
	}
	return res, nil
}

//StringSlice cast element value into string
func (elt *Element) StringSlice(defaultValue ...[]string) ([]string, error) {
	defValue := func() *[]string {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}
	def := defValue()
	if elt.Value == nil {
		if def == nil {
			var v []string
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	switch v := elt.Value.(type) {
	case []interface{}:
		return stringCastInterfaceType(v, def)
	case GSlice:
		return stringCastGSliceType(v, def)
	default:
		return nil, NewWrongTypeError("[]string", elt.Value)
	}
}

func stringCastInterfaceType(list []interface{}, def *[]string) ([]string, error){
	var res []string
	for _, e := range list{
		v, ok := e.(string)
		if !ok {
			if def == nil{
				return nil, NewWrongTypeError("string", e)
			}
			return *def, nil
		}
		res = append(res, v)
	}
	return res, nil
}

func stringCastGSliceType(list GSlice, def *[]string) ([]string, error){
	var res []string
	for _, e := range list{
		v, ok := e.(string)
		if !ok {
			if def == nil{
				return nil, NewWrongTypeError("string", e)
			}
			return *def, nil
		}
		res = append(res, v)
	}
	return res, nil
}

