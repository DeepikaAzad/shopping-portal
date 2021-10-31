import React, { SyntheticEvent, useState } from 'react';
import { Redirect } from "react-router-dom";
import axios from 'axios';
import { ToastContainer, toast } from 'react-toastify';

import 'react-toastify/dist/ReactToastify.css';

import { parseJwt } from "../util"

const Login = (props: { setName: (name: string) => void }) => {
    const [submitting, setSubmitting] = useState(false);
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [redirect, setRedirect] = useState(false);

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();
        setSubmitting(true);

        axios.post(process.env.REACT_APP_BACKEND_URL + '/api/user/login', {
            user_name: username,
            password: btoa(password)
        }).then((response) => {
            const data = response.data;
            const userPayload = parseJwt(data.token || "");
            localStorage.setItem('token', data.token);
            toast.success('Login Successful', {
                position: "bottom-right",
                autoClose: 1000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
            });
            setSubmitting(false);
            setTimeout(() => {
                setRedirect(true);
                props.setName(userPayload.name);
            }, 1250)
        }).catch((error) => {
            setSubmitting(false);
            console.log("ckadjs");
            toast.error('Invalid username/password!', {
                position: "bottom-right",
                autoClose: 3000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
            });
        });
    }

    if (redirect) {
        return <Redirect to="/" />;
    }

    return (
        <form onSubmit={submit}>
            {/* Same as */}
            <ToastContainer />
            <h1 className="h3 mb-3 fw-normal">Please sign in</h1>
            <input type="text" className="form-control" placeholder="Username" required
                onChange={e => setUsername(e.target.value)}
            />

            <input type="password" className="form-control" placeholder="Password" required
                onChange={e => setPassword(e.target.value)}
            />

            <button className="w-100 btn btn-lg btn-primary" type="submit">
                {submitting && <span className="spinner-border spinner-border-sm mr-1"></span>}
                 Sign in
            </button>
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
