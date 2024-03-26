package chat

type role string

func Role() role{
  return role("")
}

func (r role) String() string{
  return string(r)
}

func (r role) System() role{
  r = "system"
  return r
}

func (r role) User() role{
  r = "user"
  return r
}

func (r role) Assistant() role{
  r = "assistant"
  return r
}
