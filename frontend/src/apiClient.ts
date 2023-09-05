import axios from 'axios';

export const API_HOST = 'https://dev-voah.implude.kr';

export const apiClient = axios.create({
  baseURL: API_HOST,
});
