import React, { useEffect, useState } from "react";
import { connect } from "react-redux";
import { setLocation } from "../../redux/actions/selectedLocationActions";
import { loadLocations } from "../../redux/actions/locationActions";
import { loadProducts } from "../../redux/actions/productActions";
import PropTypes from "prop-types";
import LocationInput from "./LocationInput";
import SearchInput from "./SearchInput";
import ProductList from "./ProductList";
import Spinner from "../common/Spinner";
import * as searchDefaults from "../../api/searchApi";

export function SearchPage({
  setLocation,
  loadLocations,
  loadProducts,
  ...props
}) {
  const [loadingLocations, setLoadingLocations] = useState(false);
  const [loadingProducts, setLoadingProducts] = useState(false);
  const [searchTerm, setSearchTerm] = useState("");

  useEffect(() => {
    const getLocations = async () => {
      if (props.locations.length === 0) {
        setLoadingLocations(true);
        await loadLocations(
          searchDefaults.ZIP_CODE,
          searchDefaults.MAX_LOCATIONS
        );
        setLoadingLocations(false);
      }
    };
    getLocations();
  });

  async function getProducts(offset) {
    setLoadingProducts(true);
    await loadProducts(
      searchTerm,
      props.selectedLocation.locationId,
      offset,
      searchDefaults.PRODUCTS_PER_PAGE
    );
    setLoadingProducts(false);
  }

  function handleLocationChange(event) {
    const { value } = event.target;
    let selectedLocation = props.locations.find((loc) => {
      return loc.locationId === value;
    });
    setLocation(selectedLocation);
  }

  function handleSearchTermChange(event) {
    const { value } = event.target;
    setSearchTerm(value);
  }

  function handleSearchButtonClick(event) {
    event.preventDefault();
    getProducts(0);
  }

  function handleNavigateFarLeftClick() {
    if (props.products.meta.pagination.start > 0) {
      getProducts(0);
    }
  }

  function handleNavigateLeftClick() {
    let previousPage =
      props.products.meta.pagination.start - searchDefaults.PRODUCTS_PER_PAGE;
    if (props.products.meta.pagination.start > 0) {
      getProducts(previousPage);
    }
  }

  function handleNavigateRightClick() {
    let totalProducts =
      props.products.meta.pagination.total > searchDefaults.MAX_PRODUCTS
        ? searchDefaults.MAX_PRODUCTS
        : props.products.meta.pagination.total;
    let nextPage =
      props.products.meta.pagination.start + searchDefaults.PRODUCTS_PER_PAGE;
    if (nextPage < totalProducts) {
      getProducts(nextPage);
    }
  }

  function handleNavigateFarRightClick() {
    let lastPage =
      Math.floor(
        props.products.meta.pagination.total / searchDefaults.PRODUCTS_PER_PAGE
      ) * searchDefaults.PRODUCTS_PER_PAGE;
    // The Kroger API advertises that you can offset by 1 to 1000. However, in practice, the API will
    // throw a 400 error with "invalid parameters" as the explanation if you try an offset higher than
    // 250.
    lastPage = Math.min(
      lastPage,
      searchDefaults.MAX_PRODUCTS - searchDefaults.PRODUCTS_PER_PAGE
    );
    if (props.products.meta.pagination.start != lastPage) {
      getProducts(lastPage);
    }
  }

  function getErrorMessage(error) {
    let errorMessage = "";
    try {
      errorMessage = JSON.parse(error.message).message;
    } catch (e) {
      errorMessage = error;
    }
    return errorMessage;
  }

  return loadingLocations ? (
    <Spinner />
  ) : (
    <>
      <LocationInput
        name="location"
        label="Location:"
        locations={props.locations}
        value={props.selectedLocation?.locationId}
        onChange={handleLocationChange}
        placeholder="Select a location"
      />
      <SearchInput
        name="searchTerm"
        label="Search Term:"
        buttonText="Search"
        value={props.searchTerm}
        disabled={!props.selectedLocation}
        onChange={handleSearchTermChange}
        onClick={handleSearchButtonClick}
      />
      {props.error && (
        <div className="alert alert-danger">
          {"Error: " + getErrorMessage(props.error)}
        </div>
      )}
      <hr />
      {loadingProducts ? (
        <Spinner />
      ) : (
        <ProductList
          products={props.products}
          onNavigateFarLeftCLick={handleNavigateFarLeftClick}
          onNavigateLeftClick={handleNavigateLeftClick}
          onNavigateRightClick={handleNavigateRightClick}
          onNavigateFarRightClick={handleNavigateFarRightClick}
        />
      )}
    </>
  );
}

SearchPage.propTypes = {
  selectedLocation: PropTypes.object,
  locations: PropTypes.array.isRequired,
  searchTerm: PropTypes.string,
  products: PropTypes.object.isRequired,
  error: PropTypes.object,
  setLocation: PropTypes.func.isRequired,
  loadLocations: PropTypes.func.isRequired,
  loadProducts: PropTypes.func.isRequired,
};

function mapStateToProps(state) {
  return {
    selectedLocation: state.selectedLocation,
    locations: state.locations,
    searchTerm: state.searchTerm,
    products: state.products,
    error: state.error,
  };
}

const mapDispatchToProps = {
  setLocation,
  loadLocations,
  loadProducts,
};

export default connect(mapStateToProps, mapDispatchToProps)(SearchPage);
