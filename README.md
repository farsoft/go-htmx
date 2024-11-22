# App Exemplo simples de GOLANG + go template + htmx + alpinejs

## Executar

```
go run main.go

ou com make instalado

make 
```

Abrir no navegador http://localhost:8080

## Modo DEV

```
make dev
```

## Instalação do air

```
	go install github.com/cosmtrek/air@latest
 	go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest

    ou    

    go install github.com/cosmtrek/air@latest

    nano ~/.zshrc
    export PATH=$PATH:/home/farnetani/go/pkg/mod/github.com/cosmtrek/air@v1.61.0 

    air init

    OU baixar direto
    wget https://github.com/air-verse/air/releases/download/v1.61.1/air_1.61.1_linux_amd64

    mv air_1.61.1_linux_amd64 air
    chmod +x air

    E ajustar o path no .zshrc para reconhecer 
```

## Colocando o Vite na jogada
```
pnpm init

pnpm install vite --save-dev   
pnpm install @vitejs/plugin-legacy --save-dev
pnpm install alpinejs

```

```
project/
├── main.go
├── templates/
│   ├── layout.html
│   ├── list.html
│   ├── details.html
│   ├── cart.html
│   ├── checkout.html
│   └── success.html
├── public/  # Arquivos estáticos
│   ├── index.html
│   └── assets/
│       └── styles.css
├── src/
│   ├── main.js
│   └── app.js
└── package.json
```
