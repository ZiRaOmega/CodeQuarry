# Rate limiter 

- Pourquoi ?

Le `rate limiting` (ou limitation de débit) est une technique utilisée pour contrôler le nombre de requêtes qu'un client peut envoyer à un serveur dans un laps de temps donné. Cette technique permet de prévenir les abus, les attaques par déni de service (DoS) et d'autres comportements indésirables qui peuvent surcharger le serveur et affecter ses performances.

- Comment ?

```go
// RateLimiter est un struct qui gère la limitation de débit des requêtes HTTP
type RateLimiter struct {
	mu         sync.Mutex
	rateLimit  int
	timeWindow time.Duration
	counters   map[string]int
	timestamps map[string]time.Time
}

// NewRateLimiter initialise un nouveau RateLimiter avec les paramètres donnés
func NewRateLimiter(rateLimit int, timeWindow time.Duration) *RateLimiter {
	return &RateLimiter{
		rateLimit:  rateLimit,
		timeWindow: timeWindow,
		counters:   make(map[string]int),
		timestamps: make(map[string]time.Time),
	}
}

// Handle est une méthode qui prend en charge la limitation de débit pour chaque requête HTTP entrante
func (r *RateLimiter) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ip := req.RemoteAddr
		r.mu.Lock()
		counter, ok := r.counters[ip]
		if !ok {
			counter = 0
		}
		timestamp, ok := r.timestamps[ip]
		if !ok {
			timestamp = time.Now()
		}
		r.mu.Unlock()

		if time.Since(timestamp) >= r.timeWindow {
			r.mu.Lock()
			delete(r.counters, ip)
			delete(r.timestamps, ip)
			r.mu.Unlock()
		} else if counter >= r.rateLimit {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		} else {
			r.mu.Lock()
			r.counters[ip]++
			r.timestamps[ip] = time.Now()
			r.mu.Unlock()
		}
		next.ServeHTTP(w, req)
	})
}
```

Ce code crée un `middleware` HTTP qui limite le nombre de requêtes qu'un client peut envoyer en utilisant son adresse IP. Le `middleware` utilise une mutex pour synchroniser l'accès aux compteurs et aux timestamps, et vérifie si le nombre de requêtes dépasse la limite autorisée dans la fenêtre de temps donnée. Si la limite est dépassée, le `middleware` renvoie une réponse HTTP avec le code d'état "429 Too Many Requests". Sinon, le `middleware` incrémente le compteur de requêtes pour l'adresse IP donnée et met à jour le timestamp.

Pour utiliser ce `middleware`, vous pouvez l'ajouter à votre chaîne de `middlewares` HTTP avant vos handlers :

```go
rateLimiter := app.NewRateLimiter(10, time.Minute)
http.Handle("/login", rateLimiter.Handle(http.HandlerFunc(loginHandler)))
```

Cet exemple limite le nombre de requêtes à 10 par minute pour l'endpoint `/login`.

Il est important de choisir une limite de débit appropriée en fonction des besoins de votre application et de votre infrastructure. Une limite trop basse peut affecter l'expérience utilisateur, tandis qu'une limite trop élevée peut ne pas offrir une protection suffisante contre les abus et les attaques. Vous pouvez également envisager d'utiliser des limites différentes pour différents endpoints en fonction de leur criticité et de leur utilisation prévue.