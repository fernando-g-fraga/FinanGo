# FinanGo
Bot do telegram para registro de gastos diários

# 🧾 FinanGo - Roteiro de Desenvolvimento (MVP)

## ✅ Setup Inicial
- [x] Criar repositório no Git
- [x] Inicializar projeto Go com `go mod init`
- [x] Configurar estrutura MVC (cmd/, internal/controller, service, model, router)
- [x] Adicionar dependência do Gin: `go get github.com/gin-gonic/gin`
- [X] Criar `.env` e usar `github.com/joho/godotenv` para variáveis sensíveis
- [X] Definir estrutura básica do banco de dados (PostgreSQL)

## 🤖 Integração com Telegram
- [ ] Criar bot no BotFather e obter token
- [ ] Criar estrutura de request/response com Telegram (model)
- [ ] Criar endpoint `/webhook` para receber mensagens
- [ ] Criar comando `/start` com mensagem de boas-vindas
- [ ] Tratar mensagens de texto com possíveis gastos (ex: `"mercado 50.00"`)

## 🗃️ Banco de Dados (PostgreSQL)
- [ ] Configurar conexão com Postgres
- [ ] Criar migração inicial com tabela de gastos:
  - id (UUID)
  - user_id (int64 ou string Telegram ID)
  - categoria (string)
  - valor (float)
  - data (timestamp)
- [ ] Criar função para inserir gasto no banco
- [ ] (Opcional) Criar função para listar últimos gastos

## 🧠 Lógica de Processamento
- [ ] Implementar parser para transformar mensagem em gasto (ex: `"uber 23.50"` → categoria: uber, valor: 23.50)
- [ ] Tratar erro de parsing (mensagem incompreensível)
- [ ] Armazenar gasto associado ao user_id do Telegram

## 🧪 Testes
- [ ] Testar envio manual de mensagens via Telegram
- [ ] Validar persistência no banco
- [ ] Testar mensagens inválidas (sem valor, valor incorreto, etc)

## 🚀 Deploy
- [ ] Configurar webhook do Telegram apontando pro endpoint público
- [ ] Criar Dockerfile para o projeto
- [ ] Subir aplicação na VPS ou serviço de nuvem (render.com, fly.io, railway, etc)
- [ ] Configurar SSL (via Caddy ou nginx + Let's Encrypt)

## 🧩 Melhorias Futuras (para versões posteriores)
- [ ] Listar gastos por comando `/gastos`
- [ ] Categorias automáticas e customizáveis
- [ ] Exportação para CSV
- [ ] Dashboard web (opcional)
- [ ] Autenticação por chat ID
- [ ] Limite diário / alertas de orçamento
