## PRESENTATION TEST TECHNIQUE

Le projet est composé de différents microservices :
- Une base de donnee Mongodb
- Une interface graphique MongoExpress pour une gestion direct depuis le navigateur
- Un serveur backend en Golang
- Un reverse proxy Traefik pour diriger les requêtes vers les différents services
- Un volume partage entre le serveur backend et l'host pour visionner plus facilement la création du fichier data pour chaque utilisateur créé dans la base de donnée. 
- Un container éphémère type "seed" pour initialiser la base de donnée automatiquement

Pour voir les interactions avec la base de donnée : `data_impact/srcs/requirements/server/internal/adapters/framework/right/mongodb`
Pour voir la logique métier des différentes fonctions : `data_impact/srcs/requirements/server/internal/domain/services`
Pour voir le serveur http et les différentes fonctions : `data_impact/srcs/requirements/server/internal/adapters/framework/left/http/`
Pour accéder au volume partage : `data_impact/data/db/mongo/shared`

`Le projet a été code sur Linux`

## PREREQUIS 

Le projet nécessite l'installation de `Docker` et de `docker-compose` pour être lance.


## LANCER / ARRETER LE PROJET

Se placer dans le répertoire courant et lancer la commande ci-dessous :

```bash
# lancer
make install

# arreter
make fclean
```

Ouvrir le navigateur et se rendre à l'adresse : http://app1.traefik.me/ pour accéder à MongoExpress et avoir une visibilité sur la base de donnée via une interface utilisateur.

## TESTING 

#### POST

- Importer dans la base de donnée différents utilisateurs via un fichier `.json` situe ici : `data_impact/data/db/mongo`

Note : Les passwords utilisateurs sont hashes avant d'etre enregistre en base de donnee

``` bash
curl -X POST http://app2.traefik.me/add/users -F file=@data/data.json
```

#### LOGIN

``` bash
curl -X POST http://app2.traefik.me/login \
  -d '{"id": "", "password": ""}'

# exemple d'id depuis le fichier data.json : 1qS9OI4YX8daKvHpwvhrUt6PVnG6MLQMemeFirBdqzEjwibcE1y1EZJELvXWi6w7hU9GwHMQ0RgVc3uWEOEJBbwolVD7rqIUgcwN

# mot de passe correspondant a l'id : xZinMPI3moup6ymz3gVn

```
#### READ

- Récupérer une liste d'utilisateur ou bien un utilisateur en fonction de son id
- Le profil utilisateur doit s'afficher uniquement s'il est connecté

Note : Les champs `_id` et `password` ne sont volontairement pas renvoyés depuis le backend

``` bash
curl http://app2.traefik.me/users/list

curl http://app2.traefik.me/user/:id -H "Token: put_token_here"

# exemple d'id depuis le fichier data.json : 1qS9OI4YX8daKvHpwvhrUt6PVnG6MLQMemeFirBdqzEjwibcE1y1EZJELvXWi6w7hU9GwHMQ0RgVc3uWEOEJBbwolVD7rqIUgcwN
```

#### UPDATE

- Modifier un utilisateur en fonction de son id
- Modifier le fichier généré a sa création en cas de modification du champ data et/ou id

Note : Je suis parti du principe que les modifications se feront toujours avec tous les champs remplis. Je ne traite pas les champs au cas par cas, ce qui signifie que si on envoie un seul champ uniquement, la valeur des autres champs sera mise à nul, et tous les champs dans la base de donnée seront modifié à nul également. (je n'ai pas trouvé comment faire proprement, j'ai donc choisi de procéder ainsi)

``` bash
curl -X PUT http://app2.traefik.me/user/1qS9OI4YX8daKvHpwvhrUt6PVnG6MLQMemeFirBdqzEjwibcE1y1EZJELvXWi6w7hU9GwHMQ0RgVc3uWEOEJBbwolVD7rqIUgcwN \
   -H "Content-Type: application/json" \
   -d '{ 
      "id": "NewId0000", 
      "password": "MyNewSuperPassword!!!!", 
      "isActive": true, 
      "balance": "$221,000.40", 
      "age": 77, 
      "name": "New Farley", 
      "gender": "male", 
      "company": "New Compagny", 
      "email": "new@email.com", 
      "phone": "+1 (111) 688-9999", 
      "address": "588 New Street, NewZiland, MissNewri, 97",
      "about": "This is my new bio about me ! rn", 
      "registered": "2017-10-12T04:14:27 -07:00", 
      "latitude": -10.18881, 
      "longitude": 106.99986, 
      "tags": [ 
        "new", 
        "new",
        "new", 
        "new",
        "new",
        "new",
        "new"
      ],
      "friends": [
        {
          "id": 0,
          "name": "Petty New"
        },
        {
          "id": 1,
          "name": "New Schneider"
        },
        {
          "id": 2,
          "name": "Duffy New"
        }
      ],
      "data":"New Data" 
    }'
```

#### DELETE

- Supprimer un utilisateur en fonction de son id, ainsi que le fichier généré a sa création

``` bash
curl -X DELETE http://app2.traefik.me/delete/user/:id

# tester avec l'id precedement modifie : NewId0000
```
