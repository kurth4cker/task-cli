// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task

import "iter"

type Set struct {
	elements []Element
}

func (s *Set) Add(description string) {
	task := Element{
		Id:          s.newId(),
		Description: description,
	}
	s.elements = append(s.elements, task)
}

func (s *Set) All() iter.Seq[Element] {
	return func(yield func(Element) bool) {
		for _, elem := range s.elements {
			if !yield(elem) {
				break
			}
		}
	}
}

func (s *Set) Descriptions() []string {
	descriptions := make([]string, 0, s.Len())
	for _, elem := range s.elements {
		descriptions = append(descriptions, elem.Description)
	}
	return descriptions
}

func (s *Set) Ids() []uint {
	ids := make([]uint, 0, s.Len())
	for _, elem := range s.elements {
		ids = append(ids, elem.Id)
	}
	return ids
}

func (s *Set) Len() int {
	return len(s.elements)
}

func (s *Set) newId() uint {
	var id uint
	for _, elem := range s.elements {
		if id == elem.Id {
			id++
		}
	}
	return id
}
