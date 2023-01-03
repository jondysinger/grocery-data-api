import React from "react";

const AboutPage = () => (
  <div>
    <h2>About</h2>
    <p>
      This app leverages the{" "}
      <a href="https://developer.kroger.com/reference/">Kroger API</a> to
      retrieve information about stock levels of products. The backend is an API
      built in Go and the front end utilizes React & Redux served by NodeJS.
    </p>
  </div>
);

export default AboutPage;
