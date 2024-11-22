import { defineConfig } from 'vite';

// Configuração básica do Vite
export default defineConfig({
    root: './public', // Define a pasta de entrada como "public"
    server: {
        port: 3000, // Porta do servidor de desenvolvimento
    },
    build: {
        outDir: '../dist', // Define a saída para a pasta "dist"
    },
    resolve: {
        alias: {
            '/src': 'src', // Alias para facilitar importações
            // '/src': '/mnt/wsl/PHYSICALDRIVE4/estudos/golang-web/chatgpt-html-template/src',
        },
    },
});
