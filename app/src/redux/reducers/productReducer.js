import * as types from "../actions/actionTypes";
import initialState from "./initialState";

export default function productReducer(state = initialState.products, action) {
  switch (action.type) {
    case types.SET_LOCATION_SUCCESS:
      return initialState.products; // Clear the products list if the user selects another location
    case types.LOAD_PRODUCTS_SUCCESS:
      return action.products;
    default:
      return state;
  }
}
