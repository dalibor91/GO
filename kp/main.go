package main

import  (
  "fmt"
  "./manager"
)

func main() {
  db := data.NewTreeDB("test")

  tree1 := db.AddTree(data.NewTree("Tree 1"))

  tree2 := db.AddTree(data.NewTree("Tree 2"))


  for i:=0; i<10; i++ {
    tree1.AddGroup(fmt.Sprintf("Group in Test1 No.%d", i))
  }

  for i:=10; i<20; i++ {
    tree2.AddGroup(fmt.Sprintf("Group in Test2 No.%d", i))
  }


  grp1, _ := tree1.GetGroup(3)
  grp2, _ := tree2.GetGroup(5)


  for i:=20; i<50; i++ {
    record := data.NewEntry()
    record.SetTitle(fmt.Sprintf("Record %d", i))
    grp1.AddEntry(record)
  }

  for i:=50; i<90; i++ {
    record := data.NewEntry()
    record.SetTitle(fmt.Sprintf("Record %d", i))
    grp2.AddEntry(record)
  }

  fmt.Printf("Tree db '%s' size: %d\n----------------------------------\n", db.GetName(), db.Size())
  fmt.Printf("Tree1 '%s' size: %d\n", tree1.GetName(), tree1.Size())
  fmt.Printf("Tree2 '%s' size: %d\n-----------------------------------\n", tree2.GetName(), tree2.Size())

  fmt.Println("--------------------------\nTree1\n--------------------------")
  tree1.PrintGroups()
  fmt.Println("--------------------------\nTree2\n--------------------------")
  tree2.PrintGroups()

  fmt.Println("--------------------------\nGrop 3 in tree1 records\n--------------------------")
  grp1.PrintEntries()
  fmt.Println("--------------------------\nGrop 5 in tree2 records\n--------------------------")
  grp2.PrintEntries()

  fmt.Println(db.XML())
}
