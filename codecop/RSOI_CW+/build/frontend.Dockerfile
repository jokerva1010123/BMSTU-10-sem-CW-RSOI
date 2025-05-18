FROM node:latest

COPY . .

WORKDIR frontend/

EXPOSE 8887

CMD npm i  && npm start