import React from 'react'
import { Products } from '../../product'
import {Product} from './product'
import './shop.css'

export const Shop = () => {
  return (
    <div className='shop'>
        <div className='shopTitle'>
            <h1> Toi Moi Shop</h1>
        </div>
        <div className='products'>
            {""}
            {Products.map((product) =>(
                <Product data={product}/>
            ))}
        </div>
    </div>
  )
}

