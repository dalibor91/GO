package treedata

import (
  "encoding/xml"
  "os"
  "io/ioutil"
  "fmt"
  "errors"
)


type XMLPWLoader struct {
  List XMLPWList `xml:"pwlist"`
}

type XMLPWList struct {
  Entries []XMLPWEntry `xml:"pwentry"`
}

type XMLPWEntry struct {
  Group XMLPWEntryGroup `xml:"group"`
  Tree string `xml:",attr"`
  Username string `xml:"username"`
  Title string `xml:"title"`
  Url string `xml:"url"`
  Password string `xml:"password"`
  Notes string `xml:"notes"`
  CreateTime string `xml:"creationtime"`
}

type XMLPWEntryGroup struct {
  Name string `xml:",chardata"`
  Tree string `xml:"tree,attr"`
}

func ParseXML(path string) (*TreeDB, error) {
  xmlFile, err := os.Open(path)
  if err != nil {
    return nil, errors.New(fmt.Sprintf("Unable to open %s", path))
  }

  content, err := ioutil.ReadAll(xmlFile)
  if err != nil {
    return nil, errors.New(fmt.Sprintf("Unable to read %s", path))
  }

  defer xmlFile.Close()

  var q XMLPWList
  err = xml.Unmarshal(content, &q)

  if err != nil {
		return nil, errors.New(fmt.Sprintf("Error with parsing xml: %v", err))
	}

  tdb := NewTreeDB(path)

  fmt.Println(q)
  for _, d := range q.Entries {
    if tdb.TreeExists(d.Group.Tree) == false {
      tdb.AddTree(NewTree(d.Group.Tree))
    }

    tree, _ := tdb.GetTreeByName(d.Group.Tree)

    if tree.GroupExists(d.Group.Name) == false {
      tree.AddGroup(d.Group.Name)
    }

    group, _ := tree.GetGroupByName(d.Group.Name)

    e := NewEntry()
    e.SetTitle(d.Title)
    e.SetURL(d.Url)
    e.SetNotes(d.Notes)
    e.SetCreateTime(d.CreateTime)
    e.SetUsername(d.Username)
    e.SetPassword(d.Password)

    group.AddEntry(e)

  }

  return tdb, nil
}
