package treedata

import (
  "errors"
  "fmt"
)

type Tree struct {
  name string
  topGroup *Group
  next *Tree
  prev *Tree
  size int
}

func NewTree(name string) *Tree {
  t := new(Tree)

  t.name = name
  return t
}

func (t *Tree) SetName(name string) {
  t.name = name
}

func (t *Tree) GetName() string {
  return t.name
}

func (t *Tree) AddGroup(name string) *Group {
  if t.size == 0 {
    t.topGroup = &Group{name, nil, nil, nil, t, 0}
  } else {
    prevGroup := t.topGroup
    t.topGroup = &Group{name, nil, nil, prevGroup, t, 0}
    prevGroup.next = t.topGroup
  }

  t.size++
  return t.topGroup
}

func (t *Tree) GetGroup(index int) (*Group, error) {
  top := t.topGroup
  i := 0
  for top != nil {
    if (i == index) {
      return top, nil
    }
    top = top.prev
    i++
  }
  return nil, errors.New("Unable to find group")
}

func (t *Tree) GetGroupByName(name string) (*Group, error) {
  top := t.topGroup
  for top != nil {

    if (top.GetName() == name) {
      return top, nil
    }

    top = top.prev
  }

  return nil, errors.New(fmt.Sprintf("Unable to find %s", name))
}

func (t *Tree) PrintGroups() {
  top := t.topGroup
  i := 0
  for top != nil {
    fmt.Printf("{%d} %s\n", i, top.GetName())
    top = top.prev
    i++
  }
}

func (t *Tree) GroupExists(name string) bool {
  top := t.topGroup
  for top != nil {
    if top.GetName() == name {
      return true
    }
    top = top.prev
  }

  return false
}

func (t *Tree) Size() int {
  return t.size
}
