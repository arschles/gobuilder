package gobuilder

// Buildable is the interface for you to implement to transform your type
// into a Builder type. The builder type knows how to transform itself
// into real JSON according to the DSL that you write
type Buildable interface {
	ToBuilder() Builder
}
