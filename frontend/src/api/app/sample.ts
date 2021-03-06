import API, { BE_BASE_URL } from "./base";
import { AxiosResponse } from "axios";
import { TypeSample } from "constants/type";

export const list = (): Promise<AxiosResponse<any>> => {
  return API.get("/samples");
};
// samples?sort=title,-content&filter[title]=eq:title&filter[content]=in:a,b&page[limit]=20&page[offset]=1

export const get = (id: string): Promise<AxiosResponse<any>> => {
  return API.get(`/samples/${id}`);
};

export const post = (sample: TypeSample): Promise<AxiosResponse<any>> => {
  return API.post("/samples", sample);
};

export const put = (id: string, sample: any): Promise<AxiosResponse<any>> => {
  return API.put(`/samples/${id}`, sample);
};

export const destroy = (id: string) => {
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
