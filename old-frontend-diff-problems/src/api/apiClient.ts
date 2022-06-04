import axios from "axios";

export const apiClientV1 = axios.create({
  baseURL: process.env.REACT_APP_API_BASE_URL + "api/v1",
  headers: {
    "Content-type": "application/json",
  },
});
