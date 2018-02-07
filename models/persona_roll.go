package models

type PersonaRoll struct {
	Persona *Persona `orm:"column(persona);rel(fk)"`
	Roll    *Roll    `orm:"column(roll);rel(fk)"`
	Id      int      `orm:"column(Id)"`
}
