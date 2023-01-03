import React from "react";
import PropTypes from "prop-types";

const SearchInput = ({
  name,
  label,
  buttonText,
  onChange,
  onClick,
  disabled,
  placeholder,
  value,
}) => {
  return (
    <form onSubmit={onClick}>
      <div className="form-group">
        <div className="input-group mb-2">
          <div className="input-group-prepend">
            <span id="searchTermLabel" className="input-group-text rounded-0">
              {label}
            </span>
          </div>
          <input
            id="searchTermInput"
            type="text"
            name={name}
            placeholder={placeholder}
            disabled={disabled ? "disabled" : ""}
            value={value}
            onChange={onChange}
            aria-describedby="searchTermLabel"
            className="form-control rounded-0"
          />
          <div className="input-group-append">
            <button
              id="submitButton"
              type="submit"
              disabled={disabled ? "disabled" : ""}
              className="btn btn-outline-secondary rounded-0"
            >
              {buttonText}
            </button>
          </div>
        </div>
      </div>
    </form>
  );
};

SearchInput.propTypes = {
  name: PropTypes.string.isRequired,
  label: PropTypes.string.isRequired,
  buttonText: PropTypes.string.isRequired,
  onChange: PropTypes.func.isRequired,
  onClick: PropTypes.func.isRequired,
  disabled: PropTypes.bool.isRequired,
  placeholder: PropTypes.string,
  value: PropTypes.string,
};

export default SearchInput;
