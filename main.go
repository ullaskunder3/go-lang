package main

import (
	"fmt"
	"go-lang/adventure"
	"math"
)

func main() {
	// 🌟 Main character introduction: haruX (ullas), the Go hero! 🌟
	// haruX is a young adventurer who just started his journey to learn Go and become a coding master! 💻⚔️
	fmt.Println("🌟 haruX: Hello, I am haruX (ullas), the Go hero! I'm ready to learn the basics of Go programming! 🌸")

	// ⚡️ Variables: haruX's magical weapons ⚡️
	// haruX can store different types of magical energy, like numbers, words, and more!
	// First, we introduce haruX's powerful weapon variables: integers, floats, and strings!
	var health int = 100               // haruX's health (integer type)
	var magicPower float64 = 12.5      // haruX's magic power (floating-point type)
	var name string = "haruX the Hero" // haruX's name (string type)

	// haruX is checking his powers!
	fmt.Println("🌟 haruX checks his stats:")
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Health: %d\n", health)
	fmt.Printf("Magic Power: %.2f\n", magicPower)

	// 🌈 Constants: haruX’s invincible shield 🌈
	// haruX also has an unchanging, magical shield: the constant!
	const maxHealth = 200 // haruX's max health (cannot change during the adventure)
	fmt.Printf("⚡️ haruX’s maximum health is: %d\n", maxHealth)

	// 🏃‍♂️ Loops: haruX encounters a dungeon 🏰
	// While adventuring, haruX must clear multiple rooms in a dungeon. Let's use a loop to simulate his journey through 5 rooms!
	fmt.Println("\n🏰 haruX enters the dungeon with 5 rooms. Let's see how many rooms he clears!")
	for room := 1; room <= 5; room++ {
		fmt.Printf("Room %d: haruX fights off enemies and clears the room! 💥\n", room)
	}

	// 💥 Arrays: haruX’s list of potions 💊
	// haruX needs potions for each room! He stores his potions in an array (a collection of similar items).
	potions := [3]string{"Health Potion", "Mana Potion", "Speed Potion"}
	fmt.Println("\n💊 haruX checks his potion inventory:")
	for i := 0; i < len(potions); i++ {
		fmt.Printf("Potion %d: %s\n", i+1, potions[i])
	}

	// 🌀 Slices: haruX discovers a special power-up! ⚡️
	// Slices are like arrays, but they can grow! haruX finds new potions in the dungeon.
	// Slices are dynamic arrays!
	powerUps := []string{"Speed Boost", "Shield Boost", "Critical Hit"}
	powerUps = append(powerUps, "Mana Regeneration") // haruX finds an extra power-up!
	fmt.Println("\n⚡️ haruX receives a special power-up list:")
	for _, powerUp := range powerUps {
		fmt.Println(powerUp)
	}

	// ⚡️ Functions: haruX learns to cast spells 🪄
	// haruX learns a function to cast a spell!
	spellName := "Fireball"
	damage := 50
	castSpell(spellName, damage)

	// ✨ Conditional Statements: haruX faces a boss! 🏹
	// haruX must decide whether to fight or run away based on his health.
	fmt.Println("\n⚔️ haruX faces a mighty boss!")
	if health < 50 {
		fmt.Println("haruX is too weak to fight! He runs away! 🏃‍♂️")
	} else {
		fmt.Println("haruX fights the boss bravely! 💪")
	}

	// 🔢 Math operations: haruX uses his magic powers 🔮
	// haruX has a magic orb that can help him calculate things like distances or powers.
	distance := 10.0
	angle := 45.0
	radius := 5.0
	calculatePower(distance, angle, radius)

	// 🦸‍♂️ Structs: haruX’s Character Profile 🧑‍🦱
	// haruX is a complex character with multiple attributes. Let's use a struct to represent him!
	haruX := Hero{
		Name:       "haruX",
		Health:     100,
		MagicPower: 50.5,
	}
	fmt.Println("\n🦸‍♂️ haruX's detailed profile:")
	fmt.Printf("Name: %s\n", haruX.Name)
	fmt.Printf("Health: %d\n", haruX.Health)
	fmt.Printf("Magic Power: %.2f\n", haruX.MagicPower)

	adventure.Adventure()
}

// 🪄 Function to cast spells 🔮
func castSpell(spell string, damage int) {
	// haruX casts a spell with the given damage!
	fmt.Printf("\n💥 haruX casts a %s spell with %d damage! ✨\n", spell, damage)
}

// 🔢 Function to calculate magic power using distance and angle 🧠
func calculatePower(distance float64, angle float64, radius float64) {
	// Magic power is calculated using some powerful math formulas!
	power := math.Sqrt(math.Pow(distance, 2) + math.Pow(radius, 2))
	fmt.Printf("\n✨ haruX calculates his magic power: %.2f (using distance %.2f and radius %.2f)\n", power, distance, radius)
}

// 🦸‍♂️ Struct representing haruX's profile 📜
type Hero struct {
	Name       string
	Health     int
	MagicPower float64
}
