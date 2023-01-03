import * as types from "./actionTypes";
import * as searchApi from "../../api/searchApi";

export function loadProductsSuccess(products) {
  return { type: types.LOAD_PRODUCTS_SUCCESS, products };
}

export function loadProductsFailure(error) {
  return { type: types.ERROR_ENCOUNTERED, error };
}

export function loadProducts(
  filterTerm,
  locationId,
  filterOffset,
  filterLimit
) {
  return function (dispatch) {
    return searchApi
      .getProducts(filterTerm, locationId, filterOffset, filterLimit)
      .then((products) => {
        dispatch(loadProductsSuccess(products));
      })
      .catch((error) => {
        dispatch(loadProductsFailure(error));
      });
  };
}
