import React from 'react';

const Login = () => {
    return (
        <div className='formContainer'>
            <div className='formWrapper'>
                <span className='logo'> chat app </span>

                <span className='title'> Login </span>
                <form>
                    <input type='text' placeholder='email address'/>
                    <input type="text" placeholder='password' />

                    <button>Login</button>
                </form>
                <p>you don't have a account ? Register</p>
            </div>
        </div>
    )
}

export default Login
