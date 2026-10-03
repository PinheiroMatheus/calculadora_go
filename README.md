```markdown
# 🧮 Terminal Calculator em Go

Uma calculadora interativa via terminal desenvolvida em **Go (Golang)**. O programa permite construir expressões matemáticas passo a passo, atualizando a visualização no terminal de forma limpa a cada entrada.

---

## 📌 Funcionalidades Atuais

- 🧹 **Limpeza Automática de Tela:** Atualiza a tela a cada valor ou operador digitado via sequências ANSI (`\033[H\033[2J`).
- 👁️ **Visualização da Expressão:** Exibe a linha do cálculo sendo construída progressivamente no terminal.
- 🛡️ **Validação de Entrada:**
  - Checa se os números inseridos são válidos (suporta números decimais via `float64`).
  - Restringe operadores aos símbolos permitidos: `+`, `-`, `*`, `/` e `=`.

---

## 🚀 Como Executar

### Pré-requisitos
- [Go instalado](https://go.dev/dl/) (versão 1.18 ou superior).