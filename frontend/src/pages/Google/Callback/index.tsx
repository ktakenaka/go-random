import React, { Fragment, useEffect } from "react";
import { GOOGLE_STATE_KEY, GOOGLE_NONCE_KEY } from "constants/auth";
//import "url-search-params-polyfill";

const CallbackPage = () => {
  useEffect(() => {
    const params = new URLSearchParams(window.location.search);

    const state = sessionStorage.getItem(GOOGLE_STATE_KEY);
    sessionStorage.removeItem(GOOGLE_STATE_KEY);

    const nonce = sessionStorage.getItem(GOOGLE_NONCE_KEY);
    sessionStorage.removeItem(GOOGLE_NONCE_KEY);

    if (params.get("state") !== state) {
      // TODO: error handling
      alert("failed to sign in");
    } else {
      const code = params.get("code");
      // TODO: call BE url to create session
      console.log(code);
      console.log(nonce);
    }
  });

  return <Fragment>Signing In...</Fragment>;
};

export default CallbackPage;
