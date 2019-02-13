package config

type CookieConfig struct {
	Name string
}

func GetCookieConfig() *CookieConfig {
	return &CookieConfig{
		Name: "jxc_cookie_session",
	}
}
