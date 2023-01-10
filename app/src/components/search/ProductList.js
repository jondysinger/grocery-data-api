import React from "react";
import PropTypes from "prop-types";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faAngleLeft,
  faAnglesLeft,
  faAngleRight,
  faAnglesRight,
} from "@fortawesome/free-solid-svg-icons";
import * as searchDefaults from "../../api/searchApi";

function getProductThumbnail(product) {
  if (!product.images) {
    return "";
  }
  let frontImages = product.images.find((i) => {
    return i.perspective === "front";
  });
  if (!frontImages || !frontImages.sizes) {
    return "";
  }
  let thumbnailImage = frontImages.sizes.find((s) => {
    return s.size === "thumbnail";
  });
  return thumbnailImage?.url ?? "";
}

function getProductInventoryLevel(product) {
  if (!product.items || product.items.length == 0) {
    return "";
  }
  let item = product.items[0];
  switch (item.inventory.stockLevel) {
    case "HIGH":
      return "High";
    case "LOW":
      return "Low";
    case "TEMPORARILY_OUT_OF_STOCK":
      return "Out of stock";
    default:
      return "Unknown";
  }
}

const ProductList = ({
  products,
  onNavigateFarLeftCLick,
  onNavigateLeftClick,
  onNavigateRightClick,
  onNavigateFarRightClick,
}) => (
  <>
    <table className="table">
      <thead>
        <tr>
          <th className="text-center">Image</th>
          <th>Description</th>
          <th className="text-center">Stock Level</th>
        </tr>
      </thead>
      <tbody>
        {products.data.map((product) => {
          return (
            <tr key={product.productId}>
              <td className="text-center">
                <img
                  src={getProductThumbnail(product)}
                  alt="Product Thumbnail"
                ></img>
              </td>
              <td>{product.description}</td>
              <td className="text-center">
                {getProductInventoryLevel(product)}
              </td>
            </tr>
          );
        })}
      </tbody>
    </table>
    {products.data.length === 0 ? (
      <></>
    ) : (
      <div id="productNavigation">
        <FontAwesomeIcon
          icon={faAnglesLeft}
          id="navigateFarLeft"
          onClick={onNavigateFarLeftCLick}
          className="fa-lg"
        />
        <FontAwesomeIcon
          icon={faAngleLeft}
          id="navigateLeft"
          onClick={onNavigateLeftClick}
          className="fa-lg"
        />
        <span id="navigateText">
          {products.meta.pagination.start + 1}
          {" to "}
          {products.meta.pagination.start + products.meta.pagination.limit >
          products.meta.pagination.total
            ? products.meta.pagination.total
            : products.meta.pagination.start + products.meta.pagination.limit}
          {" of "}
          {products.meta.pagination.total > searchDefaults.MAX_PRODUCTS
            ? searchDefaults.MAX_PRODUCTS
            : products.meta.pagination.total}
        </span>
        <FontAwesomeIcon
          icon={faAngleRight}
          id="navigateRight"
          onClick={onNavigateRightClick}
          className="fa-lg"
        />
        <FontAwesomeIcon
          icon={faAnglesRight}
          id="navigateFarRight"
          onClick={onNavigateFarRightClick}
          className="fa-lg"
        />
      </div>
    )}
  </>
);

ProductList.propTypes = {
  products: PropTypes.object.isRequired,
  onNavigateFarLeftCLick: PropTypes.func.isRequired,
  onNavigateLeftClick: PropTypes.func.isRequired,
  onNavigateRightClick: PropTypes.func.isRequired,
  onNavigateFarRightClick: PropTypes.func.isRequired,
};

export default ProductList;
