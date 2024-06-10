# LOGGING 

### Pourquoi ?

- Débogage et diagnostic des erreurs : Les journaux peuvent aider à comprendre ce qui s'est passé lorsqu'une erreur survient. Enregistrer des détails sur les opérations effectuées peut faciliter le processus de débogage en fournissant des informations sur l'état de l'application au moment où une erreur s'est produite.

- Surveillance et performance : Les journaux peuvent aider à surveiller les performances de l'application en enregistrant des informations telles que les temps de réponse des requêtes, les erreurs fréquentes, etc. Cela peut aider à identifier les goulots d'étranglement et à optimiser les performances de l'application.

- Audit et conformité : En enregistrant les actions des utilisateurs, comme les enregistrements et les connexions réussies, nous pouvons conserver une trace de ce qui se passe dans l'application. Cela peut être important pour des raisons de sécurité, de conformité aux réglementations et de traçabilité des actions des utilisateurs.

### Comment ? 

Tout d'abord, nous déclarons plusieurs niveaux de journalisation possibles. Ceci est très utile car, dans certains cas, comme par exemple un crash serveur, il est crucial de distinguer une erreur majeure d'une simple information de débogage.

```go
type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)
```

Ce code définit un type `LogLevel` pour représenter différents niveaux de journalisation. En utilisant `iota`, chaque niveau de journalisation est automatiquement attribué une valeur numérique croissante à partir de zéro. Cela simplifie l'ajout de nouveaux niveaux et garantit la cohérence dans tout le code. Utiliser un type distinct aide à éviter les erreurs de type et rend le code plus clair.

Ensuite, viens notre fonction principale de `logging` : 

```go
func Log(level LogLevel, message string) {
	file, err := os.OpenFile("audit.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open audit.log: %s", err)
	}
	defer file.Close()

	var prefix string
	switch level {
	case DebugLevel:
		prefix = "DEBUG: "
	case InfoLevel:
		prefix = "INFO: "
	case WarnLevel:
		prefix = "WARN: "
	case ErrorLevel:
		prefix = "ERROR: "
	}

	logger := log.New(file, prefix, log.LstdFlags)
	logger.Println(message)
}
```

Cette fonction `Log` écrit des messages dans un fichier de journal appelé "audit.log" avec un niveau de gravité spécifié. Elle ouvre le fichier, ajoute un préfixe correspondant au niveau de gravité du message, crée un logger avec le fichier et le préfixe, puis écrit le message dans le fichier. 