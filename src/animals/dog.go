package animals

type Dog struct {
     Name string
}

func (d *Dog) SetName(name string) {
     d.Name = name
}

func (d *Dog) GetName() string {
     return d.Name
}

func (d *Dog) MakeNoise() string {
  return "bark"
}
