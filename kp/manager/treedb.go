package treedata

import "errors"

type TreeDB struct {
  name string
  top *Tree
  size int
}

func NewTreeDB(name string) *TreeDB {
  t := new(TreeDB)
  t.name = name
  t.size = 0
  t.top = nil
  return t
}

func (t *TreeDB) AddTree(tr *Tree) *Tree {
  if t.size == 0 {
    t.top = tr
  } else {
    prev := t.top
    tr.prev = prev
    t.top = tr
    prev.next = t.top
  }

  t.size++
  return t.top
}

func (t *TreeDB) TreeExists(name string) bool {
  top := t.top
  for top != nil {
    if (top.GetName() == name) {
      return true
    }
    top = top.prev
  }
  return false
}

func (t *TreeDB) GetTreeByName(name string) (*Tree, error) {
  top := t.top
  for top != nil {
    if (top.GetName() == name) {
      return top, nil
    }
    top = top.prev
  }
  return nil, errors.New("Unable to find tree")
}

func (t *TreeDB) XML() string {

  var str string = "<?xml version=\"1.0\" encoding=\"utf-8\" standalone=\"yes\"?>\n";

  topTree := t.top

  if topTree != nil {
    str += "<pwlist>\n"
    for topTree != nil {
      topGroup := topTree.topGroup
      for topGroup != nil {
        topEntry := topGroup.topEntry
        for topEntry != nil {
          str += "\t<pwentry>\n"
          str += "\t\t<group tree=\"" + topTree.GetName() + "\">" + topGroup.GetName() + "</group>\n"
          str += "\t\t<title></title>\n"
          str += "\t\t<username></username>\n"
          str += "\t\t<url></url>\n"
          str += "\t\t<password></password>\n"
          str += "\t\t<notes></notes>\n"
          str += "\t</pwentry>\n"
          topEntry = topEntry.prev
        }
        topGroup = topGroup.prev
      }
      topTree = topTree.prev
    }
    str += "</pwlist>\n"
  }

  return str
}

func (t *TreeDB) Size() int {
  return t.size
}

func (t *TreeDB) GetName() string {
  return t.name
}
