//go:build a && b
// +build a,b

package tag

func init() {
	Taglist = append(Taglist, "a+b")
}
