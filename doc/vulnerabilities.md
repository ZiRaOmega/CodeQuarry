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

Sur notre site, on peux voir que l'on peut essayer la même chose dans l'url 

![CodeQuarryXSS](/doc/images/codequarry-xss-tentative.png)

Mais une fois que l'on envoie la requête, on peux voir que la faille XSS n'est pas exploiter, car à cette url, le backend du site attend uniquement un ID 

![Missing ID](/doc/images/missing_id.png)