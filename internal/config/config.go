package config

type Config struct {
	db_url            string
	current_user_name string
}

const configFileName = ".gatorconfig.json"
