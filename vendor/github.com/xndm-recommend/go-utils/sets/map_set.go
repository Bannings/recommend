package sets

import (
	"github.com/deckarep/golang-set"
)

type HashSet struct {
	Set_Hash mapset.Set
}

func NewSet(s ...interface{}) *HashSet {
	return &(HashSet{Set_Hash: mapset.NewSet(s)})
}

func NewSetFromSlice(s []interface{}) *HashSet {
	return &(HashSet{Set_Hash: mapset.NewSetFromSlice(s)})
}

func (this *HashSet) IsSubSet(other HashSet) bool {
	return this.Set_Hash.IsSubset(other.Set_Hash)
}

func (this *HashSet) SetToSlice() []interface{} {
	return this.Set_Hash.ToSlice()
}

func (this *HashSet) UnionSet(other HashSet) *HashSet {
	return &(HashSet{Set_Hash: this.Set_Hash.Union(other.Set_Hash)})
}

// 生成集合 set 对集合 other 的差集
func (this *HashSet) Difference(other *HashSet) *HashSet {
	return &(HashSet{Set_Hash: this.Set_Hash.Difference(other.Set_Hash)})
}
