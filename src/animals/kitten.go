package animals

type Kitten struct {
	Name    string
	Hobbies []string
	Likes   int
}

func (k *Kitten) SetName(name string) {
	k.Name = name
}

func (k *Kitten) GetName() string {
	return k.Name
}

func (k *Kitten) MakeNoise() string {
	return "meow"
}

func (k *Kitten) SetHobbies(hobbies []string) {
	k.Hobbies = hobbies
}

func (k *Kitten) GetHobbies() []string {
	return k.Hobbies
}

func (k *Kitten) IncrementLikecounter() {
	k.Likes += 1
}

func (k *Kitten) GetNumberOfLikes() int {
	return k.Likes
}
