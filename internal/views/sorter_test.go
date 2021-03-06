package views

import (
	"sort"
	"testing"

	"github.com/derailed/k9s/internal/resource"
	"github.com/stretchr/testify/assert"
)

func TestGroupSort(t *testing.T) {
	uu := []struct {
		order  bool
		rows   []string
		expect []string
	}{
		{true, []string{"200m", "100m"}, []string{"100m", "200m"}},
		{false, []string{"200m", "100m"}, []string{"200m", "100m"}},
		{true, []string{"10", "1"}, []string{"1", "10"}},
		{false, []string{"10", "1"}, []string{"10", "1"}},
		{true, []string{"100Mi", "10Mi"}, []string{"10Mi", "100Mi"}},
		{false, []string{"100Mi", "10Mi"}, []string{"100Mi", "10Mi"}},
		{true, []string{"xyz", "abc"}, []string{"abc", "xyz"}},
		{false, []string{"xyz", "abc"}, []string{"xyz", "abc"}},
	}

	for _, u := range uu {
		g := groupSorter{rows: u.rows, asc: u.order}
		sort.Sort(g)
		assert.Equal(t, u.expect, g.rows)
	}
}

func TestRowSort(t *testing.T) {
	uu := []struct {
		order  bool
		rows   resource.Rows
		expect resource.Rows
	}{
		{
			true,
			resource.Rows{resource.Row{"200m"}, resource.Row{"100m"}},
			resource.Rows{resource.Row{"100m"}, resource.Row{"200m"}},
		},
		{
			false,
			resource.Rows{resource.Row{"200m"}, resource.Row{"100m"}},
			resource.Rows{resource.Row{"200m"}, resource.Row{"100m"}},
		},
		{
			true,
			resource.Rows{resource.Row{"200Mi"}, resource.Row{"100Mi"}},
			resource.Rows{resource.Row{"100Mi"}, resource.Row{"200Mi"}},
		},
		{
			false,
			resource.Rows{resource.Row{"200Mi"}, resource.Row{"100Mi"}},
			resource.Rows{resource.Row{"200Mi"}, resource.Row{"100Mi"}},
		},
	}

	for _, u := range uu {
		r := rowSorter{index: 0, rows: u.rows, asc: u.order}
		sort.Sort(r)
		assert.Equal(t, u.expect, r.rows)
	}
}
