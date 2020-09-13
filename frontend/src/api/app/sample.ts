import API from "./base";
import { AxiosResponse } from "axios";

export const getSamplesAPI = (): Promise<AxiosResponse<any>> => {
  return API.get("/samples");
};

export const getSampleAPI = (id: number): Promise<AxiosResponse<any>> => {
  return API.get(`/samples/${id}`);
};

// TODO: define the type for the form and use it in several places => prevent from the difference of type.
export const postSampleAPI = (title: string): Promise<AxiosResponse<any>> => {
  return API.post("/samples", { title: title });
};

export const putSampleAPI = (
  id: number,
  title: string
): Promise<AxiosResponse<any>> => {
  return API.put(`/samples/${id}`, { title: title });
};
