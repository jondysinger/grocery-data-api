import * as types from "../actions/actionTypes";
import initialState from "./initialState";

export default function errorReducer(state = initialState.error, action) {
  switch (action.type) {
    case types.LOAD_LOCATIONS_SUCCESS:
    case types.LOAD_PRODUCTS_SUCCESS:
      return initialState.error; // Clear error on successful API call
    case types.ERROR_ENCOUNTERED:
      return action.error;
    default:
      return state;
  }
}
