## Golang HTTPS Server


### HTTPS PART 
- Pourquoi [HTTPS](https://en.wikipedia.org/wiki/HTTPS) ?

Le protocole `HTTPS` est utilisé à la place du protocole HTTP pour garantir la sécurité des données échangées entre un navigateur web et un serveur web. `HTTPS` crypte ces données, ce qui rend plus difficile pour les pirates d'intercepter ou de modifier des informations sensibles telles que les mots de passe, les détails des cartes de crédit, etc. Une telle attaque peut être utilisée, par exemple, avec [MITM attack] (https://en.wikipedia.org/wiki/Man-in-the-middle_attack).


- Comment ? :

[Lien vers le tutoriel](https://gist.github.com/denji/12b3a568f092ab951456)

Générer une clé privée (server.key) :

```bash
openssl genrsa -out server.key 2048
```

Pourquoi ? : La clé privée est essentielle pour établir des connexions sécurisées entre le serveur web et les clients (comme les navigateurs).

Importance : Cette clé est un secret critique qui doit être gardé en sécurité. Elle permet de chiffrer les données envoyées par le serveur et de déchiffrer les données envoyées par les clients. Si cette clé était compromise, les attaquants pourraient déchiffrer les données confidentielles transitant entre le serveur et les clients, ce qui pourrait compromettre la sécurité et la confidentialité des utilisateurs.

Générer un certificat auto-signé (server.crt) basé sur la clé privée :

```bash
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

Pourquoi ? : Le certificat est utilisé pour prouver l'identité du serveur web aux clients. C'est essentiel pour établir la confiance et garantir que les utilisateurs se connectent au bon site.

Importance : Un certificat SSL/TLS valide est nécessaire pour établir une connexion sécurisée et chiffrée. Ce certificat contient des informations sur le serveur (comme son nom de domaine) et est signé numériquement avec la clé privée. Bien qu'un certificat auto-signé ne soit pas aussi fiable qu'un certificat signé par une autorité de certification reconnue, il est tout de même utile pour tester et développer des applications.

- Vu que nous utilisons désormais `HTTPS`, et donc le port 443, il faut lancer le serveur en tant que super utilisateur. Les ports inférieurs à 1024 sont réservés aux processus exécutés par le super utilisateur.

```bash
sudo go run .
```

### HEADERS PART

Afin de rendre notre serveur plus sûr, il est conseillé d'ajouter [des en-têtes sécurisés à notre serveur](https://wiki.owasp.org/index.php/OWASP_Secure_Headers_Project). Cela permet d'atténue plusieurs vulnérabilités.

Reprenons le code et vérifions les en-têtes que nous avons ajoutés, ce contre quoi ils nous protègent et comment nous les avons définis.

`w.Header().Add("Strict-Transport-Security", "max-age=63072000 ; includeSubdomains")`
 
Nous définissons ici l'en-tête "HSTS" (HTTP Strict Transport Security). Cela indique que les clients doivent automatiquement interagir avec le serveur en utilisant des connexions HTTPS. En définissant l'en-tête de réponse, nous spécifions la période pendant laquelle le client ne doit accéder au serveur que de manière sécurisée. Il protège contre les "attaques par rétrogradation du protocole" et le "détournement de cookies". (Wikipedia 2020c; OWASP 2020b) La valeur max-age spécifie la durée pendant laquelle le client doit se souvenir que le site doit être accessible par HTTPS. La valeur includeSubdomains implique que cette règle s'applique également à tous les sous-domaines.

`w.Header().Add("Content-Security-Policy", "default-src 'self'")`

L'en-tête "CSP" (Content Security Policy) vous permet de définir une politique concernant les ressources que le client est autorisé à charger/exécuter. Il spécifie les domaines que le client doit considérer comme des sources valides. Supposons que vous ayez une application web qui autorise le chargement de contenu provenant d'une autre source. Avec CSP, vous pouvez mettre sur liste blanche les origines des scripts, des images, des polices, des feuilles de style, etc. Il protège contre certaines "attaques de script intersites" (XSS), qui exploitent la confiance des clients dans le contenu reçu du serveur. Des scripts malveillants peuvent ainsi être exécutés par le client parce qu'il fait confiance à la source du contenu, même s'il ne provient pas de là où il semble être. La politique "default-src 'self'" définit que tout le contenu doit provenir de l'origine du site, ce qui exclut les sous-domaines.

`w.Header().Add("X-XSS-Protection", "1 ; mode=block")`
 
L'en-tête "X-XSS-Protection" empêche les pages de se charger lorsqu'elles détectent des attaques XSS (Cross-site scripting). Dans ce cas, l'attaquant fait en sorte qu'une page charge du JavaScript malveillant, en envoyant la charge utile malveillante dans le cadre de la demande. Cet en-tête permet de spécifier si la page doit être chargée ou bloquée. La valeur "1 ; mode=block ;" bloque complètement le chargement de la page. Il s'agit d'un en-tête spécial pour les navigateurs Chrome et Internet Explorer, qui semble inutile pour les API et devrait être couvert par une politique de sécurité du contenu restrictive.

`w.Header().Add("X-Frame-Options", "DENY")`
 
L'en-tête "X-Frame-Options" indique au navigateur s'il est autorisé à rendre une page dans une balise <embed>,<frame >, <iframe> ou <object>. Cela permet d'éviter le "click-jacking", c'est-à-dire que le contenu n'est pas intégré dans d'autres sites. Ici, nous avons choisi la valeur "DENY", ce qui empêche le contenu d'être intégré dans d'autres pages.

`w.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")`

L'en-tête "Referrer-Policy" permet de définir les informations envoyées aux sites/ressources externes dans l'en-tête de requête Referer. Lorsqu'un client accède à une URL à partir d'un lien hypertexte ou qu'une page web charge une ressource externe, le navigateur ajoute l'en-tête Referer (oui, c'est intentionnellement mal orthographié), pour indiquer à la destination l'origine de cette demande. Imaginez que votre page web contienne un lien vers un site externe, et que ce site externe reçoive l'en-tête Referer avec des informations qui ne devraient être utilisées qu'en interne. Cet en-tête nous permet de contrôler ce qui est envoyé à la destination. Dans cet exemple, nous allons le définir sur strict-origin-when-cross-origin, qui envoie le referer complet aux sources de la même origine, et l'url sans chemin aux destinations d'origine externe, et n'envoie pas d'en-tête aux destinations moins sécurisées (HTTPS→HTTP).

`w.Header().Add("X-Content-Type-Options", "nosniff")`

L'en-tête "X-Content-Type-Options" indique les types MIME définis par l'en-tête "Content-Type", qui ne doivent pas être modifiés et doivent être respectés. Il empêche les navigateurs d'interpréter les fichiers comme étant d'un autre type que celui spécifié dans l'en-tête "Content-Type". Il exclut le "reniflage de type Mime", qui consiste à deviner le type Mime correct en examinant les octets de la ressource par le navigateur. Sans cet en-tête, les navigateurs peuvent incorrectement détecter les fichiers comme étant des scripts et des feuilles de style, ce qui conduit à des attaques XSS. La définition de l'en-tête à "nosniff" indique que les navigateurs doivent empêcher la détection incorrecte de fichiers non scripts comme étant des scripts.

`w.Header().Add("Content-Type", "text/plain ; charset=UTF-8")`
 
Pour que 'X-Content-Type-Options' fonctionne correctement, nous devons définir le 'Content-Type' sur le bon type MIME. Pour cet exemple, nous allons le définir à "text/plain ; charset=UTF-8". Cela indique au client que le type est text, que le sous-type est plain et que l'encodage des caractères est utf-8.
