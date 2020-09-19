package patterns

import "testing"

func TestDecorator(t *testing.T) {
	reno := Renault{"Рено", "Renault LOGAN Active", 499.0}
	println(AutoBaseToString(reno))

	myreno := MediaNav{reno, "Навигация"}
	println(AutoBaseToString(myreno))

	newmyReno := SystemSecurity{MediaNav{reno, "Навигация"}, "Безопасность"}
	println(AutoBaseToString(newmyReno))
}
