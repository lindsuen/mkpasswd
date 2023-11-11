package src

import "flag"

func SetParameter(p *Password) {
	flag.Uint64Var(&p.Quantity, "N", 1, "The quantity of created passwords.")
	flag.Uint64Var(&p.Length, "l", 16, "The length of password.")
	flag.Uint64Var(&p.Number, "n", 4, "The number of Arabic numerals (0-9).")
	flag.Uint64Var(&p.Character, "c", 4, "The number of special characters (!@#$%^&*+?).")

	flag.Parse()
}
