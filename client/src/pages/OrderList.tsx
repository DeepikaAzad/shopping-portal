import React, { useState } from 'react';
import { AgGridColumn, AgGridReact } from 'ag-grid-react';
import axios from 'axios';
import { ToastContainer, toast } from 'react-toastify';
import "ag-grid-enterprise";

import 'react-toastify/dist/ReactToastify.css';

import 'ag-grid-community/dist/styles/ag-grid.css';
import 'ag-grid-community/dist/styles/ag-theme-alpine.css';

const OrderList = (props: { name: string }) => {
    const [gridApi, setGridApi] = useState(null);
    const [gridColumnApi, setGridColumnApi] = useState(null);
    const [rowData, setRowData] = useState(null);

    const onGridReady = (params: any) => {
        setGridApi(params.api);
        setGridColumnApi(params.columnApi);

        axios.get(process.env.REACT_APP_BACKEND_URL + '/api/order/list')
            .then(response => response.data)
            .then(rowData => setRowData(rowData))
            .catch((error) => {
                console.log(error);
                toast.error('Fetching item list failed!', {
                    position: "bottom-right",
                    autoClose: 3000,
                    hideProgressBar: false,
                    closeOnClick: true,
                    pauseOnHover: true,
                    draggable: true,
                    progress: undefined,
                });
            });
    };

    return (
        <div className="ag-theme-alpine" style={{ height: 400, width: 400 }}>
            <AgGridReact
                onGridReady={onGridReady}
                rowData={rowData}>
                <AgGridColumn
                    field="cart_id"
                    rowGroup={true}
                    hide={true}
                    valueGetter={orderValueGetter}
                    keyCreator={orderKeyCreator}
                    headerName="ID"></AgGridColumn>
                <AgGridColumn field="name"></AgGridColumn>
            </AgGridReact>
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


    function orderKeyCreator(params: any) {
        const orderObject = params.value;
        return orderObject.cart_id;
    }

    // https://www.ag-grid.com/react-data-grid/grouping-complex-objects/
    // Not yet working.
    function orderValueGetter(params: any) {
        const orderObj = params.data;
        if (orderObj == null) {
            return;
        }
        const itemNames: any = [];
        orderObj.items.forEach((item: any) => {
            itemNames.push(item.name);
        });
        return {
            cart_id: orderObj.cart_id,
            name: itemNames,
        };
    }

};

export default OrderList;
