package structs

// Platzi es una interfaz de PlatziCourse y PlatziCareer
type Platzi interface {
	Subscribe(name string)
}

func InterfaceTest() {
	platziCourse := PlatziCourse{Name: "Go", Slug: "go", Skills: []string{"backend"}}
	platziCareer := new(PlatziCareer)
	platziCareer.Name = "GoCareer"
	platziCareer.Slug = "go"
	callSubscribe(platziCareer)
	callSubscribe(platziCourse)
}

func callSubscribe(p Platzi) {
	p.Subscribe("Yohan")
}
