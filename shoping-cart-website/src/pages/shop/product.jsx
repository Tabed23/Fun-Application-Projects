import React  from "react";
import { useContext } from "react";
import { ShopContext } from "../../context/shop-context";

export const Product = (props) => {
  const {id, name, price, image} = props.data
  const { addToCart, cartItems } = useContext(ShopContext);
  const cartItemCount = cartItems[id];
  return (
    <div className='product'>

      <img src={image}/>

      <div className='product-info'>
        <p>
          <b>{name}</b>
        </p>

        <p>
          <b> ${price}</b>
        </p>

      </div>

      <button className="addToCartBttn" onClick={() => addToCart(id)}>
        Add To Cart {cartItemCount > 0 && <> ({cartItemCount})</>}
      </button>

    </div>
  )
}
