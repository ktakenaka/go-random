import axios from "axios";
import { CSRF_TOKEN_KEY } from "constants/auth";
import { UNAUTHORIZED } from "constants/metaCode";
import history from "browserHistory";

const csrfIgnoreMethods = ["get", "head", "options"];

export const BE_BASE_URL = process.env.REACT_APP_API_BASE_URL;

const isAuthError = (response: any): boolean => {
  return response.status === 401 && response.data.meta.code === UNAUTHORIZED;
};

const initialize = function (): void {
  axios.defaults.withCredentials = true;
  axios.defaults.baseURL = `${BE_BASE_URL}/api/v1`;

  axios.interceptors.request.use(
    function (config) {
      if (config.method && !csrfIgnoreMethods.includes(config.method)) {
        config.headers["X-CSRF-Token"] = localStorage.getItem(CSRF_TOKEN_KEY);
      }
      return config;
    },
    function (error) {
      return Promise.reject(error);
    }
  );

  axios.interceptors.response.use(
    function (response) {
      return response;
    },
    function (error) {
      if (isAuthError(error.response)) {
        // TOOD: try to refresh JWT before sign in again
        history.push("/google/sign-in");
      }
      return Promise.reject(error);
    }
  );
};

initialize();

export default axios;
