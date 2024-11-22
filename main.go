package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"html/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Product represents a product structure
type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
}

// CartItem represents an item in the shopping cart
type CartItem struct {
	Product  Product
	Quantity float64
}

var (
	products = []Product{
		{ID: 1, Name: "Produto 1", Description: "Descrição do Produto 1", Price: 10.00},
		{ID: 2, Name: "Produto 2", Description: "Descrição do Produto 2", Price: 20.00},
		{ID: 3, Name: "Produto 3", Description: "Descrição do Produto 3", Price: 30.00},
		{ID: 4, Name: "Produto 4", Description: "Descrição do Produto 4", Price: 40.00},
	}
	cart []CartItem
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", listProductsHandler)
	r.Get("/product/{id}", productDetailsHandler)
	r.Post("/cart/add/{id}", addToCartHandler)
	r.Get("/cart", viewCartHandler)
	r.Get("/checkout", checkoutHandler)
	r.Get("/success", successHandler)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}

// listProductsHandler renders the product list page with header and footer
func listProductsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Rota / chamada para listar produtos")
	renderPage(w, "templates/list.html", products)
}

// productDetailsHandler renders the product details page with header and footer
func productDetailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Rota /product/{id} chamada para detalhes do produto")

	// Captura o ID como string da URL
	idStr := chi.URLParam(r, "id")

	// Converte o ID capturado para inteiro
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Erro ao converter ID:", err)
		http.Error(w, "ID do produto inválido", http.StatusBadRequest)
		return
	}

	// Procura o produto correspondente ao ID
	for _, product := range products {
		log.Println("Verificando produto com ID:", product.ID)
		if product.ID == id {
			log.Println("Produto encontrado:", product.Name)
			// Renderiza a página de detalhes do produto
			renderPage(w, "templates/details.html", product)
			return
		}
	}

	// Caso nenhum produto seja encontrado, retorna 404
	log.Println("Produto não encontrado para o ID:", id)
	http.NotFound(w, r)
}

// addToCartHandler adds a product to the cart
func addToCartHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Rota /cart/add/{id} chamada para adicionar ao carrinho")

	// Captura o ID do produto a partir da URL
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Erro ao converter ID:", err)
		http.Error(w, "ID do produto inválido", http.StatusBadRequest)
		return
	}

	// Procura o produto correspondente
	var selectedProduct *Product
	for _, product := range products {
		if product.ID == id {
			selectedProduct = &product
			break
		}
	}

	if selectedProduct == nil {
		log.Println("Produto não encontrado")
		http.NotFound(w, r)
		return
	}

	// Adiciona o produto ao carrinho ou incrementa a quantidade
	// for i, item := range cart {
	// 	if item.Product.ID == id {
	// 		cart[i].Quantity++
	// 		log.Printf("Quantidade do produto '%s' incrementada para %d\n", item.Product.Name, cart[i].Quantity)
	// 		w.WriteHeader(http.StatusOK)
	// 		w.Write([]byte("Produto adicionado ao carrinho com sucesso"))
	// 		return
	// 	}
	// }

	// Adiciona o produto ao carrinho ou incrementa a quantidade
	for i, item := range cart {
		if item.Product.ID == id {
			cart[i].Quantity++
			log.Printf("Quantidade do produto '%s' incrementada para %f\n", item.Product.Name, cart[i].Quantity)
			w.Header().Set("HX-Trigger", "itemAdicionado") // Dispara um evento para o HTMX
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Produto adicionado ao carrinho com sucesso"))
			return
		}
	}

	// Produto não está no carrinho, adiciona um novo item
	// cart = append(cart, CartItem{Product: *selectedProduct, Quantity: 1})
	// log.Printf("Produto '%s' adicionado ao carrinho\n", selectedProduct.Name)
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte("Produto adicionado ao carrinho com sucesso"))
	// Produto não está no carrinho, adiciona um novo item
	cart = append(cart, CartItem{Product: *selectedProduct, Quantity: 1})
	log.Printf("Produto '%s' adicionado ao carrinho\n", selectedProduct.Name)
	w.Header().Set("HX-Trigger", "itemAdicionado") // Dispara um evento para o HTMX
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Produto adicionado ao carrinho com sucesso"))
}

// viewCartHandler renders the shopping cart with header and footer
func viewCartHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Rota /cart chamada para exibir o carrinho")
	log.Printf("Itens no carrinho: %+v\n", cart)
	renderPage(w, "templates/cart.html", cart)
}

// checkoutHandler renders the checkout page without header and footer
func checkoutHandler(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "templates/checkout.html", nil)
}

// successHandler renders the success page without header and footer
func successHandler(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "templates/success.html", nil)
}

// renderPage renders a page template with header and footer if applicable
// renderPage renderiza um template
func renderPage(w http.ResponseWriter, templatePath string, data interface{}) {
	funcMap := template.FuncMap{
		"mul": func(a, b float64) float64 { return a * b },
	}

	tmpl, err := template.New("").Funcs(funcMap).ParseFiles(
		"templates/layout.html",
		templatePath,
		"templates/components/header.html",
		"templates/components/footer.html",
	)

	if err != nil {
		http.Error(w, "Erro ao carregar templates: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, "Erro ao renderizar página: "+err.Error(), http.StatusInternalServerError)
	}
}