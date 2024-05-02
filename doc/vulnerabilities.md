# Vuln

## Type de vulnérabilité

- `XSS`
- `SQL Injection`
- `CSRF`
- `Command injection`

### XSS (DOM)

L'url peut être modifier afin d'inclure une balise script et executer du code non souhaité par l'administrateur du site web.
![URL d'une faille XSS](/doc/images/xss(dom).png)

Résultat :
![Réussite faille XSS](/doc/images/xss-dom-passed.png)

On peux voir ici que l'alerte écrit en javascript qui affiche `a` s'affiche bien sur la page.

this is a test to add