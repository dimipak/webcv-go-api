package config

var (
	G_APP      App
	G_MAIL     Maill
	G_DATABASE Database
	G_STORAGE  Storage
)

func InitGlobals() {
	setValues(
		&G_APP,
		&G_MAIL,
		&G_DATABASE,
		&G_STORAGE,
	)
}
