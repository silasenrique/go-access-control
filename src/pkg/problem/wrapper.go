package problem

type problemDescription interface {
	String() string
}

type Type problemDescription
type Title problemDescription
type Details problemDescription
