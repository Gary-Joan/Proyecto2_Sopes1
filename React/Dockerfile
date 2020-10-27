FROM node:12
WORKDIR /usr/src/webapp
COPY package.json ./
RUN npm install
COPY . . 
RUN npm install
EXPOSE 3000
CMD [ "npm", "start" ]