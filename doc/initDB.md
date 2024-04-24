# Installation de PostgreSQL

>Ce document guide l'installation de PostgreSQL sur les systèmes d'exploitation Ubuntu et macOS, décrit les informations d'identification pour se connecter à la base de données et explique pourquoi PostgreSQL est préférable à SQLite3 et MySQL pour certaines applications.

## Installation sur Ubuntu

1. **Mise à jour de la liste des paquets**:  
   Ouvrez un terminal et exécutez:  
   
```bash
sudo apt update
```

2. Installation de PostgreSQL:  
    *Toujours dans le terminal, installez PostgreSQL ainsi que le package contrib qui ajoute des fonctionnalités supplémentaires*  

```bash
sudo apt install postgresql postgresql-contrib
```

3. Démarrage du service PostgreSQL:  
    *Pour démarrer le service PostgreSQL*  

```bash 
sudo systemctl start postgresql
```

4. Activation du démarrage automatique:  
    *Pour que PostgreSQL démarre automatiquement au lancement du système*  

```bash
sudo systemctl enable postgresql
```

5. Création de l'utilisateur 'codequarry' dans postgreSQL  

```bash
sudo -u postgres psql
```

6. Créez l'utilisateur codequarry avec le mot de passe CQ1234:  

```bash
CREATE USER codequarry WITH PASSWORD 'CQ1234';
```

7. Création d'une base de données:
    *Toujours dans l'interface de commande PostgreSQL:*  

```bash
CREATE DATABASE codequarrydb OWNER codequarry;
```

8. Attribution des droits:
    *Donnez tous les droits sur la base de données à codequarry:*  

```bash
GRANT ALL PRIVILEGES ON DATABASE codequarrydb TO codequarry;
```

9. Lancer le projet:
    *Lorsque vous aurez suivi toutes ces étapes vous pourrez lancer votre server.*  

```bash
go run .
```


# Pourquoi choisir PostgreSQL?

- ### *Par rapport à SQLite3:*

1. Gestion multi-utilisateur: Contrairement à SQLite3 qui est conçu pour des applications légères et un accès principalement local, PostgreSQL gère le multi-utilisateur et le traitement de transactions complexes, ce qui est essentiel pour les applications d'entreprise.

2. Sécurité et conformité: PostgreSQL offre une meilleure conformité aux normes SQL et prend en charge des fonctionnalités de sécurité avancées.

- ### *Par rapport à MySQL:*

1. Conformité aux normes SQL: PostgreSQL est souvent plus conforme aux normes SQL que MySQL, ce qui facilite le portage d'applications entre différentes bases de données.

2. Fonctionnalités avancées: PostgreSQL inclut des fonctionnalités comme les index BRIN ou GIN et le support natif pour les données JSON qui sont mieux intégrées et plus performantes que dans MySQL.

3. Extensibilité: L'architecture de PostgreSQL permet une meilleure extensibilité pour les applications complexes et les grandes bases de données.

---

> *Ces caractéristiques font de PostgreSQL un excellent choix pour les entreprises et les développeurs cherchant une solution robuste, conforme et extensible pour la gestion de bases de données.*

## Utilisation de DBeaver pour la visualisation et la manipulation de la base de données

>DBeaver est un outil de gestion de base de données universel gratuit qui permet de se connecter à divers systèmes de bases de données, y compris PostgreSQL. Voici comment vous pouvez l'utiliser pour connecter et manipuler votre base de données PostgreSQL :

#### Téléchargement et Installation de DBeaver:  
*Vous pouvez ouvrir votre app Center sur votre linux pour installer dbeaver.*

#### Configuration de la connexion à la base de données:  

- 1.Ouvrez DBeaver, puis allez dans le menu Database et sélectionnez New Database Connection.  

- 2.Choisissez PostgreSQL comme type de base de données.  

![choosePOSTGRESQL](/images/Capture%20d’écran%20du%202024-04-23%2017-27-09.png)

- 3.Remplissez les détails de la connexion :

- 4.Host : localhost (ou l'adresse IP du  serveur de base de données si elle est distante)  

>Port : 5432  
>Database : codequarrydb  
>Username : codequarry  
>Password : CQ1234  

![FILL](/images/Capture%20d’écran%20du%202024-04-23%2017-26-57.png)

- Cliquez sur Test Connection pour vérifier que tout est configuré correctement.  

- Si le test est réussi, cliquez sur Finish pour sauvegarder la configuration.

#### Utilisation de DBeaver:  

*Une fois connecté, vous pouvez naviguer dans l'arborescence de la base de données sur la gauche pour voir les tables et autres objets de la base de données.  
Pour exécuter des requêtes SQL, ouvrez un nouvel onglet SQL en cliquant droit sur le nom de la base de données et en sélectionnant SQL Editor > New SQL Editor.  
Tapez vos requêtes SQL dans l'éditeur et exécutez-les en cliquant sur l'icône de l'exécution (play) ou en appuyant sur CTRL + Enter.*

#### Conseils pour l'utilisation de DBeaver

- ##### Exploration des données :  

*Utilisez l'onglet Data dans n'importe quelle table pour visualiser, filtrer et trier les données sans écrire de requête SQL.* 

- ##### Modification de la structure de la base de données :  

*Vous pouvez modifier les structures de tables, ajouter des colonnes ou des index directement à partir de l'interface graphique en cliquant droit sur les objets de la base de données et en sélectionnant l'option appropriée.* 

- ##### Sauvegarde et Restauration :  

*DBeaver offre des outils pour exporter et importer des données, ce qui est utile pour la sauvegarde ou la migration de données entre différents environnements.*

> En intégrant DBeaver dans vos outils de gestion de base de données, vous augmentez votre efficacité pour gérer, développer et analyser votre base de données PostgreSQL de manière visuelle et interactive.

