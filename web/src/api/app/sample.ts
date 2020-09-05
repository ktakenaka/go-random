import API from "./base";
import { AxiosResponse } from "axios";

export const getSamplesAPI = (): Promise<AxiosResponse<any>> => {
  return API.get("/samples");
};

export const getSampleAPI = (id: number): Promise<AxiosResponse<any>> => {
  return API.get(`/samples/${id}`);
};

export const postSampleAPI = (title: string): Promise<AxiosResponse<any>> => {
  return API.post("/samples", { title: title });
};

export const putSampleAPI = (
  id: number,
  title: string
): Promise<AxiosResponse<any>> => {
  return API.put(`/samples/${id}`, { title: title });
};
