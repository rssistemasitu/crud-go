# crud-go
Crud em GOLang com acesso a banco de dados MongoDB

## Libs 

### RestAPI
```
go get github.com/gin-gonic/gin
```

### Logs
```
go get go.uber.org/zap
```

### Validações
```
go get github.com/go-playground/validator/v10
```

### Banco de dados mongodb
```
go get go.mongodb.org/mongo-driver/mongo
```

### Token JWT
```
go get github.com/golang-jwt/jwt/v4
```

### Tests
```
go get github.com/stretchr/testify
```

### Mock de interfaces para testes de services
```
go install go.uber.org/mock/mockgen@latest
```
COMANDO PARA GERAR MOCK   
use ```pwd``` para pegar o coaminho completo
```
mockgen -source=(onde_esta_a_interface) -destination=(para_onde_vai_o_mock) -packages=(nome_do_pacote_de_mock)
```

depois para atualizar as libs  
```
go mod tidy
``` 

