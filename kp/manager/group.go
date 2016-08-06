package treedata

import (
  "fmt"
)

type Group struct {
  name string
  topEntry *Entry
  next *Group
  prev *Group
  tree *Tree
  size int
}

func NewGroup(name string) *Group {
  k := new(Group)
  k.name = name
  return k
}

func (g *Group) AddEntry(e *Entry) *Entry {
  e.group = g
  if g.size == 0 {
    g.topEntry = e
  } else {
    prevEntry := g.topEntry
    e.prev = prevEntry
    g.topEntry = e
    prevEntry.next = e
  }

  g.size++

  return g.topEntry
}

func (g *Group) GetName() string {
  return g.name
}

func (g *Group) PrintEntries() {
  top := g.topEntry
  i := 0
  for top != nil {
    fmt.Printf("{%d} Entry '%s' \n", i, top.GetTitle())
    top = top.prev
    i++
  }
}

func (g *Group) Size() int {
  return g.size
}
