package main

import (
	"fmt"
	"strconv"
)

func main() {
	arlot := Hero{
		Name:     "Arlot",
		Damage:   20,
		Cooldown: 3,
	}

	layla := Hero{
		Name:     "Layla",
		Damage:   35,
		Cooldown: 1,
	}

	grock := Hero{
		Name:     "Grock",
		Damage:   5,
		Cooldown: 4,
	}

	for i := 1; i <= 15; i++ {
		fmt.Println("Perulangan ke-" + strconv.Itoa(i))

		fmt.Println(arlot.Attack(i))
		fmt.Println(layla.Attack(i))
		fmt.Println(grock.Attack(i))
	}

}
