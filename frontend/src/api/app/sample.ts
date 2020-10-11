import API from "./base";
import { AxiosResponse } from "axios";
import { TypeSample } from "constants/type";

export const getSamplesAPI = (): Promise<AxiosResponse<any>> => {
  return API.get("/samples");
};

export const getSampleAPI = (id: number): Promise<AxiosResponse<any>> => {
  return API.get(`/samples/${id}`);
};

export const postSampleAPI = (
  sample: TypeSample
): Promise<AxiosResponse<any>> => {
  return API.post("/samples", sample);
};

export const putSampleAPI = (
  id: number,
  sample: any
): Promise<AxiosResponse<any>> => {
  return API.put(`/samples/${id}`, sample);
};
