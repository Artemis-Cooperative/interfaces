package packages

type testkit_types interface {
	AllKinds() []string
	AllTypeInstances() []any
	BasicInstances() []any
	BasicKinds() []string
	ComparableInstances() []any
	ComparableKinds() []string
	ElementInstances() []any
	ElementKinds() []string
	KeyInstances() []any
	KeyKinds() []string
	Kinds(includeTags []string, exclude []string) []string
	Instances(includeTags []string, excludeKinds []string) []any
	IntegerInstances() []any
	IntegerKinds() []string
	OrderedInstances() []any
	OrderedKinds() []string
	NumbericInstances() []any
	NumbericKinds() []string
	SliceInstances() []any
	MapInstances() []any
}

type types_structs interface {
	Equals(s1 interface{}, s2 interface{}) bool
	IsStruct(v any) bool
}
