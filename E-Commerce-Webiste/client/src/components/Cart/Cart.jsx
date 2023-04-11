import "./Cart.scss";

import { MdClose } from "react-icons/md";
import { BsCartX } from "react-icons/bs";
import CartItem from './CartItem/CartItem'
import { Button } from "@mui/material";

const Cart = ({setShowCart}) => {
    return (
        <div className="cart-panel">
            <div className="opac-layer"></div>
            <div className="cart-content">
                <div className="cart-header">
                    <span className="heading">Cart</span>
                    <span className="close-btn" onClick={()=>setShowCart(false) }>
                        <MdClose/>
                        <span className="text">close</span>
                    </span>
                </div>
                <div className="empty-cart">
                    <BsCartX />
                    <span>Nothing in the Cart. </span>
                    <button className="return-cta">Buy Something</button>
                </div>

                <>
                    <CartItem />

                    <dev className="cart-footer">
                        <div className="subtotal">
                            <span className="text">Subtotal</span>
                            <span className="text total" >&#8377;43 </span>
                        </div>

                        <div className="button">
                            <Button className="checkout-cta"> BaaByeeee</Button>
                        </div>
                    </dev>
                </>
            </div>
        </div>
    );
};

export default Cart;
