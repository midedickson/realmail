package realmail

func buildRouteDependency() map[string][]string {
	routeDependency := make(map[string][]string)
	routeDependency["regex"] = []string{}
	routeDependency["mx"] = []string{"regex"}
	routeDependency["mx_blacklist"] = []string{"mx"}
	routeDependency["smtp"] = []string{"smtp"}

	return routeDependency
}
