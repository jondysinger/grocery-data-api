import * as types from "./actionTypes";

export function setLocationSuccess(location) {
  return { type: types.SET_LOCATION_SUCCESS, location };
}

export function setLocation(location) {
  return function (dispatch) {
    return dispatch(setLocationSuccess(location));
  };
}
