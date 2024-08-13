import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost/',
  withCredentials: true,
});

api.interceptors.request.use(
  config => {
    const csrfToken = document.cookie.match(/(?:^|; )csrf_=(.*?)(?:;|$)/)?.[1];
    if (csrfToken) config.headers['X-Csrf-Token'] = csrfToken;
    return config;
  },
  error => {
    return Promise.reject(error);
  }
);

export default api;
