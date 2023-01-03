import React from "react";
import { NavLink } from "react-router-dom";

const Header = () => {
  const activeStyle = { color: "#F15B2A" };
  return (
    <>
      <h1>Is My Stuff In Stock?</h1>
      <nav>
        <NavLink to="/search" activeStyle={activeStyle}>
          Search
        </NavLink>
        {" | "}
        <NavLink to="/about" activeStyle={activeStyle}>
          About
        </NavLink>
      </nav>
    </>
  );
};

export default Header;
