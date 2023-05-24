# Estágio de construção
FROM node:16-alpine as build

WORKDIR /app

# Copia os arquivos de package.json e package-lock.json
COPY package*.json ./

# Instala as dependências do projeto
RUN npm install

# Copia o restante do código-fonte
COPY . .

# Compila a aplicação Next.js
RUN npm run build

# Estágio de produção
FROM node:16-alpine

WORKDIR /app

# Copia a pasta .next e o package.json
COPY --from=build /app/.next ./.next
COPY --from=build /app/package*.json ./

# Instala somente as dependências de produção
RUN npm install --production

# Define a porta em que a aplicação irá rodar
EXPOSE 3000

# Comando para iniciar a aplicação
CMD ["npm", "start"]