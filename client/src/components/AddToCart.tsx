import React, { useState } from 'react';

import axios from 'axios';
import { toast } from 'react-toastify';

import 'react-toastify/dist/ReactToastify.css';

export default (props: any) => {
  const cellValue = props.valueFormatted ? props.valueFormatted : props.value;
  const [submitting, setSubmitting] = useState(false);

  const buttonClicked = () => {
    setSubmitting(true);
    axios.post(process.env.REACT_APP_BACKEND_URL + '/api/cart/add', {
      item_id: cellValue
    }).then((_) => {
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
    }).catch((error) => {
      setSubmitting(false);
      console.log("ckadjs");
      toast.error('Add to cart failed!', {
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
    <span>
      <span>{cellValue}</span>&nbsp;
      <button className="btn btn-secondary btn-sm" onClick={() => buttonClicked()}>
        {submitting && <span className="spinner-border spinner-border-sm mr-1"></span>}
        Add to Cart
      </button>
    </span>
  );
};