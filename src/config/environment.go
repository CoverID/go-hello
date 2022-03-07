package config

import (
	"os"
)

func ConfigureAPM() {
	os.Setenv("ELASTIC_APM_SERVICE_NAME", "Golang-App")
	os.Setenv("ELASTIC_APM_SERVER_URL", "http://server:8200")
	os.Setenv("ELASTIC_APM_ENVIRONMENT", "development")
	os.Setenv("ELASTIC_APM_SECRET_TOKEN", "YXJKWWtOVFVLZjY0T0V4RDl5Rm5HdThWQVNnMlJMTUhRMXczb3FsYmpocEJlczcwV0l6bXRYWnZQNUNkaWM=")
	os.Setenv("ELASTIC_APM_CAPTURE_BODY", "all")
}
