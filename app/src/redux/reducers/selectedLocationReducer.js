import * as types from "../actions/actionTypes";
import initialState from "./initialState";

export default function selectedLocationReducer(
  state = initialState.selectedLocation,
  action
) {
  switch (action.type) {
    case types.SET_LOCATION_SUCCESS:
      return action.location;
    default:
      return state;
  }
}
