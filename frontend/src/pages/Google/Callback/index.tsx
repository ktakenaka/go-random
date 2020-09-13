import React, { Fragment, useEffect } from "react";
import { connect } from "react-redux";

import { GOOGLE_STATE_KEY, GOOGLE_NONCE_KEY } from "constants/auth";
import { createSessionRequest } from "store/actionCreators/session";
import { changeLocation, setMessage } from "store/actionCreators/app";

type Props = {
  loading: boolean;
  createSessionRequest: (code: string, nonce: string) => void;
  changeLocation: (location: string) => void;
  setMessage: (message: string, color?: boolean) => void;
};

const CallbackPage = ({
  loading,
  createSessionRequest,
  changeLocation,
  setMessage,
}: Props) => {
  useEffect(() => {
    function handleFailure(message: string) {
      setMessage(message, false);
      changeLocation("/");
    }
    function handleCallbackPhase() {
      const params = new URLSearchParams(window.location.search);

      const state = sessionStorage.getItem(GOOGLE_STATE_KEY);
      sessionStorage.removeItem(GOOGLE_STATE_KEY);

      const nonce = sessionStorage.getItem(GOOGLE_NONCE_KEY);
      sessionStorage.removeItem(GOOGLE_NONCE_KEY);

      const code = params.get("code");

      if (params.get("state") !== state) {
        handleFailure("csrf detected");
      } else if (!code || !nonce) {
        handleFailure("invalid callback phase");
      } else {
        createSessionRequest(code, nonce);
      }
    }
    handleCallbackPhase();
  }, [createSessionRequest, setMessage, changeLocation]);

  if (loading) {
    return <Fragment>Signing In...</Fragment>;
  } else {
    return <Fragment>Preparing...</Fragment>;
  }
};

const mapStateToProps = (state: Readonly<any>) => ({
  loading: state.get("session").loading,
});

const mapDispatchToProps = {
  createSessionRequest: createSessionRequest,
  changeLocation: changeLocation,
  setMessage: setMessage,
};

export default connect(mapStateToProps, mapDispatchToProps)(CallbackPage);
