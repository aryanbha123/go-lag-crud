package services

func HealthService(name string) string {
	if name == "" {
		return "Helath check OK"
	}
	return "Health check passed , User detected naming : " + name
}
