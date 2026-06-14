package llm

type LLMAdapter interface {
	Invoke(prompt *string, substitutions *map[string]string) string
}
