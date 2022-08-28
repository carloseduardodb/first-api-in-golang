 ## Api em Go para gerenciamento de usuários
 ### Use o ***insomnia*** para clonar o documento da aplicação e poder testar as rotas

<br />
 <p>
  Rotas disponíveis:
 </p>

```ts
  // create
  curl --request POST \
  --url http://localhost:3000/users \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "Carlos Eduardo",
	"email": "carlos@hotmail.com"
  }'

  // find users
  curl --request GET \
  --url http://localhost:3000/users

  // find user
  curl --request GET \
  --url http://localhost:3000/user/1

  // update user
  curl --request PUT \
  --url http://localhost:3000/user/4 \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "Carlos Eduardo 1 Update",
	"email": "carlos@hotmail.com"
  }'

  // delete user
  curl --request DELETE \
  --url http://localhost:3000/user/4

```