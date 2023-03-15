import React from "react";

const Navbar = () => {
    return (
        <div className="navbar">
            <span className="logo"> Chat App</span>
            <div className="user">
                <img src="https://webusupload.apowersoft.info/picwishcom/wp-content/uploads/2021/11/unblur-image.jpg" alt=""/>
                <span> Johna</span>
                <button> logout</button>
            </div>
        </div>
    )
}
export default Navbar
