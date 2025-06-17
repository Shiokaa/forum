# Projet Forum - Ynov B1 Informatique

  Bonjour et bienvenue sur le dépôt de notre projet de fin d'année !<br/>
  Ce projet a été réalisé dans le cadre du module <strong>"Forum"</strong><br/>
  pour valider notre première année de Bachelor en Informatique à Ynov Aix Campus.

---

## 🎯 Notre Objectif

L'objectif était de mettre en pratique les compétences acquises tout au long de l'année en construisant une application web complète de A à Z.

Nous avons choisi de développer un forum de discussion en utilisant **Go** pour le backend, en suivant une architecture de type **MVC (Modèle-Vue-Contrôleur)** et en respectant le **cahier des charges** fourni.

---

## 🛠️ Technologies Utilisées

- **Langage principal** : Go (Golang)  
- **Base de données** : MySQL  
- **Routing** : [gorilla/mux](https://github.com/gorilla/mux)  
- **Sessions** : [gorilla/sessions](https://github.com/gorilla/sessions)  
- **Sécurité** : [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)  
- **Configuration** : [joho/godotenv](https://github.com/joho/godotenv)  
- **Frontend** : HTML5 & CSS3 (sans framework)  

---

## ✅ Prérequis

- Go (version 1.18 ou plus récente)  
- Un serveur de base de données MySQL  

---

## 🚀 Comment Lancer le Projet ?

### 1. Cloner le dépôt

```bash
git clone https://github.com/Shiokaa/forum.git
cd forum
```

### 2. Configurer l'environnement

Créez un fichier `.env` à la racine du projet (src) avec ce contenu :

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=votre_utilisateur_mysql
DB_PSW=votre_mot_de_passe_mysql
DB_NAME=forum_db
COOKIE_SECRET=un_long_secret_aleatoire_pour_la_securite_des_sessions
```

### 3. Préparer la base de données

```sql
-- Connectez-vous à MySQL et exécutez :
CREATE DATABASE IF NOT EXISTS forum_db;
-- Puis :
source database/migrations/migrations.sql;
source database/insertions/insertions.sql;
```

### 4. Installer les dépendances

```bash
go mod tidy
```

### 5. Démarrer le serveur

```bash
go run src/main.go
```

Accédez ensuite à [http://localhost:8080](http://localhost:8080)

---

## 🌐 Routes de l'Application

### Vues (Pages affichées)

| Méthode | Route               | Description                                      |
|---------|--------------------|--------------------------------------------------|
| GET     | /                  | Page d'accueil avec les derniers topics         |
| GET     | /inscription       | Formulaire d'inscription                        |
| GET     | /connexion         | Formulaire de connexion                         |
| GET     | /profil            | Page de profil d'un utilisateur                 |
| GET     | /topic             | Page d'un topic et de ses messages              |
| GET     | /message           | Page d'un message et de ses réponses            |
| GET     | /reponse           | Formulaire de réponse                           |
| GET     | /topic/creer       | Formulaire de création de topic                 |
| GET     | /categories        | Liste des catégories et forums                  |
| GET     | /categorie         | Page d'une catégorie                            |
| GET     | /forum             | Page d'un forum                                 |
| GET     | /recherche         | Résultats d'une recherche                       |
| GET     | /admin             | Dashboard de l’administrateur                   |
| GET     | /error             | Page d'erreur générique                         |

### Actions (Traitements de données)

| Méthode | Route                        | Description                                      |
|---------|-----------------------------|--------------------------------------------------|
| POST    | /inscription/traitement     | Traite les données d'inscription                |
| POST    | /connexion/traitement       | Traite les données de connexion                 |
| GET     | /deconnexion                | Déconnecte l'utilisateur                        |
| POST    | /reponse/traitement         | Enregistre une réponse                          |
| POST    | /topic/creer/traitement     | Crée un nouveau topic                           |
| POST    | /feedback/submit            | Vote (like/dislike) sur un message              |
| POST    | /message/delete             | Supprime un message                             |
| POST    | /reply/delete               | Supprime une réponse                            |
| POST    | /topic/delete               | Supprime un topic                               |
| POST    | /admin/user/delete          | Supprime un compte utilisateur (admin)          |

---

## 👥 L'Équipe

Ce projet a été réalisé en binôme :

- **Amaru TOM**
- **Timothé CHAMPIEUX**

---

# Merci de votre attention ! 🎓
