import React, { Fragment, useEffect } from "react";
import { connect } from "react-redux";

import { GOOGLE_STATE_KEY, GOOGLE_NONCE_KEY } from "constants/auth";
import { createSessionRequest } from "store/actionCreators/session";

type Props = {
  createSessionRequest: (code: string, nonce: string) => void;
};

const CallbackPage = ({ createSessionRequest }: Props) => {
  useEffect(() => {
    const params = new URLSearchParams(window.location.search);

    const state = sessionStorage.getItem(GOOGLE_STATE_KEY);
    sessionStorage.removeItem(GOOGLE_STATE_KEY);

    const nonce = sessionStorage.getItem(GOOGLE_NONCE_KEY);
    sessionStorage.removeItem(GOOGLE_NONCE_KEY);

    const code = params.get("code");

    if (params.get("state") !== state) {
      // TODO: error handling
      console.log("csrf detected");
    } else if (!code || !nonce) {
      // TOOD: error handling
      console.log("invalid callback phase");
    } else {
      createSessionRequest(code, nonce);
    }
  });

  return (
    <Fragment>
      <div>Signing In...</div>
    </Fragment>
  );
};

// When loading is needed, please use here
// const mapStateToProps = (state: Readonly<any>) => ({
//   loading: state.get("session").loading,
// });

const mapDispatchToProps = {
  createSessionRequest: createSessionRequest,
};

export default connect(null, mapDispatchToProps)(CallbackPage);
