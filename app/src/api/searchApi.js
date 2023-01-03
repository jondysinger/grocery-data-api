import { handleResponse, handleError } from "./apiUtils";

export const ZIP_CODE = "97224";
export const MAX_LOCATIONS = 25;
export const PRODUCTS_PER_PAGE = 25;
export const MAX_PRODUCTS = 250;

export function getLocations(zipcode, filterLimit) {
  let query = `zipcode=${zipcode}&filterLimit=${filterLimit}`;
  return fetch(process.env.API_URL + "/locations?" + query, {
    headers: { "content-type": "application/json" },
  })
    .then(handleResponse)
    .catch(handleError);
}

export function getProducts(filterTerm, locationId, filterOffset, filterLimit) {
  let query = `filterTerm=${filterTerm}&locationId=${locationId}&filterOffset=${filterOffset}&filterLimit=${filterLimit}`;
  return fetch(process.env.API_URL + "/products?" + query, {
    headers: { "content-type": "application/json" },
  })
    .then(handleResponse)
    .catch(handleError);
}
