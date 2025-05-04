// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task

import (
	"encoding/json"
	"io"
	"iter"
)

type Set struct {
	elements []Element
}

func (s *Set) Add(description string) {
	task := Element{
		Id:          s.newId(),
		Description: description,
		Status: Todo,
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

func (s *Set) Len() int {
	return len(s.elements)
}

func (s *Set) AddElement(elem Element) {
	s.elements = append(s.elements, elem)
}

func (s *Set) MarshalJSON() ([]byte, error) {
	return json.Marshal(&s.elements)
}

func (s *Set) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &s.elements)
}

func (s *Set) Get(id uint) (Element, bool) {
	var last Element
	for elem := range s.All() {
		if elem.Id == id {
			return elem, true
		}
		last = elem
	}
	return last, false
}

func (s *Set) Mark(id uint, status Status) {
	for i := range s.elements {
		if s.elements[i].Id == id {
			s.elements[i].Status = status
		}
	}
}

func (s *Set) ReadFrom(r io.Reader) (int64, error) {
	// TODO: return correct read bytes
	data, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	if err := s.UnmarshalJSON(data); err != nil {
		return 0, err
	}
	return 0, nil
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
