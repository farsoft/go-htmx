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

## Instalar o Vite

```
pnpm i -g vite
```

## Instalar o bun

```
- Arch Linux:
  yay -S bun-bin

  bun --version  

- Ubuntu:
curl -fsSL https://bun.sh/install | bash

export PATH="$HOME/.bun/bin:$PATH"

source ~/.bashrc
# ou
source ~/.zshrc

bun --version
```


## Estrutura de diretório
```
├── bin
├── dist
│   ├── assets
│   │   ├── index-<build>.css
│   │   └── index-<build>.js
│   └── index.html
├── go.mod
├── go.sum
├── main.go
├── Makefile
├── README.md
├── templates
│   ├── cart.html
│   ├── checkout.html
│   ├── components
│   │   ├── footer.html
│   │   └── header.html
│   ├── details.html
│   ├── layout.html
│   ├── list.html
│   └── success.html
├── tmp
│   └── main
└── web
    ├── node_modules
    ├── package.json
    ├── pnpm-lock.yaml
    ├── public
    │   ├── index.html
    │   └── static
    │       ├── assets
    │       │   └── styles.css
    │       └── js
    │           └── alpine.min.js
    ├── src
    │   └── main.js
    ├── tailwind.config.js
    └── vite.config.js
```

