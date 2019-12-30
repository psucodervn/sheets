# develop stage
FROM node:12.14-alpine as develop-stage
WORKDIR /app
COPY package*.json ./
COPY yarn.lock ./
RUN yarn install

# build stage
FROM develop-stage as build-stage
COPY . .
RUN yarn build

# production stage
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
COPY nginx/default.conf /etc/nginx/conf.d/nginx.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
