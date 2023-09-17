package config

import(
	"os"
)

func IsProduction() (bool) {
	return os.Getenv("LANGTOOLS_BACKEND_IS_PRODUCTION") == "yes"
}

func HttpPort() (string) {
	if IsProduction() {
		return os.Getenv("LANGTOOLS_BACKEND_HTTP_PORT")
	}
	return ":8080"
}

func HttpsPort() (string) {
	if IsProduction() {
		return os.Getenv("LANGTOOLS_BACKEND_HTTPS_PORT")
	}
	return ":443"
}

func DomainName() (string) {
	if IsProduction() {
		return os.Getenv("LANGTOOLS_BACKEND_DOMAIN_NAME")
	}
	return "localhost"
}

func SqlHost() (string) {
	if IsProduction() {
		return os.Getenv("LANGTOOLS_SQL_HOST")
	}
	return "localhost"
}

func SqlPort() (string) {
	if IsProduction() {
		return os.Getenv("LANGTOOLS_SQL_PORT")
	}
	return "5432"
}

func SqlUsername() (string) {
	if IsProduction() {
		return os.Getenv("LANGTOOLS_SQL_USERNAME")
	}
	return "langtools_user"
}

func SqlPassword() (string) {
	if IsProduction() {
		return os.Getenv("LANGTOOLS_SQL_PASSWORD")
	}
	return "langtools_password"
}

func SqlDatabase() (string) {
	if IsProduction() {
		return os.Getenv("LANGTOOLS_SQL_DATABASE")
	}
	return "langtools"
}