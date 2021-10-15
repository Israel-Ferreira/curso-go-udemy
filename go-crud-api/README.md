# Go Crud API

Para rodar o container, você deve passar as variáveis de ambiente:
* `GO_CRUD_DB_PORT`
* `GO_CRUD_DB_USERNAME`
* `GO_CRUD_DB_PASSWORD`
* `GO_CRUD_DB_HOST`
* `GO_CRUD_DB`

### Exemplo

```shell
docker run --name crud-user -p 5000:5000 \
-e GO_CRUD_DB_PORT=3306 \
-e GO_CRUD_DB_USERNAME=root \
-e GO_CRUD_DB_PASSWORD=root123 \
-e GO_CRUD_DB=devbook \
-e GO_CRUD_DB_HOST=mysql-curso-go-db \
-d israelsouza17/go-crud-api
```


