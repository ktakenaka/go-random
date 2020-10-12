import API from "./base";
import { AxiosResponse } from "axios";
import { TypeSample } from "constants/type";

export const list = (): Promise<AxiosResponse<any>> => {
  return API.get("/samples");
};

export const get = (id: number): Promise<AxiosResponse<any>> => {
  return API.get(`/samples/${id}`);
};

export const post = (
  sample: TypeSample
): Promise<AxiosResponse<any>> => {
  return API.post("/samples", sample);
};

export const put = (
  id: number,
  sample: any
): Promise<AxiosResponse<any>> => {
  return API.put(`/samples/${id}`, sample);
};

export const destroy = (id: number) => {
  return API.delete(`/samples/${id}`);
};
