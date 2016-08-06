package treedata

type Entry struct {
  title string
  username string
  url string
  password string
  notes string
  createTime string
  lastAccessTime string
  expiresTime string
  expires bool
  next *Entry
  prev *Entry
  group *Group
}

func NewEntry() *Entry {
  return new(Entry)
}

func (e *Entry) SetTitle(title string) *Entry {
  e.title = title
  return e;
}

func (e *Entry) GetTitle() string {
  return e.title
}

func (e *Entry) SetUsername(username string) *Entry {
  e.username = username
  return e;
}

func (e *Entry) GetUsername() string {
  return e.username
}

func (e *Entry) SetURL(url string) *Entry {
  e.url = url
  return e;
}

func (e *Entry) GetURL() string {
  return e.url
}

func (e *Entry) SetPassword(password string) *Entry {
  e.password = password
  return e;
}

func (e *Entry) GetPassword() string {
  return e.password
}

func (e *Entry) SetNotes(notes string) *Entry {
  e.notes = notes
  return e;
}

func (e *Entry) SetCreateTime(createTime string) *Entry {
  e.createTime = createTime
  return e;
}
