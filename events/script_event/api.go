package script_event

import "strings"

// StartScript получить скрипт из папки /resources/scripts/start/<файл>.sql
func (event *Event) StartScript(script string) string {
	return event.startScripts[script+".sql"]
}

// Script получить скрипт из папки /resources/script/<папка>/<файл>.sql
func (event *Event) Script(folder, script string) string {
	// remove .sql
	script = strings.ReplaceAll(script, ".sql", "")

	// if there is no folder
	if _, ok := event.scripts[folder].(map[string]string); !ok {
		return ""
	}

	// sucess case
	return event.scripts[folder].(map[string]string)[script+".sql"]
}
