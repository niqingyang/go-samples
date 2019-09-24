package hooks

import (
	"errors"
	"fmt"
	"sort"
)

type Filter interface {
	Apply() interface{}
}

type Action interface {
	Apply()
}

type Hooks struct {
	filters      map[string]map[int]map[string]Filter
	actions      map[string]map[int]map[string]Action
	mergeFilters map[string]bool
}

func (h *Hooks) AddFilter(tag string, filter Filter, priority int) bool {

	if h.filters == nil {
		h.filters = make(map[string]map[int]map[string]Filter)
	}

	if _, ok := h.filters[tag]; !ok {
		h.filters[tag] = make(map[int]map[string]Filter)
	}

	if _, ok := h.filters[tag][priority]; !ok {
		h.filters[tag][priority] = make(map[string]Filter)
	}

	key := fmt.Sprintf("%p", filter)

	h.filters[tag][priority][key] = filter

	return true
}

func (h *Hooks) ApplyFilter(tag string) (interface{}, error) {
	if _, ok := h.filters[tag]; !ok {
		return nil, errors.New("")
	}

	keys := make([]int, len(h.filters[tag]))

	for key, _ := range h.filters[tag] {
		keys = append(keys, key)
	}

	sort.Ints(keys[:])

	for key := range keys {
		filters := h.filters[tag][key]

		for _, filter := range filters {
			filter.Apply()
		}
	}

	return nil, nil
}
