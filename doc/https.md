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

lemon was here.
