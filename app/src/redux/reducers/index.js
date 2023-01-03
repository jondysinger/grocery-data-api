import { combineReducers } from "redux";
import selectedLocation from "./selectedLocationReducer";
import locations from "./locationReducer";
import products from "./productReducer";
import error from "./errorReducer";

const rootReducer = combineReducers({
  selectedLocation,
  locations,
  products,
  error,
});

export default rootReducer;
