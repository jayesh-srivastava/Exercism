package space

type Planet string

const oneearthyearinsec = 31557600

func Age(sec float64, planet Planet) float64 {
	switch planet {
	case "Earth":
		return sec / oneearthyearinsec
	case "Mercury":
		return sec / (oneearthyearinsec * 0.2408467)
	case "Venus":
		return sec / (oneearthyearinsec * 0.61519726)
	case "Mars":
		return sec / (oneearthyearinsec * 1.8808158)
	case "Jupiter":
		return sec / (oneearthyearinsec * 11.862615)
	case "Saturn":
		return sec / (oneearthyearinsec * 29.447498)
	case "Uranus":
		return sec / (oneearthyearinsec * 84.016846)
	case "Neptune":
		return sec / (oneearthyearinsec * 164.79132)
	default:
		return 0
	}

}
