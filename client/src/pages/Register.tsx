import React, {SyntheticEvent, useState} from 'react';
import axios from 'axios';
import {Redirect} from 'react-router-dom';

const Register = () => {
    const [fullName, setFullName] = useState('');
    const [userName, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [redirect, setRedirect] = useState(false);

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();

        await axios.post(process.env.REACT_APP_BACKEND_URL + '/api/user/create', {
            full_name: fullName,
            user_name: userName,
            password: btoa(password)
        }, { headers: { 'Content-Type': 'application/json' } });

        setRedirect(true);
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

            <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
        </form>
    );
};

export default Register;
