import axios from 'axios';

const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor
api.interceptors.request.use((config) => {
  return config;
});

// Response interceptor - handles the unified response format
api.interceptors.response.use(
  (response) => {
    const { code, message, data } = response.data;

    // Handle success responses
    if (code === 'SUCCESS') {
      return data; // Return only the data part
    }

    // If we get a 204 No Content, return as-is
    if (response.status === 204) {
      return response;
    }

    // For other non-SUCCESS codes, treat as error
    return Promise.reject({ code, message, data });
  },
  (error) => {
    if (error.response) {
      const { code, message, details } = error.response.data;
      return Promise.reject({
        code: code || error.response.status,
        message: message || error.message,
        details,
      });
    }
    return Promise.reject(error);
  }
);

export default api;
