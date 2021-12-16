# CPSC 321 Final Project

## Installation Guide

[![CI](https://github.com/CameronSWilliamson/go-react-blog/actions/workflows/main.yml/badge.svg?branch=master)](https://github.com/CameronSWilliamson/go-react-blog/actions/workflows/main.yml)
### Requirements

- NodeJS
  - npm or yarn
- Go
- MySQL/MariaDB Database

### Setup

Make a `config.json` file within the `server` folder. Within this json file, copy the contents of `server/example-config.json` and replace the necessary values.

Open up two terminal windows (this can also be done using bash jobs). With one window, navigate to the server folder and run `go run .`. With the other terminal window, open up the client folder and run the following commands if you use npm as your node package manager:

```bash
npm install
npm start
```

Run these commands if you use yarn as your package manager:

```bash
yarn
yarn start
```

Please Note that the server and the client must be hosted on the same machine. This is done to make development easier. If they are not to be run on the same machine, each `fetch()` call within the client will fail.

## Available Scripts

Within the client folder there are 3 scripts available:

### `yarn start`

Runs the app in the development mode.\
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.\
You will also see any lint errors in the console.

### `yarn test`

Launches the test runner in the interactive watch mode.\
See the section about [running tests](https://facebook.github.io/create-react-app/docs/running-tests) for more information.

### `yarn build`

Builds the app for production to the `build` folder.\
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.\
Your app is ready to be deployed!

See the section about [deployment](https://facebook.github.io/create-react-app/docs/deployment) for more information.

## Improvements

- [x] Get Create Post Working
- [ ] Get likes to show up on posts
- [ ] Get Comments to show up on posts