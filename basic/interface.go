package basic

import (
	"fmt"
)

// easy one:
type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func saySomething(s Speaker) {
	fmt.Println(s.Speak())
}

// === Part 1: Basic Guild Contract ===
type Adventurer interface {
	Fight()
}
type Ninja struct{}
type Samurai struct{}

func (n Ninja) Fight() {
	fmt.Println("Ninja strikes silently!")
}
func (s Samurai) Fight() {
	fmt.Println("Samurai charges with honor!")
}

func ArenaBattle(a Adventurer) {
	a.Fight()
}

// === Part 2: Empty Interface (Universal Spirit Orb) ===
func InspectPower(v interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", v, v)
}

// === Part 3: Interface as Dependency (Strategy Pattern) ===
type BattleStrategy interface {
	Execute()
}
type Aggressive struct{}
type Defensive struct{}

func (Aggressive) Execute() {
	fmt.Println("Attacks with full power!")
}
func (Defensive) Execute() {
	fmt.Println("Blocks and counterattacks!")
}

type HeroStruct struct {
	Name     string
	Strategy BattleStrategy
}

func (h HeroStruct) Fight() {
	fmt.Print(h.Name, " ")
	h.Strategy.Execute()
}

// === Part 4: Multiple Method Interfaces ===
type UltimateFighter interface {
	Fight()
	Heal()
}
type Paladin struct{}

func (Paladin) Fight() { fmt.Println("Paladin smites evil!") }
func (Paladin) Heal()  { fmt.Println("Paladin heals party!") }

// === Part 5: Pointer Receivers ===
type Trainer interface {
	Train()
}
type Warrior struct {
	Level int
}

func (w *Warrior) Train() {
	w.Level++
	fmt.Println("Training done. New level:", w.Level)
}

// === Part 6: Interface Composition (Trait Fusion) ===
type Fighter interface{ Fight() }
type Healer interface{ Heal() }
type HolyKnight interface {
	Fighter
	Healer
}
type Saint struct{}

func (Saint) Fight() { fmt.Println("Saint pierces evil with light!") }
func (Saint) Heal()  { fmt.Println("Saint restores all HP!") }

// STRAT
func InterfaceUHWork() {
	d := Dog{}
	c := Cat{}
	saySomething(d) // prints "Woof!" method take d instance call speak
	saySomething(c) // prints "Meow!" method take c instance call speak

	fmt.Println("=== how interface work under the hood ===")

	fmt.Println("\n=== Part 1: Basic Guild Contract ===")
	ArenaBattle(Ninja{})
	ArenaBattle(Samurai{})

	fmt.Println("\n=== Part 2: Empty Interface (Universal Spirit Orb) ===")
	InspectPower("Saitama")
	InspectPower(9999)
	InspectPower(true)

	fmt.Println("\n=== Part 3: Interface as Dependency (Strategy Pattern) ===")
	hero1 := HeroStruct{"Guts", Aggressive{}}
	hero2 := HeroStruct{"Levi", Defensive{}}
	hero1.Fight()
	hero2.Fight()

	fmt.Println("\n=== Part 4: Multiple Method Interfaces ===")
	var p UltimateFighter = Paladin{}
	p.Fight()
	p.Heal()

	fmt.Println("\n=== Part 5: Pointer Receivers ===")
	// This won't work: var t Trainer = Warrior{} (invalid)
	w := &Warrior{}
	var t Trainer = w
	t.Train()
	t.Train()

	fmt.Println("\n=== Part 6: Interface Composition (Trait Fusion) ===")
	var hk HolyKnight = Saint{}
	hk.Fight()
	hk.Heal()
}
