import React, {SyntheticEvent, useState} from 'react';
import axios from 'axios';
import {Redirect} from 'react-router-dom';

import { ToastContainer, toast } from 'react-toastify';

import 'react-toastify/dist/ReactToastify.css';

const Register = () => {
    const [submitting, setSubmitting] = useState(false);
    const [fullName, setFullName] = useState('');
    const [userName, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [redirect, setRedirect] = useState(false);

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();
        setSubmitting(true);

        axios.post(process.env.REACT_APP_BACKEND_URL + '/api/user/create', {
            full_name: fullName,
            user_name: userName,
            password: btoa(password)
        }).then((_) => {
            toast.success('User created successfully. Please login.', {
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
            }, 1200);
        }).catch((error) => {
            setSubmitting(false);
            console.log(error);
            toast.error('User creation failed', {
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
        return <Redirect to="/login"/>;
    }

    return (
        <form onSubmit={submit}>
            <h1 className="h3 mb-3 fw-normal">Please register</h1>

            <input className="form-control" placeholder="Full Name" required
                   onChange={e => setFullName(e.target.value)}
            />

            <input type="text" className="form-control" placeholder="Username" required
                   onChange={e => setUsername(e.target.value)}
            />

            <input type="password" className="form-control" placeholder="Password" required
                   onChange={e => setPassword(e.target.value)}
            />

            <button className="w-100 btn btn-lg btn-primary" type="submit">
                {submitting && <span className="spinner-border spinner-border-sm mr-1"></span>}
                Create User
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

export default Register;
