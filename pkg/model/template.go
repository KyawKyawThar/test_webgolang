package model

type TemplateData struct {
	StringMap map[string]string
	FloatMap  map[string]float32
	IntMap    map[string]int
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
