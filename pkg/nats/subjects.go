package nats

import "strings"

func cleanSubjectStep(step string) string {
	return strings.ReplaceAll(step, ".", "_")
}

func createEntityIdSubject(prefix, partitionKey, entityType, entityId string) string {
	return prefix +
		".PART." + cleanSubjectStep(partitionKey) +
		".ENT." + cleanSubjectStep(entityType) +
		".ID." + cleanSubjectStep(entityId)
}
