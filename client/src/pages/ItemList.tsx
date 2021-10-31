import React, { useState, useEffect } from 'react';
import { AgGridColumn, AgGridReact } from 'ag-grid-react';
import axios from 'axios';
import { ToastContainer, toast } from 'react-toastify';
import AddToCart from "../components/AddToCart";

import 'react-toastify/dist/ReactToastify.css';

import 'ag-grid-community/dist/styles/ag-grid.css';
import 'ag-grid-community/dist/styles/ag-theme-alpine.css';

const ItemList = (props: { name: string }) => {
  const [rowData, setRowData] = useState([]);

  useEffect(() => {
    axios.get(process.env.REACT_APP_BACKEND_URL + '/api/item/list')
      .then(response => response.data.items)
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
  }, []);

  return (
    <div className="ag-theme-alpine" style={{ height: 400, width: 400 }}>
      <AgGridReact
        frameworkComponents={{
          addToCardRenderer: AddToCart,
        }}
        rowData={rowData}>
        <AgGridColumn field="name" sortable={true} filter={true} resizable={true} headerName="Item Name"></AgGridColumn>
        <AgGridColumn field="id" cellRenderer="addToCardRenderer" headerName="ID"></AgGridColumn>
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
};

export default ItemList;
