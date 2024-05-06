# Vuln

## Type de vulnérabilité

- `XSS`
- `SQL Injection`
- `CSRF`
- `Command injection`

### XSS (DOM)

- Qu'est ce qu'une faille XSS ? 

Une faille XSS (Cross-Site Scripting) est un type de vulnérabilité que l'on peut trouver dans les applications web. Elle permet à un attaquant d'injecter du code malveillant, généralement du JavaScript, dans les pages web consultées par les autres utilisateurs. Ce code peut ensuite être exécuté par le navigateur de la victime, permettant ainsi à l'attaquant de voler des informations sensibles, comme des cookies de session, ou de réaliser des actions au nom de l'utilisateur.

Il existe deux types principaux de failles XSS :

1. XSS persistant (ou stocké) : Le code malveillant est stocké de manière permanente dans les pages web, par exemple dans une base de données. Il sera exécuté chaque fois que la page sera consultée.

2. XSS non persistant (ou réfléchi) : Le code malveillant est injecté temporairement, généralement via une URL modifiée. Il ne sera exécuté que lorsque la victime cliquera sur un lien menant vers cette URL spécifique.


L'url peut être modifier afin d'inclure une balise script et executer du code non souhaité par l'administrateur du site web.
![URL d'une faille XSS](/doc/images/xss(dom).png)

Résultat :
![Réussite faille XSS](/doc/images/xss-dom-passed.png)

On peux voir ici que l'alerte écrit en javascript qui affiche `a` s'affiche bien sur la page.

Sur notre site, on peux voir que l'on peut essayer la même chose dans l'url 

![CodeQuarryXSS](/doc/images/codequarry-xss-tentative.png)

Mais une fois que l'on envoie la requête, on peux voir que la faille XSS n'est pas exploiter, car à cette url, le backend du site attend uniquement un ID 

![Missing ID](/doc/images/missing_id.png)

Afin que la faille XSS soit "fix", nous avons utiliser dans le code des fonctions afin de `sanitize` notre URL de la façon suivante : 

```go
// ContainsXSS checks if the input string contains common XSS attack vectors
func ContainsXSS(input string) bool {
	// Patterns to check for common XSS contexts:
	// - <script> tags
	// - 'javascript:' pseudo-protocol
	// - HTML event handlers (e.g., onclick)
	// - <iframe> tags
	// - <img> tags with src='data:'
	patterns := []string{
		`(?i)<script.*?>.*?</script>`,
		`(?i)javascript:`,
		`(?i)on\w+="[^"]*"`,
		`(?i)<iframe.*?>.*?</iframe>`,
		`(?i)<img.*?src=['"]data:`,
	}

	// Compile and match each pattern
	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		if re.MatchString(input) {
			return true
		}
	}

	return false
}
```

Cette fonction Go, `ContainsXSS`, vérifie si une chaîne d'entrée contient des vecteurs d'attaque XSS courants en utilisant des expressions régulières. Elle définit un tableau de motifs (patterns) correspondant à des contextes XSS fréquents, tels que les balises `<script>`, le pseudo-protocole `javascript:`, les gestionnaires d'événements HTML (comme `onclick`), les balises `<iframe>` et les balises `<img>` avec `src='data:'`.
