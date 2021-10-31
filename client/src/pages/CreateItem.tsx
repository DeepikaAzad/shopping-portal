import React, {SyntheticEvent, useState} from 'react';
import {Redirect} from 'react-router-dom';
import axios from 'axios';

const CreateItem = () => {
    const [itemName, setItemName] = useState('');
    const [redirect, setRedirect] = useState(false);

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();

        await axios.post(process.env.REACT_APP_BACKEND_URL + '/api/item/create', {
            name: itemName
        }, { headers: { 'Content-Type': 'application/json' } });

        setRedirect(true);
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

            <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
        </form>
    );
};

export default CreateItem;
