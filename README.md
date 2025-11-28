# MovieMatch 🎬 (Go)

MovieMatch est une **API REST de recommandation de films** écrite en Go
## 🚀 Stack technique

- Go 1.22
- Framework HTTP : [Gin](https://github.com/gin-gonic/gin)
- ORM : [GORM](https://gorm.io/)
- Base de données : PostgreSQL
- Authentification : JWT (github.com/golang-jwt/jwt)
- Hash mot de passe : bcrypt
- Conteneurisation : Docker & docker-compose
- Tests unitaires : `go test`

## 📁 Architecture

```bash
moviematch/
├── cmd/api/main.go        # point d'entrée de l'API
├── internal/
│   ├── config/            # config + connexion DB
│   ├── domain/            # entités métier (Movie, User)
│   ├── repository/        # accès base de données
│   ├── service/           # logique métier (auth, recommandation)
│   └── handler/           # handlers HTTP (Gin)
├── pkg/jwt/               # wrapper JWT
├── tests/                 # tests unitaires
├── Dockerfile
├── docker-compose.yml
└── README.md
```

## 🔧 Lancer en local (sans Docker)

```bash
go mod tidy
go run ./cmd/api
```

Par défaut, l'API écoute sur `http://localhost:8080`.

## 🐳 Lancer avec Docker / docker-compose

```bash
docker-compose up --build
```

- API : http://localhost:8080/api
- DB Postgres : port 5432

## 🔐 Endpoints principaux

- `POST /api/register` – inscription `{ "email": "...", "password": "..." }`
- `POST /api/login` – connexion, retourne un token JWT
- `GET /api/movies` – liste tous les films (protégé, header `Authorization: Bearer <token>`)
- `GET /api/movies/recommend` – renvoie les meilleurs films par rating

## ✅ Tests

```bash
go test ./...
```

## 🧠 Idées d'amélioration

- Algorithme de recommandation plus avancé (similarité de genres, historique utilisateur…)
- Ajout d'un vrai endpoint `POST /api/movies` pour créer des films
- Utiliser un vrai système de migration (golang-migrate, etc.)
- Ajouter CI/CD (GitHub Actions) pour lancer `go test` + build automatique
