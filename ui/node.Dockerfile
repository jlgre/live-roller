FROM node

WORKDIR /usr/src/app/fe

COPY package.json ./

RUN npm i
RUN mkdir -p node_modules/.cache && chmod -R 777 node_modules/.cache

EXPOSE 3000

CMD ["npm", "start"]
