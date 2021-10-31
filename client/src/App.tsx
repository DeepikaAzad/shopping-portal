import React, { useEffect, useState } from 'react';
import { BrowserRouter, Route } from "react-router-dom";

import './App.css';
import Login from "./pages/Login";
import Navigation from "./components/Nav";
import Register from "./pages/Register";
import CreateItem from "./pages/CreateItem";
import ItemList from "./pages/ItemList";
import ViewCart from "./pages/ViewCart";
import OrderList from "./pages/OrderList";

import { parseJwt } from "./util"

function App() {
    const [name, setName] = useState('');

    useEffect(() => {
        (
            async () => {
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
                <Navigation name={name} setName={setName} />

                <main className="form-signin">
                    <Route path="/" exact component={() => <ItemList name={name} />} />
                    <Route path="/view-cart" exact component={ViewCart} />
                    <Route path="/order-list" exact component={OrderList} />
                </main>
                <main className="form-signin">
                    <Route path="/login" component={() => <Login setName={setName} />} />
                    <Route path="/register" component={Register} />
                    <Route path="/create-item" component={CreateItem} />
                </main>
            </BrowserRouter>
        </div>
    );
}

export default App;
