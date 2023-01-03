import * as types from "./actionTypes";
import * as searchApi from "../../api/searchApi";

export function loadLocationsSuccess(locations) {
  return { type: types.LOAD_LOCATIONS_SUCCESS, locations };
}

export function loadLocationsFailure(error) {
  return { type: types.ERROR_ENCOUNTERED, error };
}

export function loadLocations(zipcode, filterLimit) {
  return function (dispatch) {
    return searchApi
      .getLocations(zipcode, filterLimit)
      .then((locations) => {
        dispatch(loadLocationsSuccess(locations.data));
      })
      .catch((error) => {
        dispatch(loadLocationsFailure(error));
      });
  };
}
