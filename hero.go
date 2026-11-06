package main

import "strconv"

type Hero struct {
	Name        string
	Damage      int
	Cooldown int
}

func (hero Hero) Attack(counter int) string {
	if counter == 1 || (counter+1)%hero.Cooldown == 0 {
		return hero.Name + " - Damage : " + strconv.Itoa(hero.Damage) + " | " + strconv.Itoa(counter)
	}

	return ""
}
