import React, {SyntheticEvent, useState} from 'react';
import {Redirect} from 'react-router-dom';
import axios from 'axios';

import { ToastContainer, toast } from 'react-toastify';

import 'react-toastify/dist/ReactToastify.css';

const CreateItem = () => {
    const [submitting, setSubmitting] = useState(false);
    const [itemName, setItemName] = useState('');
    const [redirect, setRedirect] = useState(false);

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();
        setSubmitting(true);

        axios.post(process.env.REACT_APP_BACKEND_URL + '/api/item/create', {
            name: itemName
        }).then((_) => {
            toast.success('Item created successfully', {
                position: "bottom-right",
                autoClose: 750,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
            });
            setSubmitting(false);
            setTimeout(() => {
                setRedirect(true);
            }, 850)
        }).catch((error) => {
            console.log(error);
            setSubmitting(false);
            toast.error('Item creation failed', {
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
        return <Redirect to="/"/>;
    }

    return (
        <form onSubmit={submit}>
            <h1 className="h3 mb-3 fw-normal">Create Item</h1>

            <input type="text" className="form-control" placeholder="Item Name" required
                   onChange={e => setItemName(e.target.value)}
            /><br/>

            <button className="w-100 btn btn-lg btn-primary" type="submit">
                {submitting && <span className="spinner-border spinner-border-sm mr-1"></span>}
                Submit
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

export default CreateItem;
