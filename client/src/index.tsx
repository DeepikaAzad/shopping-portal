import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import reportWebVitals from './reportWebVitals';

import './index.css';

import axios from 'axios';

const backendUrl = process.env.REACT_APP_BACKEND_URL || "http://localhost:8080";

axios.interceptors.request.use(
    config => {
        const { origin } = new URL(config.url || "");
        const allowedOrigins = [backendUrl];
        const token = localStorage.getItem('token');
        if (allowedOrigins.includes(origin)) {
            if (config.headers) {
                config.headers.authorization = `Bearer${token}`;
            }
        }
        return config;
    },
    error => {
        return Promise.reject(error);
    }
);

axios.interceptors.response.use(undefined, (error) => {
    if (error.response && error.response.status === 401) {
        localStorage.clear();
        return window.location.href = '/login';
    }
    return Promise.reject(error);
});

ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
