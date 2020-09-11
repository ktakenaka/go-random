import React, { Fragment, useEffect } from "react";
import { GOOGLE_STATE_KEY, GOOGLE_NONCE_KEY } from "constants/auth";
import "url-search-params-polyfill";
import srs from "secure-random-string";

const scope = encodeURIComponent("openid email");
const responseType = "code";
const redirectUri = `${process.env.REACT_APP_GOOGLE_REDIRECT_URI}`;
const clientId = `${process.env.REACT_APP_GOOGLE_CLIENT_ID}`;
const oauthUri = "https://accounts.google.com/o/oauth2/v2/auth";

const SignInPage = () => {
  useEffect(() => {
    const state = srs();
    sessionStorage.setItem(GOOGLE_STATE_KEY, state);

    const nonce = srs();
    sessionStorage.setItem(GOOGLE_NONCE_KEY, nonce);

    const queryParams = new URLSearchParams({
      state: state,
      nonce: nonce,
      scope: scope,
      client_id: clientId,
      redirect_uri: encodeURIComponent(redirectUri),
      response_type: responseType,
    });
    window.location.href = [
      oauthUri,
      decodeURIComponent(queryParams.toString()),
    ].join("?");
  });

  return <Fragment>Redirecting Google...</Fragment>;
};

export default SignInPage;
