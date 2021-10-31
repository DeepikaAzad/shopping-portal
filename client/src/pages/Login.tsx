import React, { SyntheticEvent, useState } from 'react';
import { Redirect } from "react-router-dom";
import axios from 'axios';
import { ToastContainer, toast } from 'react-toastify';

import 'react-toastify/dist/ReactToastify.css';

import { parseJwt } from "../util"

const Login = (props: { setName: (name: string) => void }) => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [redirect, setRedirect] = useState(false);

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();

        const response = await axios.post(process.env.REACT_APP_BACKEND_URL + '/api/user/login', {
            user_name: username,
            password: btoa(password)
        }, { headers: { 'Content-Type': 'application/json' } });

        if (response.status === 200) {
            const data = await response.data;
            const userPayload = parseJwt(data.token || "");
            localStorage.setItem('token', data.token);
            toast.success('Login Successful', {
                position: "top-right",
                autoClose: 5000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
            });
            setRedirect(true);
            props.setName(userPayload.name);
        } else {
            toast.error('Invalid username/password!', {
                position: "bottom-right",
                autoClose: 3000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
            });
        }
    }

    if (redirect) {
        return <Redirect to="/" />;
    }

    return (
        <form onSubmit={submit}>
            <h1 className="h3 mb-3 fw-normal">Please sign in</h1>
            <input type="text" className="form-control" placeholder="Username" required
                onChange={e => setUsername(e.target.value)}
            />

            <input type="password" className="form-control" placeholder="Password" required
                onChange={e => setPassword(e.target.value)}
            />

            <button className="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>
            <ToastContainer
                autoClose={5000}
                hideProgressBar={false}
                newestOnTop={false}
                closeOnClick
                rtl={false}
                pauseOnFocusLoss
                draggable
                pauseOnHover
            />
        </form>
    );
};

export default Login;
