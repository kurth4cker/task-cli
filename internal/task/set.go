// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task

import (
	"encoding/json"
	"io"
	"iter"
	"slices"
)

type Set struct {
	elements []Element
}

func (s *Set) Add(description string) Element {
	task := Element{
		Id:          s.newId(),
		Description: description,
		Status: Todo,
	}
	s.elements = append(s.elements, task)
	return task
}

func (s *Set) All() iter.Seq[Element] {
	return slices.Values(s.elements)
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

func (s *Set) Mark(id uint, status Status) bool {
	for i := range s.elements {
		if s.elements[i].Id == id {
			s.elements[i].Status = status
			return true
		}
	}
	return false
}

func (s *Set) ReadFrom(r io.Reader) (n int64, err error) {
	data, err := io.ReadAll(r)
	n = int64(len(data))
	if err != nil {
		return
	}
	if n == 0 {
		return
	}
	err = s.UnmarshalJSON(data)
	if err != nil {
		return
	}
	return
}

func (s *Set) WriteTo(w io.Writer) (int64, error) {
	data, err := s.MarshalJSON()
	if err != nil {
		return 0, err
	}
	n, err := w.Write(data)
	return int64(n), err
}

func (s *Set) IndentWriteTo(w io.Writer, prefix, indent string) (int64, error) {
	data, err := json.MarshalIndent(s, prefix, indent)
	if err != nil {
		return 0, err
	}
	n, err := w.Write(data)
	return int64(n), err
}

func (s *Set) Update(id uint, description string) bool {
	for i := range s.elements {
		if s.elements[i].Id == id {
			s.elements[i].Description = description
			return true
		}
	}
	return false
}

func (s *Set) Delete(id uint) bool {
	idx := slices.IndexFunc(s.elements, func(elem Element) bool {
		return elem.Id == id
	})
	if idx == -1 {
		return false
	}
	s.elements = slices.Delete(s.elements, idx, idx+1)
	return true
}

func (s *Set) Clone() *Set {
	return &Set{
		elements: slices.Clone(s.elements),
	}
}

func (s *Set) Equal(other *Set) bool {
	if s.Len() != other.Len() {
		return false
	}

	for e1 := range s.All() {
		e2, ok := other.Get(e1.Id)
		if !ok || !e1.Equal(e2) {
			return false
		}
	}

	return true
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
