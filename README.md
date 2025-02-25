🏗 Projeto 4 Go

🚀 Descrição
Este projeto é uma aplicação desenvolvida em Go (Golang) seguindo uma arquitetura modular e escalável. Ele implementa funcionalidades robustas para serviços API, validação de dados e persistência em banco de dados.
📂 Estrutura do Projeto

    cmd/ → Contém os comandos principais da aplicação.
        api/ → Configuração da API.
        terndotenv/ → Gerenciamento de variáveis de ambiente.

    internal/ → Código interno organizado por camadas.
        api/ → Configuração dos endpoints.
        jsonutils/ → Utilitários para manipulação de JSON.
        services/ → Implementação das regras de negócio.
        store/pgstore/ → Persistência de dados usando PostgreSQL.
        usecase/ → Casos de uso específicos da aplicação.
        validator/ → Validação de dados de entrada.

    Arquivos de Configuração
        .env → Configurações de ambiente.
        docker-compose.yml → Configuração para execução via Docker.
        .gitignore → Definição de arquivos ignorados no versionamento.

🔧 Como Executar

1️⃣ Clone o repositório:

git clone https://github.com/Jose6348/Projeto_4_Go.git
cd Projeto_4_Go

2️⃣ Instale as dependências:

go mod tidy

3️⃣ Execute a aplicação:

go run main.go

🛠 Tecnologias Utilizadas

    Go (Golang) 🔹
    PostgreSQL 🛢
    Docker 🐳
    Clean Architecture 📐

📌 Este projeto segue boas práticas de modularização e separação de responsabilidades.
📜 Licença

Este projeto está licenciado sob a MIT License.
