FROM node:18-alpine AS development

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

FROM node:18-alpine AS production

LABEL "website.name"="nolink"
LABEL "maintainer"="thuongtruong109"
LABEL "org.opencontainers.image.authors"="Tran Nguyen Thuong Truong <thuongtruongofficial@gmail.com>"
LABEL "org.opencontainers.image.version"="1.0"
LABEL "org.opencontainers.image.description"="Ship your link in an easier way"
LABEL "org.opencontainers.image.licenses"="MIT"
LABEL "org.opencontainers.image.source"="https://github.com/thuongtruong109/nolink/api"
LABEL "org.opencontainers.image.documentation"="https://github.com/thuongtruong109/nolink/blob/main/README.md"

WORKDIR /app

RUN chown -R node:node /app

USER node

COPY --chown=node:node --from=development /app/package*.json ./
COPY --chown=node:node --from=development /app/.output ./.output

EXPOSE 3000

CMD ["node", ".output/server/index.mjs"]