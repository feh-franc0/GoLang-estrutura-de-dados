# 📚 Estruturas de Dados e Algoritmos em Go (Golang)

Este repositório contém um guia completo e comentado com exemplos práticos de **estruturas de dados fundamentais** em Go. Ideal para quem quer dominar algoritmos, preparar-se para entrevistas técnicas ou reforçar lógica de programação com código limpo e direto ao ponto.

---

## ✅ O que você vai encontrar

| Estrutura        | Tipo             | Complexidade Big-O     | Destaques                              |
| ---------------- | ---------------- | ---------------------- | -------------------------------------- |
| Arrays           | Estática         | Acesso: O(1)           | Base de toda estrutura indexada        |
| Slices           | Dinâmica         | Append: O(1)\*         | Muito usado no dia a dia               |
| Linked List      | Encadeada        | Inserção: O(1)         | Uso com `container/list`               |
| Queue (Fila)     | FIFO             | Inserção/Remoção: O(1) | Implementada com slice ou list         |
| Stack (Pilha)    | LIFO             | Push/Pop: O(1)         | Base de algoritmos recursivos          |
| HashMap (map)    | Tabela Hash      | Acesso: O(1)           | Rápido e direto, ideal para busca      |
| Árvores Binárias | BST              | Busca: O(log n)\*      | Implementação manual                   |
| Grafos           | Lista Adjacência | Varia                  | Representado com `map[string][]string` |
| Trie             | Prefix Tree      | Inserção/Busca: O(m)   | Ideal para autocomplete e buscas       |

---

## 🧠 Aprendizado com Comentários

Cada trecho do código contém comentários explicando:

* Como funciona a estrutura
* Quando usar
* Complexidade computacional
* Cuidados de performance

---

## 🚀 Como executar

1. Clone o repositório:

   ```bash
   git clone https://github.com/seu-usuario/estruturas-go.git
   cd estruturas-go
   ```

2. Execute o projeto:

   ```bash
   go run main.go
   ```

Você verá no terminal a saída explicando cada estrutura com exemplos.

---

## 📂 Estrutura do Código

```go
// Arrays - base fixa e acesso rápido
var arr [5]int = [5]int{10, 20, 30, 40, 50}

// Slices - coleção dinâmica
slice := []int{1, 2, 3}
slice = append(slice, 4, 5)

// Linked List - container/list
list.New().PushBack("A")

// Queue - FIFO com slice
queue := []string{"A", "B"}
queue = queue[1:] // remove do início

// Stack - LIFO com slice
stack := append(stack, 1, 2)
stack = stack[:len(stack)-1]

// HashMap
hash := make(map[string]int)
hash["go"] = 1

// Árvores Binárias
root := &Node{Value: 10}
root.Insert(5)

// Grafos
graph := map[string][]string{"A": {"B", "C"}}

// Trie
trie.Insert("go")
trie.Search("golang")
```

---

## 🎓 Recomendações de Estudo

### 🔸 Fundamentos

* Estude a teoria de cada estrutura: quando usar arrays, listas, árvores, grafos, etc.
* Utilize o site [Big-O Cheat Sheet](https://www.bigocheatsheet.com/) para memorizar complexidades.

### 🔸 Prática Guiada

* Refaça os exemplos desse projeto digitando manualmente.
* Crie suas próprias versões: Stack com lista, Queue com canais (`chan`), Trie com busca parcial.
* Transforme cada estrutura em um pacote separado para treinar modularização.

### 🔸 Algoritmos Clássicos

* Implemente **DFS**, **BFS**, **Merge Sort**, **Quick Sort**, **Dijkstra**, etc.
* Simule estruturas como **heap**, **union-find**, **AVL** e **B-trees**.

### 🔸 Sites para treino

* [LeetCode](https://leetcode.com/) — ótimo para desafios de entrevistas
* [HackerRank](https://www.hackerrank.com/domains/tutorials/10-days-of-javascript) — treinos com foco educacional
* [Beecrowd](https://www.beecrowd.com.br/) — problemas em português com ranking

### 🔸 Dica final

> O melhor jeito de aprender é **construir**. Pegue um problema real e resolva usando essas estruturas: sistemas de agendamento, filas de atendimento, mecanismos de busca, gerenciadores de tarefas.

---

## 🛠 Tecnologias

* Linguagem: **Go (Golang)**
* Versão sugerida: `go1.21+`

---

## 📬 Contato

Desenvolvido por [Fernando Franco](https://github.com/feh-franc0) • Contribuições e melhorias são bem-vindas!

---

## 📄 Licença

Este projeto está sob a licença MIT.
