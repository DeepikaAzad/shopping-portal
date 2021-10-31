import React, { useState } from 'react';

import axios from 'axios';
import { toast } from 'react-toastify';

import 'react-toastify/dist/ReactToastify.css';

const RemoveFromCart = (props: any) => {
  const cellValue = props.valueFormatted ? props.valueFormatted : props.value;
  const [submitting, setSubmitting] = useState(false);

  const buttonClicked = () => {
    setSubmitting(true);
    axios.post(process.env.REACT_APP_BACKEND_URL + '/api/cart/remove', {
      item_id: cellValue
    }).then((_) => {
      toast.success('Removed Successfully', {
        position: "bottom-right",
        autoClose: 1000,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
      });
      setSubmitting(false);
      window.location.reload();
    }).catch((error) => {
      setSubmitting(false);
      toast.error('Remove from cart failed: ' + error.message || "", {
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
        Remove from Cart
      </button>
    </span>
  );
};

export default RemoveFromCart;