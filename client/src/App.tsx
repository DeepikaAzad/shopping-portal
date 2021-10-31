import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { BrowserRouter, Route } from "react-router-dom";

import './App.css';
import Login from "./pages/Login";
import Nav from "./components/Nav";
import Home from "./pages/Home";
import Register from "./pages/Register";
import CreateItem from "./pages/CreateItem";

import { parseJwt } from "./util"

const backendUrl = process.env.REACT_APP_BACKEND_URL || "";

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

function App() {
    const [name, setName] = useState('');

    useEffect(() => {
        (
            async () => {
                // const response = await fetch(process.env.REACT_APP_BACKEND_URL + '/api/user/login', {
                //     headers: { 'Content-Type': 'application/json' },
                //     credentials: 'include',
                // });

                // const content = await response.json();
                const storedJwt = localStorage.getItem('token');
                if (storedJwt) {
                    const userInfo = parseJwt(storedJwt);
                    setName(userInfo.name);
                }
            }
        )();
    });


    return (
        <div className="App">
            <BrowserRouter>
                <Nav name={name} setName={setName} />

                <main className="form-signin">
                    <Route path="/" exact component={() => <Home name={name} />} />
                    <Route path="/login" component={() => <Login setName={setName} />} />
                    <Route path="/register" component={Register} />
                    <Route path="/create-item" component={CreateItem} />
                </main>
            </BrowserRouter>
        </div>
    );
}

export default App;
