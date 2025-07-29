# Arquitetura Hexagonal - Estrutura do Projeto

Este projeto foi reorganizado para seguir os princípios da **Arquitetura Hexagonal (Ports and Adapters)** conforme definido por Alistair Cockburn, usando as convenções idiomáticas da linguagem Go.

## Estrutura de Diretórios

### 🚀 `cmd/` - Aplicações Executáveis
Ponto de entrada das aplicações, seguindo a convenção Go.

- **`api/`** - Aplicação principal da API REST
  - **`main.go`** - Ponto de entrada da aplicação

### 🔒 `internal/` - Código Interno da Aplicação
Código que não pode ser importado por outros projetos (proteção do Go).

#### 📦 `domain/` - Centro da Arquitetura (Camada de Domínio)
O **coração** da aplicação, contém a lógica de negócio pura, independente de frameworks ou tecnologias externas.

- **`entities/`** - Entidades de domínio com identidade única
- **`valueobjects/`** - Objetos de valor imutáveis  
- **`interfaces/`** - Interfaces de domínio (contratos)

#### 🎯 `application/` - Camada de Aplicação (Casos de Uso)
Orquestra as regras de negócio e coordena a execução dos casos de uso.

- **`usecases/`** - Implementação dos casos de uso de negócio
- **`ports/input/`** - Interfaces de entrada (o que a aplicação oferece)
- **`ports/output/`** - Interfaces de saída (o que a aplicação precisa)

#### 🔌 `adapters/` - Camada de Adaptadores
Conecta o mundo externo ao domínio da aplicação através de implementações concretas.

##### `input/` - Adaptadores de Entrada
- **`http/`** - Controllers HTTP, rotas e handlers REST
  - **`routes/`** - Definição das rotas da API
  - **Arquivos de controller** - Handlers para cada operação

##### `output/` - Adaptadores de Saída  
- **`persistence/`** - Implementações de repositórios, acesso a dados

#### 🏗️ `infrastructure/` - Camada de Infraestrutura
Detalhes técnicos, configurações e dependências externas.

- **`config/`** - Configurações, tipos de erro, constantes
- **`database/`** - Configuração de banco de dados, migrações
- **`external/`** - Clientes para APIs externas, serviços third-party

#### 🧪 `test/` - Testes
Mantido na estrutura original para testes diversos.

## Estrutura Final (Idiomática Go + Hexagonal)

```
go-api/
├── cmd/                    # Aplicações executáveis
│   └── api/
│       └── main.go        # Ponto de entrada da API
├── internal/              # Código interno da aplicação
│   ├── domain/           # Lógica de negócio pura
│   │   ├── entities/
│   │   ├── valueobjects/
│   │   └── interfaces/
│   ├── application/      # Casos de uso
│   │   ├── usecases/
│   │   └── ports/
│   │       ├── input/
│   │       └── output/
│   ├── adapters/         # Implementações concretas
│   │   ├── input/
│   │   │   └── http/     # Controllers e rotas
│   │   └── output/
│   │       └── persistence/
│   ├── infrastructure/   # Configurações e serviços
│   │   ├── config/
│   │   ├── database/
│   │   └── external/
│   └── test/            # Testes
├── go.mod               # Dependências do projeto
├── go.sum
└── README.md           # Documentação principal
```

## Princípios da Arquitetura Hexagonal

### 1. **Inversão de Dependência**
- O domínio **NÃO** depende de infraestrutura
- Adapters implementam interfaces definidas no domínio
- Fluxo de dependência: Adapters → Application → Domain

### 2. **Ports (Portas)**
- **Input Ports**: Interfaces que definem o que a aplicação oferece
- **Output Ports**: Interfaces que definem o que a aplicação precisa

### 3. **Adapters (Adaptadores)**  
- **Input Adapters**: Implementam como o mundo externo acessa a aplicação
- **Output Adapters**: Implementam como a aplicação acessa recursos externos

## Convenções Go + Hexagonal

### 🔒 **`internal/` Package**
- Código que não pode ser importado por outros projetos
- Proteção nativa do Go para encapsular a arquitetura interna
- Ideal para toda a lógica da arquitetura hexagonal

### 🚀 **`cmd/` Package**
- Padrão Go para aplicações executáveis
- Permite múltiplas aplicações no mesmo repositório
- Separa configuração de inicialização da lógica de negócio

### 📁 **Estrutura de Imports**
```go
// ✅ Correto - usando internal/
import "api/internal/domain/entities"
import "api/internal/application/usecases"

// ❌ Evitar - expondo estrutura interna
import "api/src/controller"
```

## Benefícios da Reorganização

✅ **Idiomática Go**: Segue convenções da linguagem  
✅ **Encapsulamento**: `internal/` protege arquitetura interna  
✅ **Testabilidade**: Lógica de negócio isolada e facilmente testável  
✅ **Flexibilidade**: Mudança de frameworks sem impactar o domínio  
✅ **Manutenibilidade**: Separação clara de responsabilidades  
✅ **Independência**: Domínio independente de detalhes técnicos  
✅ **Escalabilidade**: Estrutura que suporta crescimento da aplicação

## Migração Completa

### Antes (MVC)
```
src/
├── controller/     # Misturava entrada HTTP com lógica
├── model/         # Entidades misturadas com persistência  
├── view/          # Não usado em APIs
└── configuration/ # Configurações espalhadas
main.go            # Na raiz do projeto
```

### Depois (Go Idiomático + Hexagonal)
```
cmd/api/main.go     # Ponto de entrada isolado
internal/           # Código protegido da aplicação
├── domain/         # Lógica de negócio pura
├── application/    # Casos de uso organizados
├── adapters/       # Separação clara de entrada/saída
└── infrastructure/ # Detalhes técnicos isolados
```

## Próximos Passos

1. **Implementar Entities** em `internal/domain/entities/`
2. **Definir Use Cases** em `internal/application/usecases/`  
3. **Criar Input Ports** em `internal/application/ports/input/`
4. **Definir Output Ports** em `internal/application/ports/output/`
5. **Implementar Repositories** em `internal/adapters/output/persistence/`
6. **Refatorar Controllers** para chamar casos de uso via ports

## Comandos de Build

```bash
# Compilar a aplicação
go build -o bin/api ./cmd/api

# Executar a aplicação
./bin/api

# Ou executar diretamente
go run ./cmd/api
```

## Referências

📚 [Hexagonal Architecture Explained - Alistair Cockburn](https://www.amazon.com/Hexagonal-Architecture-Explained-Alistair-Cockburn/dp/173751978X)  
🐹 [Go Project Layout - golang-standards](https://github.com/golang-standards/project-layout) 