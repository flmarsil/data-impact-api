# Test technique

### Sujet: 
Ce test a pour but d’évaluer votre niveau de connaissances, d’adaptation, de bonnes pratiques et de logique en développement logiciel. Le test propose d’implémenter plusieurs fonctionnalités, libre à vous de proposer votre approche.
Vous trouverez ci-dessous quelques liens utiles pour vous orienter.
Vous devez créer une API REST en Golang, permettant d’enregistrer des données dans une base de données de type MongoDB. Vous pouvez vous aider du routeur Gin Gonic. Aidez vous aussi du fichier dataset fournir.

Cette API devra supporter les 4 fonctions CRUD suivantes:

   ## Create 
   
   Requête: `POST /add/users`
	
- La donnée de la méthode POST est un fichier, le format du fichier est fourni dans le dataset au format JSON. 
- La donnée doit être de-sérialisée, puis enregistrée dans une base de données MongoDB de manière concurrente pour chaque utilisateur. 
- Les entrées déjà insérées ne devront plus être traitées à nouveau. 
- Le mot de passe doit être crypté avec bcrypt et seulement le hash de celui-ci  qui est inséré en base.
- En plus de l’insertion en base, vous devrez générer un fichier par utilisateur avec comme nom de fichier l’id de l’utilisateur, ce fichier devra contenir uniquement le champ “data” (disponible dans le dataset).

## Login
Requête: `POST /login`

- L’utilisateur doit pouvoir se connecter sur un son profil afin de pouvoir accéder à toutes les données qui lui sont attribuées. On considère que l’utilisateur à un id et un mot de passe.

## Delete
   Requête: `DELETE /delete/user/:id`

- Dois supprimer un user avec son Id, ainsi que le fichier généré.
 
 ## Read
   Requête: `GET /users/list`
   Requête: `GET /user/:id`

- Dois récupérer un utilisateur avec son id ou une liste d'utilisateurs.

## Update
Requête: `UPDATE /user/:id`

- Dois modifier un utilisateur avec son id, si le champ data change le fichier doit être modifié.


Liens utiles:
https://godoc.org/golang.org/x/crypto/bcrypt 
https://github.com/mongodb/mongo-go-driver 
https://github.com/gin-gonic/gin 