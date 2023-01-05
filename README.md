# grocery-data

The motivation behind this application was that my local supermarket was dealing with supply chain issues and was constantly reconfiguring / remodeling their store. This resulted in frustration in our shopping trips being that you never knew if the items you wanted were in stock and even if they were in stock it was challenging to find them.

I found that Kroger (the parent company for my supermarket) has an [API](https://developer.kroger.com/reference/) that is free to signup and use. Interestingly, that API provides more information than the supermarket's website. Specifically, the supermarket's website only tells you which aisle a product is located on and some basic stock information. In contrast, the API tells you the aisle, which side of the aisle, the bay number, the shelf number, and even the number of "facings" the product has. Upon further visits to the physical store, I was able to recognize these markings on the shelves.

The backend is written in Go and utilizes an `http.Client` to retrieve the grocery data from the Kroger API. The frontend is a single page application that uses React and Redux served via NodeJS. The intention is that both the frontend and backend coexist in the same Docker container. The choice of this stack was primarily that these are technologies I enjoy working in and I wanted to get further experience with them.

## Configuring the API environment variables

The Go API requires a few environment variables to be set. An example .env file is as follows:

```
PORT=5000
KROGER_API_BASE_URL=https://api.kroger.com/v1
KROGER_API_CLIENT_ID=########################################################################
KROGER_API_CLIENT_SECRET=########################################
KROGER_API_CHAIN=FRED
GROCERY_DATA_APP_URL=http://localhost:3000
```

- The `PORT` is configured in the .env file and tells the Go web API which port to listen on. It must match to the value that the React/NodeJS app uses in the `webpack.config._environment_.js` file for the `process.env.API_URL` setting.
- The `KROGER_API_BASE_URL` is either `https://api.kroger.com/v1` for production or `https://api-ce.kroger.com/v1/` for development/certification.
- The `KROGER_API_CLIENT_ID` and `KROGER_API_CLIENT_SECRET` values are obtained by [signing up](https://developer.kroger.com/manage/apps) for the Kroger API.
- Kroger owns several supermarket chains and the `KROGER_API_CHAIN` filters locations to a specific store chain. In my case this is "FRED" for "Fred Meyer".
- The `GROCERY_DATA_APP_URL` is the URL of the frontend application. It's necessary for enabling Cross-Origin Requests from the frontend to the backend.

## Build & deploy locally to a docker container

1. Execute the `build.sh` script.
2. Then run: `docker build -t <imageName> .`
3. Then run: `docker run -di -p 3000:3000 -p 5000:5000 --env-file api/.env --name <containerName> <imageName>`

## Future features

- [ ] Implement OAuth2 and allow the user to save a list of items to check
- [ ] Include instructions for deployment to Google Cloud Run
