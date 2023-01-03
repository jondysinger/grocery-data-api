import React from "react";
import { Route, Switch } from "react-router-dom";
import AboutPage from "./about/AboutPage";
import Header from "./common/Header";
import PageNotFound from "./PageNotFound";
import SearchPage from "./search/SearchPage"; // eslint-disable-line import/no-named-as-default

function App() {
  return (
    <div className="container-fluid">
      <Header />
      <Switch>
        <Route exact path="/" component={SearchPage} />
        <Route path="/about" component={AboutPage} />
        <Route path="/search" component={SearchPage} />
        <Route component={PageNotFound} />
      </Switch>
    </div>
  );
}

export default App;
