// Estruturas de Dados e Algoritmos em Go
// Autor: Exemplo completo com arrays, slices, listas ligadas, pilhas, filas, hashmaps, árvores binárias, grafos e tries.
// Objetivo: Servir como guia de aprendizado com explicações detalhadas e comentários sobre uso e complexidade

package main

import (
	"container/list"
	"fmt"
	"sort"
)

func main() {
	// === ARRAYS ===
	fmt.Println("\n--- Arrays ---")
	// Arrays têm tamanho fixo, definidos em tempo de compilação
	var arr [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array:", arr)
	fmt.Println("Acesso direto arr[2]:", arr[2])
	// Big-O: Acesso por índice = O(1)

	// === SLICES ===
	fmt.Println("\n--- Slices ---")
	// Slices são estruturas dinâmicas sobre arrays; muito mais usados que arrays diretamente
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5)
	fmt.Println("Slice:", slice)
	fmt.Println("Slice ordenado:")
	sort.Ints(slice)
	fmt.Println(slice)

	// Remoção (índice 1) em slice (requer reconstrução do slice)
	slice = append(slice[:1], slice[2:]...)
	fmt.Println("Slice após remoção índice 1:", slice)
	// Big-O: Append (média) = O(1), Remoção = O(n), Ordenação = O(n log n)

	// === Linked List ===
	fmt.Println("\n--- Linked List ---")
	// Lista duplamente encadeada da stdlib (container/list)
	l := list.New()
	l.PushBack("A")
	l.PushBack("B")
	l.PushFront("Start")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	// Big-O: Inserção/remoção O(1) com ponteiro direto; busca = O(n)

	// === Queue (Fila) ===
	fmt.Println("\n--- Queue ---")
	// Implementação com slice: FIFO
	queue := []string{"primeiro"}
	queue = append(queue, "segundo")
	head := queue[0]
	queue = queue[1:] // remove da frente
	fmt.Println("Dequeued:", head)
	fmt.Println("Fila restante:", queue)
	// Big-O: Inserção final = O(1), Remoção início = O(n) devido shift (use list para O(1))

	// === Stack (Pilha) ===
	fmt.Println("\n--- Stack ---")
	// Implementação com slice: LIFO
	stack := []int{}
	stack = append(stack, 1, 2, 3) // push
	top := stack[len(stack)-1]     // peek
	stack = stack[:len(stack)-1]   // pop
	fmt.Println("Topo da pilha:", top)
	fmt.Println("Pilha restante:", stack)
	// Big-O: Push e Pop = O(1)

	// === HashMap (map em Go) ===
	fmt.Println("\n--- HashMap ---")
	// Estrutura chave-valor; tabela hash
	hash := make(map[string]int)
	hash["go"] = 1
	hash["java"] = 2
	fmt.Println("Valor da chave 'go':", hash["go"])
	delete(hash, "java")
	fmt.Println("HashMap após delete:", hash)
	// Big-O: Média Acesso/Inserção/Remoção = O(1), Pior caso raro = O(n)

	// === Árvores Binárias ===
	fmt.Println("\n--- Binary Tree ---")
	// Árvore binária de busca (BST), sem balanceamento
	root := &Node{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(12)
	fmt.Println("Valores em ordem:")
	root.PrintInOrder()
	// Big-O: O(log n) média, O(n) pior caso (lista encadeada degenerada)

	// === Grafos ===
	fmt.Println("\n--- Grafo ---")
	// Representação por listas de adjacência
	graph := make(map[string][]string)
	graph["A"] = []string{"B", "C"}
	graph["B"] = []string{"D"}
	graph["C"] = []string{"E"}
	fmt.Println("Vizinhos de A:", graph["A"])
	// Big-O: Depende da operação. Verificar conexão direta = O(1) se hash/lista curta

	// === Trie ===
	fmt.Println("\n--- Trie ---")
	// Árvore de prefixos, eficiente para buscas por palavras
	trie := NewTrie()
	trie.Insert("go")
	trie.Insert("golang")
	fmt.Println("Contém 'go'?", trie.Search("go"))
	fmt.Println("Contém 'java'?", trie.Search("java"))
	// Big-O: Inserção/Busca = O(m), onde m = comprimento da palavra

	// Extras possíveis: AVL, Red-Black Tree, B-Tree, Heap, Union-Find, Dijkstra, etc.
}

// Estrutura para árvore binária
// Cada nó contém valor, e ponteiros para a subárvore à esquerda e à direita

// --- Binary Tree ---
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(val int) {
	if val < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: val}
		} else {
			n.Left.Insert(val)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: val}
		} else {
			n.Right.Insert(val)
		}
	}
}

func (n *Node) PrintInOrder() {
	if n == nil {
		return
	}
	n.Left.PrintInOrder()
	fmt.Println(n.Value)
	n.Right.PrintInOrder()
}

// --- Trie (Prefix Tree) ---
type TrieNode struct {
	Children map[rune]*TrieNode
	End      bool
}

type Trie struct {
	Root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{Root: &TrieNode{Children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
	node := t.Root
	for _, ch := range word {
		if _, exists := node.Children[ch]; !exists {
			node.Children[ch] = &TrieNode{Children: make(map[rune]*TrieNode)}
		}
		node = node.Children[ch]
	}
	node.End = true
}

func (t *Trie) Search(word string) bool {
	node := t.Root
	for _, ch := range word {
		if _, exists := node.Children[ch]; !exists {
			return false
		}
		node = node.Children[ch]
	}
	return node.End
}
