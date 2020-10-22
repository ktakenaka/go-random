import API, { BE_BASE_URL } from "./base";
import { AxiosResponse } from "axios";
import { TypeSample } from "constants/type";

export const list = (): Promise<AxiosResponse<any>> => {
  return API.get("/samples");
};

export const get = (id: number): Promise<AxiosResponse<any>> => {
  return API.get(`/samples/${id}`);
};

export const post = (sample: TypeSample): Promise<AxiosResponse<any>> => {
  return API.post("/samples", sample);
};

export const put = (id: number, sample: any): Promise<AxiosResponse<any>> => {
  return API.put(`/samples/${id}`, sample);
};

export const destroy = (id: number) => {
  return API.delete(`/samples/${id}`);
};

export const exportURL = (charset: string): string => {
  return `${BE_BASE_URL}/export/samples?charset=${charset}`;
};

export const importFile = (file: File): Promise<AxiosResponse<any>> => {
  const formData = new FormData();
  formData.append("file", file);

  return API.post("/samples/import", formData, {
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
};
