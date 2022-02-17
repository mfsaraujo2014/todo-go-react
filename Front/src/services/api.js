import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:5000",
  withCredentials: false,
        headers: {
          'Access-Control-Allow-Origin' : '*',
          'Access-Control-Allow-Methods':'GET,PUT,POST,DELETE,PATCH,OPTIONS',   
      }
});

export default api;