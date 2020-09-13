import React from "react";
import { connect } from "react-redux";

import { ActionMessageWrapper } from "./style";

type Props = {
  message?: string | null;
  success?: boolean;
};

const ActionMessage = ({ message = null, success = true }: Props) => {
  if (message) {
    return (
      <ActionMessageWrapper
        className={`color-${success ? "success" : "failure"}`}
      >
        {message}
      </ActionMessageWrapper>
    );
  } else {
    return null;
  }
};

const mapStateToProps = (state: Readonly<any>) => ({
  message: state.getIn(["app", "message"]),
  success: state.getIn(["app", "success"]),
});

export default connect(mapStateToProps, null)(ActionMessage);
