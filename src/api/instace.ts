import axios, { AxiosInstance } from "axios";

const instance: AxiosInstance = axios.create({
  baseURL: "https://jsonplaceholder.typicode.com",
  withCredentials: true,
  headers: {
    accept: "application/json"
  }
});
export default instance;
