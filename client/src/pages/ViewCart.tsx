import React, { useState, useEffect } from 'react';
import { AgGridColumn, AgGridReact } from 'ag-grid-react';
import axios from 'axios';
import { ToastContainer, toast } from 'react-toastify';
import RemoveFromCart from "../components/RemoveFromCart";

import 'react-toastify/dist/ReactToastify.css';

import 'ag-grid-community/dist/styles/ag-grid.css';
import 'ag-grid-community/dist/styles/ag-theme-alpine.css';

const ViewCart = () => {
    const [submitting, setSubmitting] = useState(false);
    const [rowData, setRowData] = useState([]);
    const [cartId, setCartId] = useState(null);

    useEffect(() => {
        axios.get(process.env.REACT_APP_BACKEND_URL + '/api/cart/list')
            .then(response => {
                if (response.data.cart_id && response.data.cart_id !== 0) {
                    setCartId(response.data.cart_id);
                } else {
                    setCartId(null);
                }
                return response.data.item_list || [];
            })
            .then(rowData => setRowData(rowData))
            .catch((error) => {
                console.log(error);
                toast.error('Fetching cart list failed!', {
                    position: "bottom-right",
                    autoClose: 3000,
                    hideProgressBar: false,
                    closeOnClick: true,
                    pauseOnHover: true,
                    draggable: true,
                    progress: undefined,
                });
            });
    }, [cartId]);

    function placeOrder() {
        setSubmitting(true);
        axios.post(process.env.REACT_APP_BACKEND_URL + `/api/cart/${cartId}/complete`)
            .then((_) => {
                toast.success('Order Placed Successfully!', {
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
                    return window.location.href = '/';
                }, 1250)
            }).catch((error) => {
                setSubmitting(false);
                toast.error('Order failed: ' + error.response.data.message, {
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

    return (
        <div className="ag-theme-alpine">
            <div style={{ height: 400, width: 400 }}>

                <div>
                    {/* Cart List */}
                </div>
                <AgGridReact
                    frameworkComponents={{
                        removeFromCartRenderer: RemoveFromCart,
                    }}
                    rowData={rowData}>
                    <AgGridColumn field="name" sortable={true} filter={true} resizable={true} headerName="Item Name"></AgGridColumn>
                    <AgGridColumn field="id" cellRenderer="removeFromCartRenderer" headerName="ID" enableValue={false}></AgGridColumn>
                </AgGridReact>
            </div><br />
            <button className="btn btn-primary btn-lg" onClick={() => placeOrder()} disabled={cartId == null}>
                {submitting && <span className="spinner-border spinner-border-sm mr-1"></span>}
                Place Order
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
        </div>
    );
};
export default ViewCart;
