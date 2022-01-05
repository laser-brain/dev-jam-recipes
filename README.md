# Goal

The requirements of this application were defined in the Dev-Jam discord community which can be found at [https://devjam.vercel.app/](https://devjam.vercel.app/). If you want some challenge for your evolving development skills, come and join us!

As the initial requirements were very basic, I used this opportunity to learn more about the go programming language, mongoDB and the hosting solutions of heroku and netlify that I had not yet used.

The Database is a free mongoDB Shared Cluster hosted on [Atlas](https://cloud.mongodb.com/), the SAAS platform of mongo.
The API Backend is written in go and hosted on Heroku, while the frontend is a Vue3 and Typescript application deployed on netlify.

You can see the application live [here](https://festive-ardinghelli-413138.netlify.app/).

## Requirements

- [x] User can see a list of recipe titles
- [x] User can click a recipe title to display a recipe card containing the recipe title, meal type (breakfast, lunch, supper, or snack), number of people it serves, its difficulty level (beginner, intermediate, advanced), the list of ingredients (including their amounts), and the preparation steps.

## Bonus features (optional)

- [x] User can see a photo showing what the item looks like after it has been prepared.
- [x] User can search for a recipe not in the list of recipe titles by entering the meal name into a search box and clicking a 'Search' button. Any open source recipe API may be used as the source for recipes (see The MealDB below).
- [x] User can see a list of recipes matching the search terms
- [x] User can click the name of the recipe to display its recipe card.
- [x] User can see a warning message if no matching recipe was found.
- [x] User can click a 'Save' button on the cards for recipes located through the API to save a copy to this apps recipe file or database.

# Project Setup

## Prerequisites

- node
- go
- yarn

## Local Development

### .env files

This project assumes a .env file in the root folder of the frontend and server projects

- frontend environment variable:

```
VITE_API_URL=http://localhost:8080/recipes
```

- server environment variables:

```
MONGO_DB_USER=[your mongo-db user]
MONGO_DB_PASSWORD=[db user's password]
MONGO_DB_URL=[url to a mongo-db cluster]
PORT=8080
```

### Start local environment

To start the server:

- `$ cd ./server`
- `$ go run .`

Start the frontend in a separate terminal:

- `$ cd ./frontend`
- `$ yarn`
- `$ yarn dev`
- Open your browser at [http://localhost:3000](http://localhost:3000)
