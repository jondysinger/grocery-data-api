import React from "react";
import PropTypes from "prop-types";

const LocationInput = ({
  name,
  label,
  locations,
  value,
  onChange,
  placeholder,
}) => {
  return (
    <div className="input-group mb-2">
      <div className="input-group-prepend">
        <span id="locationLabel" className="input-group-text rounded-0">
          {label}
        </span>
      </div>
      <select
        id="locationSelect"
        name={name}
        onChange={onChange}
        value={value}
        className="form-select rounded-0"
      >
        <option value={placeholder}>{placeholder}</option>
        {locations.map((location) => (
          <option key={location.locationId} value={location.locationId}>
            {location.Name}
          </option>
        ))}
      </select>
    </div>
  );
};

LocationInput.propTypes = {
  name: PropTypes.string.isRequired,
  label: PropTypes.string.isRequired,
  locations: PropTypes.array.isRequired,
  value: PropTypes.string,
  onChange: PropTypes.func.isRequired,
  placeholder: PropTypes.string.isRequired,
};

export default LocationInput;
