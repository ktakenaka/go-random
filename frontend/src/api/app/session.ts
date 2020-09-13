import API from "./base";
import { AxiosResponse } from "axios";

type Data = {
  code: string;
  nonce: string;
};

export const postSessionAPI = (data: Data): Promise<AxiosResponse<any>> => {
  return API.post("/sessions/google", data);
};
