package racers

func SessionType(id int) string {
	switch id {
	case 170:
		return "Practice"
	case 293:
		return "Qualifying"
	default:
		return "-"
	}
}

func Car(id int) string {
	switch id {
	case 1:
		return "Skippy"
	case 4:
		return "Pro Mazda"
	case 55:
		return "BMW Z4 GT3"
	case 72:
		return "Mercedes AMG GT3"
	case 74:
		return "Formula Renault 2.0"
	case 91:
		return "VW Beetle GRC"
	case 93:
		return "Ferrari 488 GTE"
	case 94:
		return "Ferrari 488 GT3"
	default:
		return "-"
	}
}

func Track(id int) string {
	switch id {
	case 46, 99, 100:
		return "Barber"
	case 47, 158:
		return "Laguna Seca"
	case 145, 146:
		return "Brands Hatch"
	case 147, 148, 149, 150, 151:
		return "Zandvoort"
	case 152:
		return "Phillip Island"
	case 163, 164, 165:
		return "Spa-Francorchamps"
	case 199:
		return "Zolder"
	case 212, 213:
		return "Interlagos"
	case 239, 240, 241, 242, 243, 244, 245, 246, 247:
		return "Monza"
	case 249, 253:
		return "Nordschleife"
	case 266, 267:
		return "Imola"
	case 297, 298, 299:
		return "Snetterton"
	default:
		return "-"
	}
}
