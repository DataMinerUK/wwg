package animals

type Pet interface {
	SetName(name string)
	GetName() string
	SetHobbies(hobbies []string)
	GetHobbies() []string
	MakeNoise() string
	IncrementLikecounter()
	GetNumberOfLikes() int
}
