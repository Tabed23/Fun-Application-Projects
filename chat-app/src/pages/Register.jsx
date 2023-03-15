import React from 'react';
import UploadImage from '../img/addAvatar.png'

const Register = () => {
    return (
        <div className='formContainer'>
            <div className='formWrapper'>
                <span className='logo'> chat app </span>

                <span className='title'> register </span>
                <form>
                    <input type='text' placeholder='display name'/>
                    <input type='text' placeholder='email address'/>
                    <input type="text" placeholder='password' />
                    <input type="text" placeholder='phone number' />
                    <input style={{display: "none"}} type='file' id='file'/>
                    <label htmlFor='file'>
                        <img src={UploadImage} alt="" />
                        <span> Upload Picture</span>
                    </label>

                    <button>Sign Up</button>
                </form>
                <p>do you have a account already ? Login</p>
            </div>
        </div>
    )
}

export default Register
