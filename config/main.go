package config

type Config interface {
	setValues()
}

func setValues(c ...Config) {
	for _, v := range c {
		v.setValues()
	}
}
